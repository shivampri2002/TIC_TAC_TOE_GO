package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Show loads a tic-tac-toe example window for the specified app context
func show(win fyne.Window) fyne.CanvasObject {
	board := &board{}

	board.grid = container.NewGridWithColumns(3)

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			addNewBoardIcon(uint8(r), uint8(c), board)
		}
	}

	win.SetOnClosed(func() {
		if board.tcpClient != nil {
			board.tcpClient.Close()
		}
	})

	//two buttons for the Practice and Multiplayer
	board.multiPlayerMode = false

	board.multiPlayLabStr = binding.NewString()
	board.multiPlayLabStr.Set("Initial value")

	reset := widget.NewButtonWithIcon("Reset Board", theme.ViewRefreshIcon(), func() {
		board.Reset()
	})

	multiPlayLabel := widget.NewLabelWithData(board.multiPlayLabStr)

	reslabInnerContainer := container.NewStack(reset)
	resetLabOuterContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), reslabInnerContainer, layout.NewSpacer())

	updateResetLabContent := func() {
		reslabInnerContainer.Objects = nil // Clear current content
		if board.multiPlayerMode {
			reslabInnerContainer.Add(reset)
		} else {
			reslabInnerContainer.Add(multiPlayLabel)
		}
		reslabInnerContainer.Refresh()
		resetLabOuterContainer.Refresh()
		log.Println(true)
		log.Println(reslabInnerContainer)
	}

	multiPlayerBtn := widget.NewButton("MultiPlayer", func() {
		if board.multiPlayerMode {
			return
		}
		board.multiPlayerMode = true
		updateResetLabContent()

		board.Reset()

		//making the tcp connection
		err := board.tcpClient.Connect("localhost:3000")
		if err != nil {
			board.multiPlayerMode = false
			updateResetLabContent()
			log.Println("Failed to connect to server:", err)
			dialog.ShowInformation("Error: ", "Failed to connect to server!", fyne.CurrentApp().Driver().AllWindows()[0])
			return
		}

		go MultiPlayerMoveController(board)
	})

	practiceGameBtn := widget.NewButton("Practice", func() {
		if !board.multiPlayerMode {
			return
		}
		board.multiPlayerMode = false
		if board.tcpClient != nil {
			board.tcpClient.Close()
		}
		updateResetLabContent()

		board.Reset()
	})

	playingModeBtnContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), multiPlayerBtn, practiceGameBtn, layout.NewSpacer())

	gameControlContainer := container.New(layout.NewVBoxLayout(), playingModeBtnContainer, resetLabOuterContainer)

	return container.NewBorder(gameControlContainer, nil, nil, nil, board.grid)
}
