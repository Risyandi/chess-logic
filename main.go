// main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Color represents the color of a chess piece.
type Color string

const (
	White Color = "white"
	Black Color = "black"
)

// Piece is the interface that all chess pieces implement.
type Piece interface {
	GetName() string
	GetColor() Color
	GetSymbol() string
	GetValidMoves(position Position, board *Board) []Position
}

// Position represents a position on the chessboard.
type Position struct {
	Row int
	Col int
}

// King struct
type King struct {
	color Color
}

func (k King) GetName() string {
	return "King"
}

func (k King) GetColor() Color {
	return k.color
}

func (k King) GetSymbol() string {
	if k.color == White {
		return "K"
	}
	return "k"
}

func (k King) GetValidMoves(position Position, board *Board) []Position {
	directions := []Position{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1} /* King */, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	var moves []Position
	for _, d := range directions {
		newRow := position.Row + d.Row
		newCol := position.Col + d.Col
		if board.IsWithinBounds(newRow, newCol) {
			target := board.GetPiece(Position{newRow, newCol})
			if target == nil || target.GetColor() != k.color {
				moves = append(moves, Position{newRow, newCol})
			}
		}
	}
	return moves
}

// Queen struct
type Queen struct {
	color Color
}

func (q Queen) GetName() string {
	return "Queen"
}

func (q Queen) GetColor() Color {
	return q.color
}

func (q Queen) GetSymbol() string {
	if q.color == White {
		return "Q"
	}
	return "q"
}

