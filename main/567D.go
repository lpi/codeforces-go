package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Sol567D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, k, sz, m int
	Fscan(in, &n, &k, &sz, &m)
	x := make([]int, m)
	for i := range x {
		Fscan(in, &x[i])
	}

	ans := sort.Search(m+1, func(end int) bool {
		if end == 0 {
			return false
		}
		y := make([]int, end)
		copy(y, x)
		sort.Ints(y)
		y = append(y, n)
		cnt := (y[0] - 1) / sz
		for i := 1; i <= end; i++ {
			cnt += (y[i] - y[i-1]) / sz
			if cnt >= k {
				return false
			}
		}
		return true
	})
	if ans == m+1 {
		ans = -1
	}
	Fprint(out, ans)
}

func main() {
	Sol567D(os.Stdin, os.Stdout)
}
