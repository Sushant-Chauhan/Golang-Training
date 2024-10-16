/*
Design :
Game interface -  facade design pattern
TicTacToe struct  - implements the Game interface (board and players)
NewGame method   - sets up the game, players, and the board.
play method - updates the board and checks for a winner.
TPlayer, Board, and Cell structs - for tictactoe game
*/

/*//////////////////////////// Game deign ////////////////////

package main
import (
    "fmt"
)
func main() {
   var g1 Game = NewGame()   // or NewTicTacToe()
//    var g2 Game = NewTicTaeToe()    - 2nd game on another board 

   g1.play(2) //1st player - 'X'
   g1.play(5) // '0'
   g1.play(3) // 'X'
   g1.play(1)// '0'
   g1.play(8) //'X'
   g1.play(4)// '0'
   g1.play(6) //'X'
   g1.play(7) //1
   var g1 Game = &Game{} //don't know is it required	
   var g1 Game = &Game{} 
   var g1 Game = &Game{}

 }

// Facade design pattern
type Game interface {
	play(int)
	NewGame(playerNames ...string)   //makeing Game generised to make scalable for other Games 
}

type TicTaeToe struct {
	//code
}

func NewGame(player1Name, player2Name string) {
	player1 := NewPlayer(player1Name, 'X')
	player2 := NewPlayer(player2Name, '0')
	player2 := NewBoard()

}

func (g Game) play(cellNo int) {

}

type Player struct {s
	name string
	symbol string
}

type Board struct {
	// cell [9]Cell     //instead of 2D matrix of 3*3 , taking 1D matrix of 9 cells
}

type Cell struct {
	symbol string
	isEmpty bool
}

*/