func (q Queen) GetValidMoves(position Position, board *Board) []Position {
	directions := []Position{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	var moves []Position
	for _, d := range directions {
		newRow, newCol := position.Row, position.Col
		for {
			newRow += d.Row
			newCol += d.Col
			if !board.IsWithinBounds(newRow, newCol) {
				break
			}
			target := board.GetPiece(Position{newRow, newCol})
			if target == nil {
				moves = append(moves, Position{newRow, newCol})
			} else {
				if target.GetColor() != q.color {
					moves = append(moves, Position{newRow, newCol})
				}
				break
			}
		}
	}
	return moves
}

// Rook struct
type Rook struct {
	color Color
}

func (r Rook) GetName() string {
	return "Rook"
}

func (r Rook) GetColor() Color {
	return r.color
}

func (r Rook) GetSymbol() string {
	if r.color == White {
		return "R"
	}
	return "r"
}

func (r Rook) GetValidMoves(position Position, board *Board) []Position {
	directions := []Position{
		{-1, 0}, {0, -1}, {0, 1}, {1, 0},
	}
	var moves []Position
	for _, d := range directions {
		newRow, newCol := position.Row, position.Col
		for {
			newRow += d.Row
			newCol += d.Col
			if !board.IsWithinBounds(newRow, newCol) {
				break
			}
			target := board.GetPiece(Position{newRow, newCol})
			if target == nil {
				moves = append(moves, Position{newRow, newCol})
			} else {
				if target.GetColor() != r.color {
					moves = append(moves, Position{newRow, newCol})
				}
				break
			}
		}
	}
	return moves
}

// Bishop struct
type Bishop struct {
	color Color
}

func (b Bishop) GetName() string {
	return "Bishop"
}

func (b Bishop) GetColor() Color {
	return b.color
}

func (b Bishop) GetSymbol() string {
	if b.color == White {
		return "B"
	}
	return "b"
}

func (b Bishop) GetValidMoves(position Position, board *Board) []Position {
	directions := []Position{
		{-1, -1}, {-1, 1},
		{1, -1}, {1, 1},
	}
	var moves []Position
	for _, d := range directions {
		newRow, newCol := position.Row, position.Col
		for {
			newRow += d.Row
			newCol += d.Col
			if !board.IsWithinBounds(newRow, newCol) {
				break
			}
			target := board.GetPiece(Position{newRow, newCol})
			if target == nil {
				moves = append(moves, Position{newRow, newCol})
			} else {
				if target.GetColor() != b.color {
					moves = append(moves, Position{newRow, newCol})
				}
				break
			}
		}
	}
	return moves
}

// Knight struct
type Knight struct {
	color Color
}

func (n Knight) GetName() string {
	return "Knight"
}

func (n Knight) GetColor() Color {
	return n.color
}

func (n Knight) GetSymbol() string {
	if n.color == White {
		return "N"
	}
	return "n"
}

func (n Knight) GetValidMoves(position Position, board *Board) []Position {
	deltas := []Position{
		{-2, -1}, {-2, 1},
		{-1, -2}, {-1, 2},
		{1, -2}, {1, 2},
		{2, -1}, {2, 1},
	}
	var moves []Position
	for _, d := range deltas {
		newRow := position.Row + d.Row
		newCol := position.Col + d.Col
		if board.IsWithinBounds(newRow, newCol) {
			target := board.GetPiece(Position{newRow, newCol})
			if target == nil || target.GetColor() != n.color {
				moves = append(moves, Position{newRow, newCol})
			}
		}
	}
	return moves
}

// Pawn struct
type Pawn struct {
	color Color
}

func (p Pawn) GetName() string {
	return "Pawn"
}

func (p Pawn) GetColor() Color {
	return p.color
}

func (p Pawn) GetSymbol() string {
	if p.color == White {
		return "P"
	}
	return "p"
}

func (p Pawn) GetValidMoves(position Position, board *Board) []Position {
	var moves []Position
	direction := 1
	startRow := 1
	if p.color == Black {
		direction = -1
		startRow = 6
	}
	// Move forward
	newRow := position.Row + direction
	if board.IsWithinBounds(newRow, position.Col) && board.GetPiece(Position{newRow, position.Col}) == nil {
		moves = append(moves, Position{newRow, position.Col})
		// Double move from starting position
		if position.Row == startRow {
			newRow2 := newRow + direction
			if board.GetPiece(Position{newRow2, position.Col}) == nil {
				moves = append(moves, Position{newRow2, position.Col})
			}
		}
	}
	// Captures
	for _, dc := range []int{-1, 1} {
		newCol := position.Col + dc
		if board.IsWithinBounds(newRow, newCol) {
			target := board.GetPiece(Position{newRow, newCol})
			if target != nil && target.GetColor() != p.color {
				moves = append(moves, Position{newRow, newCol})
			}
		}
	}
	return moves
}

// Board struct
type Board struct {
	grid [8][8]Piece
}

// InitializeBoard sets up the chess board with pieces.
func (b *Board) InitializeBoard() {
	// Place Pawns
	for col := 0; col < 8; col++ {
		b.grid[1][col] = Pawn{White}
		b.grid[6][col] = Pawn{Black}
	}
	// Place Rooks
	b.grid[0][0] = Rook{White}
	b.grid[0][7] = Rook{White}
	b.grid[7][0] = Rook{Black}
	b.grid[7][7] = Rook{Black}
	// Place Knights
	b.grid[0][1] = Knight{White}
	b.grid[0][6] = Knight{White}
	b.grid[7][1] = Knight{Black}
	b.grid[7][6] = Knight{Black}
	// Place Bishops
	b.grid[0][2] = Bishop{White}
	b.grid[0][5] = Bishop{White}
	b.grid[7][2] = Bishop{Black}
	b.grid[7][5] = Bishop{Black}
	// Place Queens
	b.grid[0][3] = Queen{White}
	b.grid[7][3] = Queen{Black}
	// Place Kings
	b.grid[0][4] = King{White}
	b.grid[7][4] = King{Black}
}

// PrintBoard displays the current state of the board.
func (b *Board) PrintBoard() {
	fmt.Println("  a b c d e f g h")
	for row := 7; row >= 0; row-- {
		fmt.Print(row+1, " ")
		for col := 0; col < 8; col++ {
			piece := b.grid[row][col]
			if piece == nil {
				fmt.Print(". ")
			} else {
				fmt.Print(piece.GetSymbol(), " ")
			}
		}
		fmt.Println(row + 1)
	}
	fmt.Println("  a b c d e f g h")
}

// GetPiece retrieves the piece at a given position.
func (b *Board) GetPiece(pos Position) Piece {
	if !b.IsWithinBounds(pos.Row, pos.Col) {
		return nil
	}
	return b.grid[pos.Row][pos.Col]
}

// SetPiece places a piece at a given position.
func (b *Board) SetPiece(pos Position, piece Piece) {
	if !b.IsWithinBounds(pos.Row, pos.Col) {
		return
	}
	b.grid[pos.Row][pos.Col] = piece
}

// MovePiece moves a piece from start to end position. Returns any captured piece.
func (b *Board) MovePiece(start, end Position) Piece {
	piece := b.GetPiece(start)
	if piece == nil {
		return nil
	}
	captured := b.GetPiece(end)
	b.SetPiece(end, piece)
	b.SetPiece(start, nil)
	return captured
}

// IsWithinBounds checks if the given row and column are within the board.
func (b *Board) IsWithinBounds(row, col int) bool {
	return row >= 0 && row < 8 && col >= 0 && col < 8
}

// ChessGame struct
type ChessGame struct {
	board         Board
	currentPlayer Color
	gameOver      bool
}

// NewChessGame initializes a new chess game.
func NewChessGame() *ChessGame {
	board := Board{}
	board.InitializeBoard()
	return &ChessGame{
		board:         board,
		currentPlayer: White,
		gameOver:      false,
	}
}

// SwitchPlayer changes the current player.
func (cg *ChessGame) SwitchPlayer() {
	if cg.currentPlayer == White {
		cg.currentPlayer = Black
	} else {
		cg.currentPlayer = White
	}
}

// ParseInput parses user input into start and end positions.
func (cg *ChessGame) ParseInput(input string) (Position, Position, error) {
	input = strings.TrimSpace(input)
	var start, end Position
	if strings.Contains(input, ",") {
		// Coordinate notation: e.g., 2,4 4,4
		parts := strings.Fields(input)
		if len(parts) != 2 {
			return start, end, fmt.Errorf("invalid input format")
		}
		startParts := strings.Split(parts[0], ",")
		endParts := strings.Split(parts[1], ",")
		if len(startParts) != 2 || len(endParts) != 2 {
			return start, end, fmt.Errorf("invalid input format")
		}
		startRow, err1 := strconv.Atoi(startParts[0])
		startCol, err2 := strconv.Atoi(startParts[1])
		endRow, err3 := strconv.Atoi(endParts[0])
		endCol, err4 := strconv.Atoi(endParts[1])
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			return start, end, fmt.Errorf("invalid numbers in input")
		}
		start = Position{Row: startRow - 1, Col: startCol - 1}
		end = Position{Row: endRow - 1, Col: endCol - 1}
	} else if strings.Contains(input, "to") {
		// Algebraic notation: e.g., e2 to e4
		parts := strings.Split(input, "to")
		if len(parts) != 2 {
			return start, end, fmt.Errorf("invalid input format")
		}
		startPos, err1 := cg.AlgebraicToPosition(strings.TrimSpace(parts[0]))
		endPos, err2 := cg.AlgebraicToPosition(strings.TrimSpace(parts[1]))
		if err1 != nil || err2 != nil {
			return start, end, fmt.Errorf("invalid algebraic positions")
		}
		start = startPos
		end = endPos
	} else {
		// Algebraic notation with comma: e.g., e2,e4
		parts := strings.Split(input, ",")
		if len(parts) != 2 {
			return start, end, fmt.Errorf("invalid input format")
		}
		startPos, err1 := cg.AlgebraicToPosition(strings.TrimSpace(parts[0]))
		endPos, err2 := cg.AlgebraicToPosition(strings.TrimSpace(parts[1]))
		if err1 != nil || err2 != nil {
			return start, end, fmt.Errorf("invalid algebraic positions")
		}
		start = startPos
		end = endPos
	}
	return start, end, nil
}

