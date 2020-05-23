package apple_orange

func countApplesAndOranges(s int, t int, a int, b int, apples []int, oranges []int) (int, int) {
	appleCount := 0
	for _, v := range apples {
		d := a + v
		if d <= t && d >= s {
			appleCount++
		}
	}

	orgCount := 0
	for _, v := range oranges {
		d := b + v
		if d <= t && d >= s {
			orgCount++
		}
	}

	return appleCount, orgCount
}
