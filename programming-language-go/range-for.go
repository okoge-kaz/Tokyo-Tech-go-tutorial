package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, arg := range os.Args[1:] {// index, element value
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
// unix の echo コマンドの実装