// AlgebraicToPosition converts algebraic notation to Position.
func (cg *ChessGame) AlgebraicToPosition(alg string) (Position, error) {
	if len(alg) != 2 {
		return Position{}, fmt.Errorf("invalid algebraic notation")
	}
	colChar := strings.ToLower(string(alg[0]))
	rowChar := string(alg[1])
	columns := "abcdefgh"
	col := strings.Index(columns, colChar)
	if col == -1 {
		return Position{}, fmt.Errorf("invalid column in algebraic notation")
	}
	row, err := strconv.Atoi(rowChar)
	if err != nil || row < 1 || row > 8 {
		return Position{}, fmt.Errorf("invalid row in algebraic notation")
	}
	return Position{Row: row - 1, Col: col}, nil
}

// PositionToAlgebraic converts Position to algebraic notation.
func (cg *ChessGame) PositionToAlgebraic(pos Position) string {
	columns := "abcdefgh"
	if pos.Col < 0 || pos.Col >= 8 || pos.Row < 0 || pos.Row >= 8 {
		return "Invalid"
	}
	return fmt.Sprintf("%c%d", columns[pos.Col], pos.Row+1)
}

// IsValidMove checks if moving from start to end is a valid move.
func (cg *ChessGame) IsValidMove(start, end Position) bool {
	piece := cg.board.GetPiece(start)
	if piece == nil {
		fmt.Println("No piece at the starting position.")
		return false
	}
	if piece.GetColor() != cg.currentPlayer {
		fmt.Printf("It's %s's turn.\n", cg.currentPlayer)
		return false
	}
	validMoves := piece.GetValidMoves(start, &cg.board)
	for _, move := range validMoves {
		if move == end {
			return true
		}
	}
	fmt.Println("Invalid move for the selected piece.")
	return false
}

