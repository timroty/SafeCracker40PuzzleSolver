package safe

type Level struct {
	Values [][]int
}

var (
	// Bottom Level
	level0 = Level{
		Values: [][]int{
			{2, 15, 23, 19, 3, 2, 3, 27, 20, 11, 27, 10, 19, 10, 13, 10},
			{22, 9, 5, 10, 5, 1, 24, 2, 10, 9, 7, 3, 12, 24, 10, 9},
		},
	}

	level1 = Level{
		Values: [][]int{
			{10, 0, 15, 0, 6, 0, 9, 0, 16, 0, 17, 0, 2, 0, 2, 0},
			{8, 3, 6, 15, 22, 6, 1, 1, 11, 27, 14, 5, 5, 7, 8, 24},
		},
	}

	level2 = Level{
		Values: [][]int{
			{2, 0, 22, 0, 2, 0, 17, 0, 15, 0, 14, 0, 5, 0, 10, 0},
			{10, 6, 10, 2, 6, 10, 4, 1, 5, 5, 4, 8, 6, 3, 1, 6},
		},
	}

	// Top Level
	level3 = Level{
		Values: [][]int{
			{3, 0, 6, 0, 10, 0, 10, 0, 10, 0, 6, 0, 13, 0, 3, 0},
		},
	}

	// Define safe with ordered levels
	Levels = []Level{
		level0,
		level1,
		level2,
		level3,
	}

	LevelValues = 16
	TotalAmount = 40
)
