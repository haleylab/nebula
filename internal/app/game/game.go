package game

import (
	"context"
	"fmt"
)

type Game struct{}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Tick(ctx context.Context) error {
	fmt.Printf("Press Enter for next tick or Ctrl-C to quit.")
	fmt.Scanln()
	return nil
}