// CheckGameOver checks if either king has been captured.
func (cg *ChessGame) CheckGameOver() {
	whiteKing := false
	blackKing := false
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			piece := cg.board.GetPiece(Position{Row: row, Col: col})
			if piece != nil {
				if king, ok := piece.(King); ok {
					if king.GetColor() == White {
						whiteKing = true
					} else {
						blackKing = true
					}
				}
			}
		}
	}
	if !whiteKing {
		fmt.Println("Black wins! White king has been captured.")
		cg.gameOver = true
	} else if !blackKing {
		fmt.Println("White wins! Black king has been captured.")
		cg.gameOver = true
	}
}

// PlayTurn handles a single turn of the game.
func (cg *ChessGame) PlayTurn(scanner *bufio.Scanner) {
	cg.board.PrintBoard()
	fmt.Printf("%s's move:\n", strings.Title(string(cg.currentPlayer)))
	fmt.Print("Enter your move (e.g., e2,e4 or 2,4 to 4,4): ")
	if !scanner.Scan() {
		fmt.Println("Error reading input.")
		cg.gameOver = true
		return
	}
	input := scanner.Text()
	start, end, err := cg.ParseInput(input)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}
	if !cg.IsValidMove(start, end) {
		return
	}
	captured := cg.board.MovePiece(start, end)
	fmt.Printf("Moved from %s to %s\n", cg.PositionToAlgebraic(start), cg.PositionToAlgebraic(end))
	if captured != nil {
		fmt.Printf("Captured %s at %s\n", captured.GetName(), cg.PositionToAlgebraic(end))
	}
	cg.CheckGameOver()
	if !cg.gameOver {
		cg.SwitchPlayer()
	}
}

// StartGame begins the chess game loop.
func (cg *ChessGame) StartGame() {
	fmt.Println("Welcome to Console Chess in Go!")
	scanner := bufio.NewScanner(os.Stdin)
	for !cg.gameOver {
		cg.PlayTurn(scanner)
	}
	fmt.Println("Game Over.")
}

func main() {
	game := NewChessGame()
	game.StartGame()
}
