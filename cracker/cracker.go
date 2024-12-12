package cracker

import (
	"SafeCracker40PuzzleSolver/safe"
	"fmt"
)

func Crack(safe []safe.Level) {
	levelIndices := make([]int, len(safe))

	for {
		if checkSafe(safe, levelIndices) {
			fmt.Println("Solved!")
			fmt.Println("--------")
			fmt.Println(levelIndices)
			fmt.Println("--------")
			printSafeSolution(safe, levelIndices)

			break
		}

		// Iterate through each level
		if !iterateIndices(levelIndices, len(levelIndices)-1) {
			fmt.Println("Failed :(")
			break
		}
	}
}

func iterateIndices(slice []int, i int) bool {
	slice[i] = slice[i] + 1

	if slice[i] >= safe.LevelValues {
		slice[i] = 0

		// If we have wrapped around on the bottom level
		// that means no solution was found
		if i == 0 {
			return false
		}

		return iterateIndices(slice, i-1)
	}

	return true
}

func iterateIndex(i int, amount int) int {
	i = i + amount

	// Wrap around
	if i >= safe.LevelValues {
		i = i - safe.LevelValues
	}

	return i
}

func checkSafe(safeLevels []safe.Level, indices []int) bool {
	// Assume if four in a row equal the total
	// amount, then we have found a solution
	for i := 0; i < 4; i++ {
		sum := 0
		for j := 0; j < len(safeLevels); j++ {
			indexValue := iterateIndex(indices[j], i)
			positionValue := safeLevels[j].Values[0][indexValue]

			// If empty, sum the underneath value
			if positionValue == 0 {
				positionValue = safeLevels[j-1].Values[1][iterateIndex(indices[j-1], i)]
			}

			sum += positionValue
		}

		if sum != safe.TotalAmount {
			return false
		}
	}

	// We solved the puzzle!
	return true
}

func printSafeSolution(safe []safe.Level, levelIndices []int) {
	for i, level := range safe {
		fmt.Printf("Level %d:\n", i)
		fmt.Printf(" Top: %d\n", level.Values[0][levelIndices[i]])

		if i < len(safe)-1 {
			fmt.Printf(" Bottom: %d\n", level.Values[1][levelIndices[i]])
		}
	}
}
