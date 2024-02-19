package main

import (
	"context"
	"fmt"
	"os"

	"github.com/haleylab/nebula/internal/app/game"
)

func main() {
	ctx := context.Background()
	if err := run(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	g := game.NewGame()
	for {
		if err := g.Tick(ctx); err != nil {
			return err
		}
	}
}
