package main

import (
	"context"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/josudoey/judo/cmd"
	_ "github.com/josudoey/judo/script"
)

func main() {
	if err := cmd.NewCommand().ExecuteContext(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
