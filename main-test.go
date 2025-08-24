// main_test.go
package main

import (
	"testing"
)

func TestBoardInitialization(t *testing.T) {
	board := Board{}
	board.InitializeBoard()

	// Test Kings
	piece := board.GetPiece(Position{Row: 0, Col: 4})
	if king, ok := piece.(King); !ok || king.GetColor() != White {
		t.Errorf("Expected White King at (1,5), got %+v", piece)
	}

	piece = board.GetPiece(Position{Row: 7, Col: 4})
	if king, ok := piece.(King); !ok || king.GetColor() != Black {
		t.Errorf("Expected Black King at (8,5), got %+v", piece)
	}

	// Test Queens
	piece = board.GetPiece(Position{Row: 0, Col: 3})
	if queen, ok := piece.(Queen); !ok || queen.GetColor() != White {
		t.Errorf("Expected White Queen at (1,4), got %+v", piece)
	}

	piece = board.GetPiece(Position{Row: 7, Col: 3})
	if queen, ok := piece.(Queen); !ok || queen.GetColor() != Black {
		t.Errorf("Expected Black Queen at (8,4), got %+v", piece)
	}

	// Test Pawns
	for col := 0; col < 8; col++ {
		piece = board.GetPiece(Position{Row: 1, Col: col})
		if _, ok := piece.(Pawn); !ok || piece.GetColor() != White {
			t.Errorf("Expected White Pawn at (2,%d), got %+v", col+1, piece)
		}
		piece = board.GetPiece(Position{Row: 6, Col: col})
		if _, ok := piece.(Pawn); !ok || piece.GetColor() != Black {
			t.Errorf("Expected Black Pawn at (7,%d), got %+v", col+1, piece)
		}
	}

	// Test Rooks
	piece = board.GetPiece(Position{Row: 0, Col: 0})
	if _, ok := piece.(Rook); !ok || piece.GetColor() != White {
		t.Errorf("Expected White Rook at (1,1), got %+v", piece)
	}
	piece = board.GetPiece(Position{Row: 0, Col: 7})
	if _, ok := piece.(Rook); !ok || piece.GetColor() != White {
		t.Errorf("Expected White Rook at (1,8), got %+v", piece)
	}
	piece = board.GetPiece(Position{Row: 7, Col: 0})
	if _, ok := piece.(Rook); !ok || piece.GetColor() != Black {
		t.Errorf("Expected Black Rook at (8,1), got %+v", piece)
	}
	piece = board.GetPiece(Position{Row: 7, Col: 7})
	if _, ok := piece.(Rook); !ok || piece.GetColor() != Black {
		t.Errorf("Expected Black Rook at (8,8), got %+v", piece)
	}
}

func TestValidPawnMove(t *testing.T) {
	game := NewChessGame()
	start := Position{Row: 1, Col: 0} // White Pawn at a2
	end := Position{Row: 3, Col: 0}   // Move to a4
	if !game.IsValidMove(start, end) {
		t.Errorf("Expected move from a2 to a4 to be valid.")
	}
}

func TestInvalidPawnMove(t *testing.T) {
	game := NewChessGame()
	start := Position{Row: 1, Col: 0} // White Pawn at a2
	end := Position{Row: 2, Col: 1}   // Attempting to move to b3 without capture
	if game.IsValidMove(start, end) {
		t.Errorf("Expected move from a2 to b3 to be invalid.")
	}
}

func TestCaptureMove(t *testing.T) {
	game := NewChessGame()
	// Move white pawn from a2 to a4
	game.board.MovePiece(Position{Row: 1, Col: 0}, Position{Row: 3, Col: 0})
	// Move black pawn from b7 to b5
	game.board.MovePiece(Position{Row: 6, Col: 1}, Position{Row: 4, Col: 1})
	// White pawn captures black pawn from a4 to b5
	start := Position{Row: 3, Col: 0}
	end := Position{Row: 4, Col: 1}
	if !game.IsValidMove(start, end) {
		t.Errorf("Expected capture move from a4 to b5 to be valid.")
	}
}

func TestKingCapture(t *testing.T) {
	game := NewChessGame()
	// Manually place the black king in a vulnerable position
	game.board.SetPiece(Position{Row: 4, Col: 4}, King{White})
	game.board.SetPiece(Position{Row: 4, Col: 5}, King{Black})

	// White moves King from e1 to f2
	start := Position{Row: 0, Col: 4}
	end := Position{Row: 1, Col: 5}
	if !game.IsValidMove(start, end) {
		t.Errorf("Expected move from e1 to f2 to be valid.")
	}
	game.board.MovePiece(start, end)
	game.CheckGameOver()
	if !game.gameOver {
		t.Errorf("Expected game to be over after capturing the black king.")
	}
}

func TestSwitchPlayer(t *testing.T) {
	game := NewChessGame()
	initialPlayer := game.currentPlayer
	game.SwitchPlayer()
	if game.currentPlayer == initialPlayer {
		t.Errorf("Expected player to switch from %s", initialPlayer)
	}
}

func TestAlgebraicToPosition(t *testing.T) {
	game := NewChessGame()
	pos, err := game.AlgebraicToPosition("e2")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := Position{Row: 1, Col: 4}
	if pos != expected {
		t.Errorf("Expected e2 to be %+v, got %+v", expected, pos)
	}
}

func TestInvalidInputFormat(t *testing.T) {
	game := NewChessGame()
	_, _, err := game.ParseInput("invalid_input")
	if err == nil {
		t.Errorf("Expected error for invalid input format.")
	}
}
