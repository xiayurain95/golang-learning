package main

func main() {

}
func canCompleteCircuit(gas []int, cost []int) int {
	//m := make(map[int]int)
	tmp := make([]int, len(gas))
	for i := 0; i < len(gas); i++ {
		tmp[i] = gas[i] - cost[i]
	}
	out := 0
	for i := 0; i < len(tmp); i++ {
		out += tmp[i]
	}
	if out < 0 {
		return -1
	}
	ans := 0
	for i := 0; i < len(tmp); {
		if tmp[i] < 0 {
			i++
		} else {
			preAns := ans
			ans = i
			tmpCost := 0
			for i < len(tmp) && tmpCost >= 0 {
				tmpCost += tmp[i]
			}
			if tmpCost < 0 {
				ans = preAns
			}
		}
	}
	return ans
}
