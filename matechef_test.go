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
	fmt.Println(t.Search())
}
