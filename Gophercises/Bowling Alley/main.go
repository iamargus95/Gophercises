package main

import "fmt"

const (
	Strike  = 10
	Spare   = 10
	LastTry = 9
)

func main() {
	fmt.Println(bowlingScores(player1Frame))
}

var player1Frame = [][]int{{10, 0}, {3, 3}, {3, 3}, {3, 3}, {8, 2}, {2, 8}, {7, 3}, {3, 7}, {5, 5}, {10, 0}}

func bowlingScores(frames [][]int) int {

	score := 0

	for try, frame := range frames {

		if try == LastTry {
			for j := range frame {
				score += frame[j]
			}
		} else {

			if isStrike(frame) {

				score += frame[0]
				score += frames[try+1][0]

				if try < 8 {
					if isStrike(frames[try+1]) {
						score += frames[try+2][0]
					} else {
						score += frames[try+1][1]
					}
				} else if try == 8 {
					score += frames[try+1][1]
				}
			}

			if frame[0]+frame[1] != Spare {
				score += frame[0]
				score += frame[1]
			}

			if isSpare(frame) && !isStrike(frame) {
				score += frame[0]
				score += frame[1]
				score += frames[try+1][0]
			}
		}
	}
	return score
}

func isStrike(frame []int) bool {
	return frame[0] == 10
}

func isSpare(frame []int) bool {
	if frame[0] != 10 && frame[0]+frame[1] == 10 {
		return true
	}
	return false
}
