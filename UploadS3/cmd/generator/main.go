package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < 20; i++ {
		f, err := os.Create(fmt.Sprintf("./tmp/file%d.txt", i))
		if err != nil {
			panic(err)
		}

		if _, werr := f.WriteString("Hello, World!"); werr != nil {
			f.Close()
			panic(werr)
		}
		if cerr := f.Close(); cerr != nil {
			panic(cerr)
		}
	}
}
