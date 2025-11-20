package chesspiece

/* Ziel: Entwurf eines Datentyps für Schachfiguren.
Es soll eine Datenstruktur für Schachfiguren entworfen werden,
die eine Methode enthält, mit der geprüft werden kann,
ob ein bestimmter Zug für diese Figur erlaubt ist.

DISCLAIMER:
Die Aufgabe soll den Umgang mit Structs und Consts/Enums üben.
Es geht nicht darum, Schachzüge akkurat in allen Details umzusetzen
(z.B. keine Rochade oder andere komplexere Bedingungen).
*/

type PType int

const (
	BISHOP PType = iota
	KNIGHT
	QUEEN
	KING
	ROOK
	PAWN
)

type Colour int

const (
	WHITE Colour = iota
	BLACK
)

// ChessPiece repräsentiert eine Schachfigur auf einem Spielfeld.
type ChessPiece struct {
	pieceType PType
	colour    Colour
	row       int
	column    int
}

// Methoden: TODO

// MoveAllowed erwartet eine Feld-Angabe und liefert true,
// falls die Figur nach den Bewegungsregeln beim Schach
// auf dieses Feld ziehen darf.
// Besonderheiten wie Rochade oder im Weg stehende Figuren
// Spielen keine Rolle.
func (p ChessPiece) MoveAllowed(row, col int) bool {

	// switch p.pieceType {
	// case BISHOP: ...
	// case KNIGHT: ...
	// }
	if row < 7 && row >= 0 && col < 7 && col >= 0 {

		if p.pieceType == BISHOP {
			// Diagonale von links unten nach rechts oben.
			if row-col == p.row-p.column {
				return true
			}
			// Diagonale von links oben nach rechts unten.
			if row+col == p.row+p.column {
				return true
			}
		}
		if p.pieceType == KNIGHT {
			// "L"-förmige Bewegungen
			if (((row-p.row) == 2 || (row-p.row) == -2) && ((col-p.column) == 1 || (col-p.column) == -1)) ||
				(((row-p.row) == 1 || (row-p.row) == -1) && ((col-p.column) == 2 || (col-p.column) == -2)) {
				return true
			}
		}

		if p.pieceType == ROOK {
			// gleiche Reihe oder gleiche Spalte
			if row == p.row || col == p.column {
				return true
			}
		}

		if p.pieceType == PAWN {
			// Ein Feld vorwärts (ohne Berücksichtigung von Anfangs-Doppelschritt oder Schlagen)
			if p.colour == WHITE && row == p.row-1 && col == p.column {
				return true
			}
			if p.colour == BLACK && row == p.row+1 && col == p.column {
				return true
			}
		}

		if p.pieceType == KING {
			// Ein Feld in jede Richtung.
			// Prüft, ob der Reihen-Unterschied zwischen -1 und 1 liegt.
			if (row-p.row) >= -1 && (row-p.row) <= 1 {
				// UND ob der Spalten-Unterschied zwischen -1 und 1 liegt.
				if (col-p.column) >= -1 && (col-p.column) <= 1 {
					return true
				}
			}
		}

		if p.pieceType == QUEEN {
			// Kombination aus Turm und Läufer
			// Diagonale von links unten nach rechts oben.
			if row-col == p.row-p.column {
				return true
			}
			// Diagonale von links oben nach rechts unten.
			if row+col == p.row+p.column {
				return true
			}
			// gleiche Reihe oder gleiche Spalte
			if row == p.row || col == p.column {
				return true
			}
		}

		
	}
	return false
}
