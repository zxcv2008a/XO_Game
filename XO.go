package main

import (
	"fmt"
)

func main() {

	xoBoard := [3][3]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}}
	Player := [2]string{"X", "O"}

	var row int
	var column int
	var play int
	play = 0

	for i := 1; i <= 9; i++ {
		fmt.Print(Player[play])
		fmt.Println(" , Please Enter column number")
		fmt.Scanln(&column)
		fmt.Print(Player[play])
		fmt.Println(" , Please Enter row number")
		fmt.Scanln(&row)

		if xoBoard[column-1][row-1] != "-" {
			fmt.Println("Illegal Move Try again")
			i--
			continue
		} else {

			xoBoard[column-1][row-1] = Player[play]

			fmt.Println(xoBoard[0])
			fmt.Println(xoBoard[1])
			fmt.Println(xoBoard[2])

			if xoBoard[0][0] == Player[play] && xoBoard[0][1] == Player[play] && xoBoard[0][2] == Player[play] {
				fmt.Println("The Winner is ", Player[play])
				return
			} else if xoBoard[1][0] == Player[play] && xoBoard[1][1] == Player[play] && xoBoard[1][2] == Player[play] {
				fmt.Println("The Winner is ", Player[play])
				return
			} else if xoBoard[2][0] == Player[play] && xoBoard[2][1] == Player[play] && xoBoard[2][2] == Player[play] {
				fmt.Println("The Winner is ", Player[play])
				return
			} else if xoBoard[0][0] == Player[play] && xoBoard[1][0] == Player[play] && xoBoard[2][0] == Player[play] {
				fmt.Println("The Winner is ", Player[play])
				return
			} else if xoBoard[0][1] == Player[play] && xoBoard[1][1] == Player[play] && xoBoard[2][1] == Player[play] {
				fmt.Println("The Winner is ", Player[play])
				return
			} else if xoBoard[0][2] == Player[play] && xoBoard[1][2] == Player[play] && xoBoard[2][2] == Player[play] {
				fmt.Println("The Winner is ", Player[play])
				return
			}

			if play == 1 {
				play = 0
			} else {
				play = 1
			}
		}

	}
	fmt.Println("Stale Mate, Game Over")
}
