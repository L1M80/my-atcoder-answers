// https://atcoder.jp/contests/abc132/tasks/abc132_a
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func readLine() string {
	scanner.Scan()
	return scanner.Text()
}

func main() {
	s := strings.Split(readLine(), "")
	m := make(map[string]bool)
	uniq := []string{}

	for _, elem := range s {
		if !m[elem] {
			m[elem] = true
			uniq = append(uniq, elem)
		}
	}

	if len(uniq) == 2 {
		fmt.Fprintf(writer, "Yes")
	} else {
		fmt.Fprintf(writer, "No")
	}
	writer.Flush()
}
