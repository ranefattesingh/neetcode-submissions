func carFleet(target int, position []int, speed []int) int {
	cars := make([][2]int, 0)

	for i := 0; i < len(position); i++ {
		car := [2]int{position[i], speed[i]}
		cars = append(cars, car)
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] > cars[j][0]
	})

	stack := []float64{}
	for _, car := range cars {
		time := float64(target - car[0])/float64(car[1])
		if len(stack) == 0 {
			stack = append(stack, time)
		} else {
			top := stack[len(stack)-1]
			if time > top {
				stack = append(stack, time)
			}
		}
	}

	return len(stack)
}