package main

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

func numberOfGoodSubsets(nums []int) (ans int) {
	const mod int = 1e9 + 7
	freq := [31]int{}
	for _, num := range nums {
		freq[num]++
	}

	f := make([]int, 1<<len(primes))
	f[0] = 1
	for i := 0; i < freq[1]; i++ {
		f[0] = f[0] * 2 % mod
	}
next:
	for i := 2; i < 31; i++ {
		if freq[i] == 0 {
			continue
		}

		// 检查 i 的每个质因数是否均不超过 1 个
		subset := 0
		for j, prime := range primes {
			if i%(prime*prime) == 0 {
				continue next
			}
			if i%prime == 0 {
				subset |= 1 << j
			}
		}

		// 动态规划
		for mask := 1 << len(primes); mask > 0; mask-- {
			if mask&subset == subset {
				f[mask] = (f[mask] + f[mask^subset]*freq[i]) % mod
			}
		}
	}

	for _, v := range f[1:] {
		ans = (ans + v) % mod
	}
	return
}
