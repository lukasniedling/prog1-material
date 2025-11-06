package main



import "fmt"



func main() {

    var board [3][3]string

    currentPlayer := "X"

   

    for {

        printBoard(board)

        fmt.Printf("Spieler %s ist am Zug.\n", currentPlayer)

        fmt.Print("Gib die Zeile (1-3) ein: ")

        var row int

        fmt.Scanln(&row)

        row = row - 1  // Konvertiere 1-3 zu 0-2

       

        fmt.Print("Gib die Spalte (1-3) ein: ")

        var col int

        fmt.Scanln(&col)

        col = col - 1  // Konvertiere 1-3 zu 0-2

       



        // Überprüfe ob die Position gültig und frei ist

        if row >= 0 && row < 3 && col >= 0 && col < 3 && board[row][col] == "" {

            board[row][col] = currentPlayer

            // Wechsel den Spieler

            if currentPlayer == "X" {

                currentPlayer = "O"

            } else {

                currentPlayer = "X"

            }

        } else {

            fmt.Println("Ungültige Position! Bitte erneut versuchen.")

        }

        if checkWin(board, "X") {

            printBoard(board)

            fmt.Println("Spieler X hat gewonnen!")

            break

        } else if checkWin(board, "O") {

            printBoard(board)

            fmt.Println("Spieler O hat gewonnen!")

            break

        }

    }

}



func printBoard(board [3][3]string) {

    for _, row := range board {

        for _, cell := range row {

            if cell == "" {

                fmt.Print("_ ")

            } else {

                fmt.Print(cell + " ")

            }

        }

        fmt.Println()

    }

}



func checkWin(board [3][3]string, player string) bool {

    // Überprüfe Reihen und Spalten

    for i := 0; i < 3; i++ {

        if (board[i][0] == player && board[i][1] == player && board[i][2] == player) ||

            (board[0][i] == player && board[1][i] == player && board[2][i] == player) {

            return true

        }

    }

    // Überprüfe Diagonalen

    if (board[0][0] == player && board[1][1] == player && board[2][2] == player) ||

        (board[0][2] == player && board[1][1] == player && board[2][0] == player) {

        return true

    }  

    return false

}