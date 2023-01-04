package minimumroundstocompletealltasks

func minimumRounds(tasks []int) int {
	m := make(map[int]int)

	for _, task := range tasks {
		m[task] += 1
	}

	res := 0
	for _, val := range m {
		if val == 1 {
			return -1
		}
		res += (val + 2) / 3
	}

	return res
}

/**
2 2 = 1 task

2 2 2 = 1 task

2 2 2 2 = 2 tasks

2 2 2 2 2 = 2 tasks

2 2 2 2 2 2 = 2 tasks

2 2 2 2 2 2 2 = 3 tasks

2 2 2 2 2 2 2 2 = 3 tasks

2 2 2 2 2 2 2 2 2 = 3 tasks

2 2 2 2 2 2 2 2 2 2 = 4 tasks

2 2 2 2 2 2 2 2 2 2 2 = 4 tasks

2 2 2 2 2 2 2 2 2 2 2 2 = 4 tasks


2 -> 1
3 -> 1

4 -> 2
5 -> 2
6 -> 2

7 -> 3
8 -> 3
9 -> 3
**/
