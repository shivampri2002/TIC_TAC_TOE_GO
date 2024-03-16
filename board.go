package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//defining the board for the game
type board struct {
	pieces   [3][3]uint8
	turn     uint8
	finished bool
}

//checking for the any winner in the game
func (b *board) result() uint8 {
	//checking if the top left to bottom right diagonal has any winner
	if b.pieces[0][0] != 0 && b.pieces[0][0] == b.pieces[1][1] && b.pieces[1][1] == b.pieces[2][2] {
		return b.pieces[0][0]
	}

	//checking if the top right to bottom left diagonal has any winner
	if b.pieces[0][2] != 0 && b.pieces[0][2] == b.pieces[1][1] && b.pieces[1][1] == b.pieces[2][0] {
		return b.pieces[0][2]
	}

	//checking for the vertical and horizontal line have any winner
	for i := 0; i < 3; i++ {
		if b.pieces[i][0] != 0 && b.pieces[i][0] == b.pieces[i][1] && b.pieces[i][1] == b.pieces[i][2] {
			return b.pieces[i][0]
		}

		if b.pieces[0][i] != 0 && b.pieces[0][i] == b.pieces[1][i] && b.pieces[1][i] == b.pieces[2][i] {
			return b.pieces[0][i]
		}
	}

	return 0
}

func (b *board) newClick(row, column uint8) {
	b.pieces[row][column] = b.turn % 2 + 1;

	if b.turn > 3 {
		winner := b.result()

		if winner == 0 {
			if b.turn == 8 {
				dialog.ShowInformation("It is a tie!", "Nobody has won. Better luck next time.", fyne.CurrentApp().Driver().AllWindows()[0])
				b.finished = true
			}
			return
		}

		number := string(winner + 48)

		dialog.ShowInformation("Player "+number+" has won!", "Congratulations to player "+number+" for winning.", fyne.CurrentApp().Driver().AllWindows()[0])
		b.finished = true
	}
}

func (b *board) Reset() {
	for i := range b.pieces {
		b.pieces[i][0] = 0
		b.pieces[i][1] = 0
		b.pieces[i][2] = 0
	}

	b.finished = false
	b.turn = 0
}


//defining the boardIcon struct
type boardIcon struct {
	widget.Icon
	board *board
	row, column uint8
}

func (i *boardIcon) Reset() {
	i.SetResource(theme.ViewFullScreenIcon())
}

func (i *boardIcon) Tapped(ev *fyne.PointEvent) {
	if i.board.pieces[i.row][i.column] != 0 || i.board.finished {
		return
	}

	if i.board.turn % 2 == 0 {
		i.SetResource(theme.RadioButtonIcon())
	} else {
		i.SetResource(theme.CancelIcon())
	}

	i.board.newClick(i.row, i.column)
	i.board.turn++
}

func newBoardIcon(row, column uint8, b *board) *boardIcon {
	bdIcon := &boardIcon{row: row, column: column, board: b}
	bdIcon.SetResource(theme.ViewFullScreenIcon())
	bdIcon.ExtendBaseWidget(bdIcon)
	return bdIcon
}