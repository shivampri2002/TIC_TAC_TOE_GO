package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"

	"fyne.io/fyne/v2/theme"
)

func randomMoveGenerate(b *board) [2]int {
	emptyCell := make([][2]int, 0, 9)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.pieces[i][j] == 0 {
				emptyCell = append(emptyCell, [2]int{i, j})
			}
		}
	}

	return emptyCell[rand.Intn(9)]
}

func MultiPlayerMoveController(b *board) {
	b.turn = 1

	for {
		if b.tcpClient == nil {
			//show the dialog of disconnection and move to practice mode
			log.Println("Connection with the server disconnected!")
			return
		}

		message, err := b.tcpClient.Read()
		if err != nil {
			log.Println("Error reading from server:", err)
			b.tcpClient.Close()
			return
		}

		mesArr := strings.Split(strings.TrimSpace(message), " ")
		log.Println(mesArr)

		if mesArr[0] == "GAME_OVER" {
			if mesArr[1] == "W" {
				log.Println("You have won the Game.")
				//show the dialog with animation
			} else if mesArr[1] == "L" {
				log.Println("You have lost the Game.")
				//show the dialog
			} else if mesArr[1] == "T" {
				log.Println("The Game has been tied.!")
				//show the dialog
			}

			b.tcpClient.Close()
			log.Println("Game over")
			return
		} else if mesArr[0] == "MOVE" {

			x, _ := strconv.Atoi(mesArr[1])
			y, _ := strconv.Atoi(mesArr[2])

			if x != 10 || y != 10 {
				b.grid.Objects[(x*3 + y)].(*boardIcon).SetResource(theme.RadioButtonIcon())
				b.pieces[x][y] = 2
				b.turn += 1

				b.multiPlayLabStr.Set(fmt.Sprintf("Opponent moved %d, %d. Your Turn!", x, y))
			}

			if x == 10 {
				b.multiPlayLabStr.Set("Start the Game. Your Turn !")
				b.turn = 0
			}
			//before sending the random move
			// b.mutx.Lock()
			// defer b.mutx.Unlock()
		} else if mesArr[0] == "WAIT" {
			b.multiPlayLabStr.Set("Waiting for the Opponent to Join!")
		} else {
			b.multiPlayLabStr.Set(message)
		}

		// log.Println("Received from server:", message)
	}
}
