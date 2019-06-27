// https://atcoder.jp/contests/abc088/tasks/abc088_b
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func readLine() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	n, _ := strconv.Atoi(readLine())
	a := strings.Split(readLine(), " ")
	cards := make([]int, n)
	alice := 0
	bob := 0

	for i := 0; i < n; i++ {
		cards[i], _ = strconv.Atoi(a[i])
	}

	for i := 0; i < n; i++ {
		sort.Sort(sort.Reverse(sort.IntSlice(cards)))
		point := cards[i]
		if i%2 == 0 {
			alice += point
		} else {
			bob += point
		}
	}

	fmt.Fprintf(writer, strconv.Itoa(alice-bob))
	_ = writer.Flush()
}
