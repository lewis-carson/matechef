package main

//import "testing"

import (
	"fmt"
	"testing"

	"github.com/notnil/chess"
	matechef "github.com/torvim/matechef/engine"
)

func TestMain(m *testing.M) {
	t := matechef.Node{
		Position: chess.NewGame().Position(),
	}

	s := matechef.Settings{
		MinDepth: 4,
	}

	move, _ := t.Search(s)
	fmt.Println(move)
}

func TestM1(t *testing.T) {
	fenStr := "6k1/pp3ppp/2b5/8/7N/6P1/P4P1P/3rQ1K1 b - - 0 1"
	fen, _ := chess.FEN(fenStr)
	g := chess.NewGame(fen).Position()

	tree := matechef.Node{
		Position: g,
	}

	s := matechef.Settings{
		MinDepth: 4,
	}

	_, score := tree.Search(s)
	if score != 99999999 {
		t.Errorf("Didn't find mate")
	}
}
