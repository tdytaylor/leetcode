package _322

func coinChange(coins []int, amount int) int {
	len, val := len(coins), make([]int, amount+1)

	val[0] = 0

	for i := 1; i <= amount; i++ {
		val[i] = -1
		for j := 0; j < len; j++ {
			if i >= coins[j] && val[i-coins[j]] != -1 {
				if val[i] == -1 || val[i-coins[j]]+1 < val[i] {
					val[i] = val[i-coins[j]] + 1
				}
			}
		}
	}
	return val[amount]
}
