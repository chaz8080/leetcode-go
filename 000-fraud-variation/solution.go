package fraudvariation

const (
	FAIL           = "Fail"
	PASS           = "Pass"
	NEED_MORE_INFO = "Need More Info"
	INCONCLUSIVE   = "Inconclusive"
)

type State struct {
	isFail          bool
	isPass          bool
	isNeedsMoreInfo bool
}

func checkData(data map[string]int, state *State) {
	_, hasFs := data["F"]
	_, hasPs := data["P"]
	_, hasUs := data["_"]

	if len(data) == 1 && hasFs {
		state.isFail = true
		return
	}

	if len(data) == 1 && hasPs {
		state.isPass = true
	}

	if len(data) == 2 && hasUs {
		state.isNeedsMoreInfo = true
	}
}

func detectFraud(grid [][]string) string {
	state := &State{}
	gridLen := len(grid)

	// check rows
	for row := 0; row < gridLen; row++ {
		chars := map[string]int{}
		for col := 0; col < gridLen; col++ {
			val := grid[row][col]
			chars[val] += 1
		}

		checkData(chars, state)
		if state.isFail {
			return FAIL
		}
	}

	// check cols
	for col := 0; col < gridLen; col++ {
		chars := map[string]int{}
		for row := 0; row < gridLen; row++ {
			val := grid[row][col]
			chars[val] += 1
		}

		checkData(chars, state)
		if state.isFail {
			return FAIL
		}
	}

	leftToRight := map[string]int{}
	for position := 0; position < gridLen; position++ {
		val := grid[position][position]
		leftToRight[val] += 1
	}

	checkData(leftToRight, state)
	if state.isFail {
		return FAIL
	}

	rightToLeft := map[string]int{}
	for position := 0; position < gridLen; position++ {
		val := grid[position][gridLen-position-1]
		rightToLeft[val] += 1
	}

	checkData(rightToLeft, state)
	if state.isFail {
		return FAIL
	}

	if state.isPass {
		return PASS
	}

	if state.isNeedsMoreInfo {
		return NEED_MORE_INFO
	}

	return INCONCLUSIVE
}
