package day2

func EveryStrTotargetStrShortestDistance(s string, c byte) []int {
	Sarr := []byte{}
	SP := []int{}
	for _, v := range s {
		Sarr = append(Sarr, byte(v))
	}
	for i := 0; i < len(Sarr); i++ {
		step := 0

		for true {
			if i+step < len(Sarr) && Sarr[i+step] == c {
				SP = append(SP, step)
				step = 1
				break
			}
			if i-step >= 0 && Sarr[i-step] == c {
				SP = append(SP, step)
				step = 1
				break
			}
			step++
		}
	}
	return SP
}
