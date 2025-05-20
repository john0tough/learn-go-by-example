package mocking

import (
	"fmt"
	"io"
)

func Counter(w io.Writer, sleeper Sleeper) {
	countDownStart := 3
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		sleeper.Sleep()
	}
	fmt.Fprintf(w, "GO!")
}
