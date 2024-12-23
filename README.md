# judo usage

provide utils tool for development

```sh
go install github.com/josudoey/judo/cmd/judo@latest
export POSTGRES_URL=postgres://postgres@localhost:5432/postgres?sslmode=disable
judo dump-pg-dbml

# Output:
# Project postgres {
#   database_type: 'PostgreSQL'
# }
```
