package main

import (
	//"math"
	//"reflect"
	//"time"
	//"github.com/notnil/chess/uci"

	"strings"

	"github.com/abiosoft/ishell"
	"github.com/notnil/chess"
)

func main() {
	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.New()

	// display welcome info.
	shell.Println("Kimchi v2.1")
	shell.ShowPrompt(false)

	game := chess.NewGame()

	// register a function for "greet" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "uci",
		Func: func(c *ishell.Context) {
			c.Println("id name mateChef")
			c.Println("id author lew")
			c.Println("uciok")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "isready",
		Func: func(c *ishell.Context) {
			c.Println("readyok")
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "ucinewgame",
		Func: func(c *ishell.Context) {
			return
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "position",
		Func: func(c *ishell.Context) {
			if c.Args[0] == "startpos" {
				game = chess.NewGame()
				if len(c.Args) < 3 {
					game = chess.NewGame()
					return
				}
				moves := c.Args[2:len(c.Args)]
				for _, move := range moves {
					m, _ := chess.UCINotation{}.Decode(nil, move)
					game.Move(m)
				}
				return
			}
			fenString := strings.Join(c.Args, " ")
			fen, _ := chess.FEN(fenString)
			game = chess.NewGame(fen)
		},
	})

	shell.AddCmd(&ishell.Cmd{
		Name: "go",
		Func: func(c *ishell.Context) {
			return
		},
	})

	// run shell
	shell.Run()
}
