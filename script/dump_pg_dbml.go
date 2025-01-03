package script

import (
	"context"
	"fmt"

	"github.com/josudoey/judo/cmd"
	"github.com/josudoey/judo/core"
	"github.com/josudoey/judo/dbml"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var _ = cmd.AddCommand(
	&cobra.Command{
		Use: "dump-pg-dbml",
		RunE: cmd.Run(func(ctx context.Context, args []string) error {
			logger := core.UseLogger(ctx)
			zap.Any("logger", logger)
			script := NewPGXScript(ctx)

			databaseName, err := script.GetCurrentDatabaseName(ctx)
			if err != nil {
				return err
			}

			tableNames, err := script.ListTableNames(ctx)
			if err != nil {
				return err
			}

			tableColumns, err := script.ListTableColumn(ctx)
			if err != nil {
				return err
			}

			tableIndexes, err := script.ListTableIndex(ctx)
			if err != nil {
				return err
			}

			foreignKeyConstraints, err := script.ListForeignKeyConstraint(ctx)
			if err != nil {
				return err
			}

			doc := dbml.New()
			doc.Append(&dbml.Project{
				Name:         databaseName,
				DatabaseType: dbml.PostgreSQL,
			})

			tableMap := map[string]*dbml.Table{}
			tableFooters := map[string][]dbml.Element{}
			columnMap := map[string]map[string]*dbml.Column{}
			columnUniqMap := map[string]map[string]bool{}
			for _, column := range tableColumns {
				table, ok := tableMap[column.TableName]
				if !ok {
					table = &dbml.Table{Name: column.TableName}
					tableMap[column.TableName] = table
				}

				c := &dbml.Column{
					Name:     column.ColumnName,
					Type:     column.DataType,
					Settings: NewColumnSettings(column),
				}
				if columnMap[column.TableName] == nil {
					columnMap[column.TableName] = map[string]*dbml.Column{}
				}

				if columnUniqMap[column.TableName] == nil {
					columnUniqMap[column.TableName] = map[string]bool{}
				}
				columnMap[column.TableName][column.ColumnName] = c
				table.Columns = append(table.Columns, c)
			}

			for _, tableIndex := range tableIndexes {
				if len(tableIndex.IndexedColumnNames) == 1 {
					if tableIndex.IsPrimary || tableIndex.IsUnique {
						columnUniqMap[tableIndex.TableName][tableIndex.IndexedColumnNames[0]] = true
					}
				}

				for _, columnName := range tableIndex.IndexedColumnNames {
					m, ok := columnMap[tableIndex.TableName]
					if !ok {
						continue
					}

					if tableIndex.IsPrimary {
						column, ok := m[columnName]
						if !ok {
							continue
						}
						column.Settings = append([]dbml.ColumnSetting{dbml.PK()}, column.Settings...)
					}
				}

				t, ok := tableMap[tableIndex.TableName]
				if !ok {
					continue
				}

				tableIndexSettings := []dbml.TableIndexSetting{}
				if tableIndex.IsPrimary {
					tableIndexSettings = append(tableIndexSettings, dbml.PK())
				}

				if tableIndex.IsUnique {
					tableIndexSettings = append(tableIndexSettings, dbml.Unique())
				}
				tableIndexSettings = append(tableIndexSettings, dbml.TableIndexType(tableIndex.IndexType))

				if tableIndex.IndexCondition != nil {
					tableIndexSettings = append(tableIndexSettings, dbml.TableIndexNote("IndexCondition: "+*tableIndex.IndexCondition))
				}
				tableIndexSettings = append(tableIndexSettings, dbml.TableIndexName(tableIndex.IndexName))

				t.TableIndexes = append(t.TableIndexes, &dbml.TableIndex{
					ColumnNames: tableIndex.IndexedColumnNames,
					Settings:    tableIndexSettings,
				})

				if tableIndex.IsPrimary {
					continue
				}
				if len(tableFooters[tableIndex.TableName]) == 0 {
					tableFooters[tableIndex.TableName] = append(tableFooters[tableIndex.TableName], dbml.SingleLineComment("--- Indices DDL"))
				}
				tableFooters[tableIndex.TableName] = append(tableFooters[tableIndex.TableName], dbml.SingleLineComment(tableIndex.IndexDefinition))
			}

			for _, fk := range foreignKeyConstraints {
				localColumn := columnMap[fk.LocalTableName][fk.LocalColumnName]
				if localColumn == nil {
					continue
				}

				localUniq := columnUniqMap[fk.LocalTableName][fk.LocalColumnName]
				foreignUniq := columnUniqMap[fk.ForeignTable][fk.ForeignColumnName]
				if localUniq {
					if foreignUniq {
						localColumn.Settings = append(localColumn.Settings, dbml.ColumnOneToOne(fk.ForeignTable, fk.ForeignColumnName))
					} else {
						localColumn.Settings = append(localColumn.Settings, dbml.ColumnOneToMany(fk.ForeignTable, fk.ForeignColumnName))
					}
				} else {
					if foreignUniq {
						localColumn.Settings = append(localColumn.Settings, dbml.ColumnManyToOne(fk.ForeignTable, fk.ForeignColumnName))
					} else {
						localColumn.Settings = append(localColumn.Settings, dbml.ColumnManyToMany(fk.ForeignTable, fk.ForeignColumnName))
					}
				}
			}

			for _, tableName := range tableNames {
				table, ok := tableMap[tableName]
				if !ok {
					continue
				}

				doc.Append(table)
				for _, footer := range tableFooters[tableName] {
					doc.Append(footer)
				}
				doc.Append(dbml.Blankline())
			}

			fmt.Printf("%v", doc)
			return err
		}),
	},
)

func (s *PgxScript) GetCurrentDatabaseName(ctx context.Context) (string, error) {
	rows, err := s.PgxConn.Query(ctx, `SELECT current_database();`)
	if err != nil {
		return "", err
	}

	defer rows.Close()

	var result string
	for rows.Next() {
		if err := rows.Scan(
			&result,
		); err != nil {
			return "", err
		}
	}

	if rows.Err() != nil {
		return "", err
	}

	return result, nil
}

func (s *PgxScript) ListTableNames(ctx context.Context) ([]string, error) {
	rows, err := s.PgxConn.Query(ctx, `
SELECT table_name
FROM information_schema.tables 
WHERE 
    table_schema='public'
    AND table_type='BASE TABLE'
ORDER BY table_name`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []string
	for rows.Next() {
		var item string
		if err := rows.Scan(&item); err != nil {
			return nil, err
		}

		result = append(result, item)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return result, nil
}

func NewColumnSettings(c *TableColumn) []dbml.ColumnSetting {
	var result []dbml.ColumnSetting

	if c.ColumnDefault != nil {
		result = append(result, dbml.DefaultString(*c.ColumnDefault))
	}

	if c.IsNullable == "NO" {
		result = append(result, dbml.NotNull())
	}

	return result
}

type TableColumn struct {
	TableName       string
	DataType        string
	ColumnName      string
	IsNullable      string
	ColumnDefault   *string
	OrdinalPosition int
}

func (s *PgxScript) ListTableColumn(ctx context.Context) ([]*TableColumn, error) {
	rows, err := s.PgxConn.Query(ctx, `
SELECT
    table_name,
    data_type,
    column_name,
    is_nullable,
    column_default,
    ordinal_position
FROM
    information_schema.columns
WHERE
    table_schema='public'
ORDER BY table_name, ordinal_position;`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []*TableColumn
	for rows.Next() {
		var item TableColumn
		if err := rows.Scan(
			&item.TableName,
			&item.DataType,
			&item.ColumnName,
			&item.IsNullable,
			&item.ColumnDefault,
			&item.OrdinalPosition,
		); err != nil {
			return nil, err
		}

		result = append(result, &item)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return result, nil
}

type ForeignKeyConstraint struct {
	ForeignKeyName    string
	LocalTableName    string
	LocalColumnName   string
	ForeignTable      string
	ForeignColumnName string
}

func (s *PgxScript) ListForeignKeyConstraint(ctx context.Context) ([]*ForeignKeyConstraint, error) {
	rows, err := s.PgxConn.Query(ctx, `
SELECT
    tc.constraint_name AS foreign_key_name,
    tc.table_name AS local_table,
    kcu.column_name AS local_column,
    ccu.table_name AS foreign_table,
    ccu.column_name AS foreign_column
FROM 
    information_schema.table_constraints AS tc
JOIN 
    information_schema.key_column_usage AS kcu
    ON tc.constraint_name = kcu.constraint_name
    AND tc.table_schema = kcu.table_schema
JOIN 
    information_schema.constraint_column_usage AS ccu
    ON ccu.constraint_name = tc.constraint_name
    AND ccu.table_schema = tc.table_schema
WHERE 
    tc.constraint_type = 'FOREIGN KEY';`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []*ForeignKeyConstraint
	for rows.Next() {
		var item ForeignKeyConstraint
		if err := rows.Scan(
			&item.ForeignKeyName,
			&item.LocalTableName,
			&item.LocalColumnName,
			&item.ForeignTable,
			&item.ForeignColumnName,
		); err != nil {
			return nil, err
		}

		result = append(result, &item)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return result, nil
}

type TableIndex struct {
	TableName          string
	IndexName          string
	IndexType          string
	IsUnique           bool
	IsPrimary          bool
	IndexDefinition    string
	IndexCondition     *string
	IndexedColumnNames []string
}

func (s *PgxScript) ListTableIndex(ctx context.Context) ([]*TableIndex, error) {
	rows, err := s.PgxConn.Query(ctx, `
WITH unnest_index AS (SELECT 
  ix.indexrelid,
  ix.indrelid,
  unnest.indkey,
  unnest.ordinal
FROM pg_index ix
LEFT JOIN LATERAL UNNEST(ix.indkey) WITH ORDINALITY AS UNNEST(indkey, ordinal) ON true)
,indexed_columns AS (SELECT
    uix.indrelid,
    uix.indexrelid,
    ARRAY_AGG(a.attname ORDER BY uix.ordinal) AS names
FROM
    pg_class t,
    pg_class i,
    unnest_index uix,
    pg_attribute a
WHERE
    t.oid = uix.indrelid
    AND i.oid = uix.indexrelid
    AND a.attrelid = t.oid
    AND a.attnum = uix.indkey
    AND t.relkind = 'r'
GROUP BY
    uix.indrelid,
    uix.indexrelid
)
SELECT
    t.relname AS table_name,
    i.relname AS index_name,
    am.amname AS index_type,
    ix.indisunique AS is_unique,
    ix.indisprimary AS is_primary,
    pg_get_indexdef(ix.indexrelid) AS index_definition,
    pg_get_expr(ix.indpred, ix.indrelid) AS index_condition,
    ic.names AS index_column_names
FROM
    pg_class t,
    pg_class i,
    pg_index ix,
    pg_am am,
    indexed_columns ic
WHERE
    t.relnamespace = (SELECT oid FROM pg_namespace WHERE nspname = 'public')
    AND t.oid = ix.indrelid
    AND i.oid = ix.indexrelid
    AND i.relam = am.oid
    AND ic.indrelid = ix.indrelid
    AND ic.indexrelid = ix.indexrelid
ORDER BY
    t.relname,
    i.relname
;`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []*TableIndex
	for rows.Next() {
		var item TableIndex
		if err := rows.Scan(
			&item.TableName,
			&item.IndexName,
			&item.IndexType,
			&item.IsUnique,
			&item.IsPrimary,
			&item.IndexDefinition,
			&item.IndexCondition,
			&item.IndexedColumnNames,
		); err != nil {
			return nil, err
		}

		result = append(result, &item)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return result, nil
}
