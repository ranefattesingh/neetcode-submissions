func carFleet(target int, position []int, speed []int) int {
	cars := make([][2]int, len(position))
	for i := 0; i < len(position); i++ {
		cars[i] = [2]int{position[i], speed[i]}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] > cars[j][0]
	})

	stack := make([]float64, 0)

	for i := 0; i < len(cars); i++ {
		time := float64((target-cars[i][0]))/float64(cars[i][1])
		stack = append(stack, time)
        if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
            stack = stack[:len(stack)-1]
        }
	}


	return len(stack)
}
