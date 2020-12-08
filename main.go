//TEST INTERVIEW YOOPY CHRISTIAN

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	counterKocokan = 1
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	inputPlayer := 3
	inputDice := 4

	playerDataScore := make([][]int, 0) // berisi data skor index [0] dan jumlah dadu pemain index[1]
	playerDataDice := make([][]int, 0)
	initScore(&playerDataScore, inputPlayer, inputDice)

	for {
		getDice(&playerDataDice, playerDataScore, inputPlayer)
		printResult(playerDataScore, playerDataDice)
		printEvaluatedResult(playerDataScore, playerDataDice)

		countPlayerLeft := 0
		for _, data := range playerDataScore { // hitung sisa player yang masih punya dadu
			if data[1] > 0 {
				countPlayerLeft++
			}
		}
		if countPlayerLeft == 1 { // jika player sisa satu, break
			break
		}

		time.Sleep(1000 * time.Millisecond)
	}

	for i, data := range playerDataScore {
		if data[1] > 0 {
			fmt.Println(`Game berakhir karena hanya pemain #`, i, `yang memiliki dadu`)
		}
	}
}

func printEvaluatedResult(dataScore [][]int, dataDice [][]int) {
	counterKocokan++
	totalTransferDiceNumberOne := make([]int, len(dataScore)) // catatan berapa angka satu yang dipindahkan
	// Data di proses
	for i := 0; i < len(dataScore); i++ { // i mewakili person
		// nge loop dadu nya person tersebut
		for j := 0; j < len(dataDice[i]); j++ {
			if dataDice[i][j] == 6 { // kalau dadu nya angka 6
				dataScore[i][0]++                                           // tambahkan skor orang tersebut
				dataScore[i][1]--                                           // kurangin dadu pemain tersebut
				dataDice[i] = append(dataDice[i][:j], dataDice[i][j+1:]...) // ilangkan dadu angka 1
				j--                                                         // abis apus elemen harus mundur
				continue                                                    // mulai dari atas lagi
			}

			if dataDice[i][j] == 1 {
				// kalau lebih dari 0 berarti angka satu nya dia boleh dapet dari pemain sebelumya
				// diginiin biar angka satu nya ga di oper2 terus selama loop
				if totalTransferDiceNumberOne[i] > 0 {
					totalTransferDiceNumberOne[i]-- // kurangin aja
					continue                        // ulang lagi dari awal
				}
				// kalau  totalTransferDiceNumberOne[i] sudah 0 baru proses kodingan di bawah

				dataScore[i][1]--                                           // kurangin jumlah dadu pemain tersebut
				dataDice[i] = append(dataDice[i][:j], dataDice[i][j+1:]...) // ilangin dadu angka 1 pemain tersebut

				if i == len(dataScore)-1 { // jika dadu angka satu nya ada di pemain terakhir
					if dataScore[0][1] == 0 { // jangan kasih dadu 1 ke pemain yang sudah selesai
						continue
					}
					dataScore[0][1]++                    // tambahkan jumlah dadu pemain pertama
					dataDice[0] = append(dataDice[0], 1) // tambahkan dadu ke pemain pertama
					continue
				}

				if dataScore[i+1][1] == 0 { // jangan kasih dadu 1 ke pemain yang sudah selesai
					continue
				}

				dataScore[i+1][1]++                      // tambahkan JUMLAH dadu orang sebelah
				dataDice[i+1] = append(dataDice[i+1], 1) // tambahkan DADU angka satu orang sebelah
				totalTransferDiceNumberOne[i+1]++        // tambahin datanya, biar tau berapa yang udah dipindahkan
			}
		}
	}

	// Print hasil proses data
	fmt.Println(`Setelah evaluasi:`)
	for i := 0; i < len(dataScore); i++ {
		if dataScore[i][1] == 0 {
			fmt.Println("Pemain #", i, "->", dataScore[i][0], ":_(Berhenti bermain karena tidak memiliki dadu)")
		} else {
			fmt.Println("Pemain #", i, "->", dataScore[i][0], ":", dataDice[i])
		}
	}
	fmt.Println(`===================================`)
}

func printResult(dataScore [][]int, dataDice [][]int) {

	fmt.Println(`Giliran`, counterKocokan, `lempar dadu :`)
	for i := 0; i < len(dataScore); i++ {
		if dataScore[i][1] == 0 {
			fmt.Println("Pemain #", i, "->", dataScore[i][0], ":_(Berhenti bermain karena tidak memiliki dadu)")
		} else {
			fmt.Println("Pemain #", i, "->", dataScore[i][0], ":", dataDice[i])
		}
	}
}

func getDice(playerDataDice *[][]int, dataScore [][]int, totalPlayer int) {
	*playerDataDice = [][]int{}

	for i := 0; i < totalPlayer; i++ {
		playerDice := make([]int, 0)

		for j := 0; j < dataScore[i][1]; j++ {
			playerDice = append(playerDice, rollDice())
		}

		*playerDataDice = append(*playerDataDice, playerDice)
	}
}

func initScore(playerData *[][]int, totalPlayer, totalDice int) {
	for i := 0; i < totalPlayer; i++ {
		temp := []int{0, totalDice}
		*playerData = append(*playerData, temp)
	}
}

func rollDice() int {
	dice := []int{1, 2, 3, 4, 5, 6}
	return dice[rand.Intn(len(dice))]
}
