package butils

import "fmt"

func PrintHex(in []byte, slice int) {
	// revise
	if slice <= 0 {
		slice = 100
	} else if slice <= 4 {
		slice = 4
	}

	// print
	var tmp string
	for i, v := range in {
		fmt.Printf("%02x ", v)
		tmp += string(v)

		if (i+1)%slice == 0 {
			fmt.Printf("%s\n", tmp)
			tmp = ""
		}
	}

	if len(in)%slice != 0 {
		for i := 0; i < (slice - (len(in) % slice)); i++ {
			fmt.Printf("%s ", "  ")
		}
		fmt.Printf("%s\n", tmp)
	}
}
