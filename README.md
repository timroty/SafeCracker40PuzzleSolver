# SafeCracker40PuzzleSolver

Solver for Creative Crafthouse's SafeCracker Puzzle.

## About the Puzzle

<p align="center">
  <img src="./assets/SafeCrackerPuzzle.png" width="200" height="200" />
</p>

The SafeCracker Puzzle is a puzzle created by Creative Crafthouse, based on a 1911 English design, with the objective of getting each of the 16 columns of numbers to add to 40 at the same time. The top ring layers are notched and will expose and hide numbers as the layers are rotated. There is only 1 possible solution.

Find out more about the [SafeCracker Puzzle](https://www.creativecrafthouse.com/safecracker-40-puzzle-can-you-find-the-combination.html)!

## Solving the Puzzle

This program solves the puzzle by systematically iterating through combinations of layer positions until the solution is found.

### How It Works

The solver uses an exhaustive search algorithm:
- It tries all possible rotations of the puzzle layers
- Checks each configuration to see if it satisfies the 40-point condition
- Stops when the unique solution is found

### Installation

1. Clone the repository: `git clone https://github.com/yourusername/SafeCracker40PuzzleSolver.git`

2. Navigate to the project directory: `cd SafeCracker40PuzzleSolver`

3. Run the program: `go run main.go`

### Requirements

- Go version 1.23.4 or higher

### Results

The program will print the solution and additional information to the console. The first block represents the number of rotations needed to solve the puzzle for each level. The levels are ordered from bottom to top, with the leftmost value representing the bottom layer and the rightmost value representing the top layer.

To solve the puzzle, rotate each layer *counter-clockwise* by the number of spaces corresponding to its position in the solution array. The second block prints the values at one specific position to help you align the puzzle correctly during rotation.
