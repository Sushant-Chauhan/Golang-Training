// /*
// Game interface -  facade design pattern
// TicTacToe struct  - implements the Game interface (board and players)
// NewGame method   - sets up the game, players, and the board.
// play method - updates the board and checks for a winner.
// TPlayer, Board, and Cell structs - for tictactoe game
// */


// package main

// import (
// 	"errors"
// 	"fmt"
// )

// // Facade Design Pattern
// type Game interface {
// 	Play(cellNo int) error
// 	NewGame(player1Name, player2Name string) *TicTacToe
// }

// // TicTacToe struct - game logic
// type TicTacToe struct {
// 	board   *Board
// 	player1 *Player
// 	player2 *Player
// 	current *Player
// }

// // NewGame - initializes new TicTacToe game
// func NewGame(player1Name, player2Name string) *TicTacToe {
// 	player1 := NewPlayer(player1Name, "X")
// 	player2 := NewPlayer(player2Name, "O")
// 	board := NewBoard()
// 	return &TicTacToe{
// 		board:   board,
// 		player1: player1,
// 		player2: player2,
// 		current: player1, // Player 1 starts the game
// 	}
// }

// // Play - handles the move for the current player and switches turns
// func (t *TicTacToe) Play(cellNo int) error {
// 	if err := t.board.UpdateBoard(cellNo, t.current.symbol); err != nil {
// 		return err
// 	}

// 	t.board.Display()

// 	// Check for winner
// 	if t.board.CheckWin(t.current.symbol) {
// 		fmt.Printf("Player %s (%s) wins!\n", t.current.name, t.current.symbol)
// 		return nil
// 	}

// 	// Switch turn
// 	t.switchTurn()

// 	return nil
// }

// // switchTurn switches the active player
// func (t *TicTacToe) switchTurn() {
// 	if t.current == t.player1 {
// 		t.current = t.player2
// 	} else {
// 		t.current = t.player1
// 	}
// }

// // Player struct - player in the game
// type Player struct {
// 	name   string
// 	symbol string
// }

// // NewPlayer - creates a new player
// func NewPlayer(name, symbol string) *Player {
// 	return &Player{
// 		name:   name,
// 		symbol: symbol,
// 	}
// }

// // Board struct - represents game board
// type Board struct {
// 	cells [9]Cell
// }

// // NewBoard creates a new, empty game board
// func NewBoard() *Board {
// 	board := &Board{}
// 	for i := 0; i < 9; i++ {
// 		board.cells[i] = Cell{symbol: " ", isEmpty: true}
// 	}
// 	return board
// }

// // UpdateBoard places the player's symbol on the chosen cell
// func (b *Board) UpdateBoard(cellNo int, symbol string) error {
// 	if cellNo < 1 || cellNo > 9 {
// 		return errors.New("invalid cell number")
// 	}
// 	if !b.cells[cellNo-1].isEmpty {
// 		return errors.New("cell already taken")
// 	}

// 	b.cells[cellNo-1] = Cell{symbol: symbol, isEmpty: false}
// 	return nil
// }

// // Display  - prints the current state of the board after each turn
// func (b *Board) Display() {
// 	fmt.Println("Board:")
// 	for i := 0; i < 9; i += 3 {
// 		fmt.Printf("%s | %s | %s\n", b.cells[i].symbol, b.cells[i+1].symbol, b.cells[i+2].symbol)
// 		if i < 6 {
// 			fmt.Println("---------")
// 		}
// 	}
// 	fmt.Println()
// }

// // CheckWin - checks if the current player has won
// func (b *Board) CheckWin(symbol string) bool {
// 	winConditions := [][3]int{
// 		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Rows
// 		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Columns
// 		{0, 4, 8}, {2, 4, 6},            // Diagonals
// 	}

// 	for _, condition := range winConditions {
// 		if b.cells[condition[0]].symbol == symbol &&
// 			b.cells[condition[1]].symbol == symbol &&
// 			b.cells[condition[2]].symbol == symbol {
// 			return true
// 		}
// 	}

// 	return false
// }

// // Cell struct represents a cell on the board
// type Cell struct {
// 	symbol  string
// 	isEmpty bool
// }

// func main() {
// 	// Initialize a new TicTacToe game
// 	game := NewGame("Player 1", "Player 2")

// 	// Simulate moves
// 	game.Play(2)  // Player 1 ('X')
// 	game.Play(5)  // Player 2 ('O')
// 	game.Play(3)  // Player 1 ('X')
// 	game.Play(1)  // Player 2 ('O')
// 	game.Play(8)  // Player 1 ('X')
// 	game.Play(4)  // Player 2 ('O')
// 	game.Play(6)  // Player 1 ('X')
// 	game.Play(7)  // Player 2 ('O')
// 	game.Play(5)  // Player 2 ('O')

// 	game.Play(3)  // Player 2 ('O')


//  }

package main

import (
	"errors"
	"fmt"
)

// Game interface for generalization and scalability
type Game interface {
	Play(cellNo int) error
	// NewGame(player1Name, player2Name string) *TicTacToe

}

