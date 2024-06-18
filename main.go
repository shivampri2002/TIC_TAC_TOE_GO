package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.New()

	w := myApp.NewWindow("Tic Tac Toe Game")
	w.SetContent(container.NewStack(show(w)))
	w.Resize(fyne.NewSize(960, 540))

	w.ShowAndRun()
}
