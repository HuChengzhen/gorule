package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"hcz_go_rule/gorule/gotype"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	//game := gotype.NewGame(19)
	//move := gotype.NewPlayMove(gotype.NewPoint(1, 1))
	//game = game.ApplyMove(&move)
	//move = gotype.NewPlayMove(gotype.NewPoint(1, 2))
	//game = game.ApplyMove(&move)
	//move = gotype.NewPlayMove(gotype.NewPoint(16, 16))
	//game = game.ApplyMove(&move)
	//move = gotype.NewPlayMove(gotype.NewPoint(2, 1))
	//game = game.ApplyMove(&move)
	//
	//drawBoard(game.Board)

	http.HandleFunc("/ws", handleConnections)
	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}

type Point struct {
	Row uint8 `json:"row"`
	Col uint8 `json:"col"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	game := gotype.NewGame(19)

	for {
		var msg Point
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Printf("Received: %s\n", msg)
		move := gotype.NewPlayMove(gotype.NewPoint(msg.Row, msg.Col))
		if game.IsValidMove(move) {
			game = game.ApplyMove(&move)
		} else {
			fmt.Println("Invalid move")
		}
		drawBoard(game.Board)

		err = ws.WriteJSON("ok")
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

// drawBoard draws the board to the console.
func drawBoard(board *gotype.Board) {
	for i := uint8(1); i <= board.NumRows; i++ {
		for j := uint8(1); j <= board.NumCols; j++ {
			point := gotype.NewPoint(20-i, j)
			switch board.GetColor(point) {
			case gotype.Black:
				print("B ")
			case gotype.White:
				print("W ")
			default:
				print(". ")
			}
		}
		println()
	}
}