// TicTacToe struct implements the Game interface
type TicTacToe struct {
	board   *Board
	player1 *Player
	player2 *Player
	current *Player
	isOver  bool
	winner  *Player
}

// NewTicTacToe initializes a new TicTacToe game with two players
func NewTicTacToe(player1Name, player2Name string) Game {
	player1 := NewPlayer(player1Name, "X")
	player2 := NewPlayer(player2Name, "O")
	board := NewBoard()
	return &TicTacToe{
		board:   board,
		player1: player1,
		player2: player2,
		current: player1, // Player 1 starts the game
		isOver:  false,
	}
}

// Play handles the move for the current player and checks for the game status
func (t *TicTacToe) Play(cellNo int) error {
	// Game already over, prevent further moves
	if t.isOver {
		return errors.New("the game is over, no more moves are allowed")
	}

	// Validate the cell number is within the valid range
	if cellNo < 1 || cellNo > 9 {
		return errors.New("invalid cell number, please choose a number between 1 and 9")
	}

	// Attempt to make a move on the board
	if err := t.board.UpdateBoard(cellNo, t.current.symbol); err != nil {
		return err
	}

	// Display the updated board after the move
	t.DisplayBoard()

	// Check for a winner
	if t.board.CheckWin(t.current.symbol) {
		t.isOver = true
		t.winner = t.current
		fmt.Printf("Player %s (%s) wins!\n", t.current.name, t.current.symbol)
		return nil
	}

	// Check for a draw (if all cells are filled and no winner)
	if t.board.IsDraw() {
		t.isOver = true
		fmt.Println("The game is a draw!")
		return nil
	}

	// Switch the turn to the other player
	t.switchTurn()

	return nil
}

// DisplayBoard prints the current state of the board
func (t *TicTacToe) DisplayBoard() {
	t.board.Display()
}

// switchTurn switches the active player
func (t *TicTacToe) switchTurn() {
	if t.current == t.player1 {
		t.current = t.player2
	} else {
		t.current = t.player1
	}
}

// Player struct represents a player in the game
type Player struct {
	name   string
	symbol string
}

// NewPlayer creates a new player with a name and symbol (X or O)
func NewPlayer(name, symbol string) *Player {
	return &Player{
		name:   name,
		symbol: symbol,
	}
}

// Board struct represents the game board
type Board struct {
	cells [9]Cell
}

// NewBoard creates a new, empty game board
func NewBoard() *Board {
	board := &Board{}
	for i := 0; i < 9; i++ {
		board.cells[i] = Cell{symbol: " ", isEmpty: true}
	}
	return board
}

// UpdateBoard places the player's symbol on the chosen cell
func (b *Board) UpdateBoard(cellNo int, symbol string) error {
	if !b.cells[cellNo-1].isEmpty {
		return errors.New("cell already taken, choose another cell")
	}
	b.cells[cellNo-1] = Cell{symbol: symbol, isEmpty: false}
	return nil
}

// Display prints the current state of the board
func (b *Board) Display() {
	fmt.Println("Current Board:")
	for i := 0; i < 9; i += 3 {
		fmt.Printf("%s | %s | %s\n", b.cells[i].symbol, b.cells[i+1].symbol, b.cells[i+2].symbol)
		if i < 6 {
			fmt.Println("---------")
		}
	}
	fmt.Println()
}

// CheckWin checks if the current player has won the game
func (b *Board) CheckWin(symbol string) bool {
	winConditions := [][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Rows
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Columns
		{0, 4, 8}, {2, 4, 6},            // Diagonals
	}

	for _, condition := range winConditions {
		if b.cells[condition[0]].symbol == symbol &&
			b.cells[condition[1]].symbol == symbol &&
			b.cells[condition[2]].symbol == symbol {
			return true
		}
	}
	return false
}

// IsDraw checks if the game is a draw (no empty cells and no winner)
func (b *Board) IsDraw() bool {
	for _, cell := range b.cells {
		if cell.isEmpty {
			return false
		}
	}
	return true
}

// Cell struct represents a cell on the board
type Cell struct {
	symbol  string
	isEmpty bool
}

func main() {
	// Initialize a new TicTacToe game using the Game interface
	var g1 Game = NewTicTacToe("Player 1", "Player 2")

	// Simulate moves
	g1.Play(2)  // Player 1 ('X')
	g1.Play(5)  // Player 2 ('O')
	g1.Play(3)  // Player 1 ('X')
	g1.Play(1)  // Player 2 ('O')
	g1.Play(8)  // Player 1 ('X')
	g1.Play(4)  // Player 2 ('O')
	g1.Play(6)  // Player 1 ('X')
	g1.Play(7)  // Player 2 ('O')
	g1.Play(7)  // Player 2 ('O')
	g1.Play(6)  // Player 1 ('X')
	g1.Play(7)  // Player 2 ('O')
	g1.Play(7)  // Player 2 ('O')


	// Uncomment below to initialize another game on a new board
	// var g2 Game = NewTicTacToe("Player 3", "Player 4")
	// g2.Play(1) // Player 3 ('X')
}
