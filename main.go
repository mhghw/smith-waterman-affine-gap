package main

import (
	"fmt"
)

func SmithWaterman(seq1, seq2 string, match, mismatch, gapOpen, gapExtend int) (int, string, string) {
	// Initialize score matrix
	score := make([][]int, len(seq1)+1)
	for i := range score {
		score[i] = make([]int, len(seq2)+1)
	}

	// Initialize traceback matrix
	traceback := make([][]string, len(seq1)+1)
	for i := range traceback {
		traceback[i] = make([]string, len(seq2)+1)
	}

	// Fill the matrices
	maxScore := 0
	maxI, maxJ := 0, 0
	for i := 1; i <= len(seq1); i++ {
		for j := 1; j <= len(seq2); j++ {
			diagonalScore := score[i-1][j-1] + match
			if seq1[i-1] != seq2[j-1] {
				diagonalScore = score[i-1][j-1] + mismatch
			}
			gapPenalty1 := gapOpen
			if traceback[i-1][j] != "diagonal" {
				gapPenalty1 = gapExtend
			}
			gapScore1 := score[i-1][j] + gapPenalty1
			gapPenalty2 := gapOpen
			if traceback[i][j-1] != "diagonal" {
				gapPenalty2 = gapExtend
			}
			gapScore2 := score[i][j-1] + gapPenalty2

			// Determine the score for the current cell
			score[i][j] = 0
			traceback[i][j] = ""
			if diagonalScore > score[i][j] {
				score[i][j] = diagonalScore
				traceback[i][j] = "diagonal"
			}
			if gapScore1 > score[i][j] {
				score[i][j] = gapScore1
				traceback[i][j] = "up"
			}
			if gapScore2 > score[i][j] {
				score[i][j] = gapScore2
				traceback[i][j] = "left"
			}

			// Update maximum score
			if score[i][j] > maxScore {
				maxScore = score[i][j]
				maxI = i
				maxJ = j
			}
		}
	}

	// Traceback to find the aligned sequences
	align1 := ""
	align2 := ""
	i, j := maxI, maxJ
	for traceback[i][j] != "" {
		switch traceback[i][j] {
		case "diagonal":
			align1 = string(seq1[i-1]) + align1
			align2 = string(seq2[j-1]) + align2
			i--
			j--
		case "up":
			align1 = string(seq1[i-1]) + align1
			align2 = "-" + align2
			i--
		case "left":
			align1 = "-" + align1
			align2 = string(seq2[j-1]) + align2
			j--
		}
	}

	return maxScore, align1, align2
}

func main() {
	seq1 := "FKHMEDPLE"
	seq2 := "FMDTPLNE"
	match := 10
	mismatch := -15
	gapOpen := -5
	gapExtend := -1

	score, align1, align2 := SmithWaterman(seq1, seq2, int(match), int(mismatch), gapOpen, gapExtend)

	fmt.Println("Alignment Score:", score)
	fmt.Println("Aligned Sequence 1:", align1)
	fmt.Println("Aligned Sequence 2:", align2)
}
