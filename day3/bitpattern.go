package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var (
		inputWidth = 12
		o2         = make([]byte, inputWidth)
		co2        = make([]byte, inputWidth)
	)

	for bit := 0; bit < inputWidth; bit++ {
		var o2zeroes, o2ones, co2zeroes, co2ones int
		var o2prefix, co2prefix string
		if bit > 0 {
			o2prefix = string(o2[:bit])
			co2prefix = string(co2[:bit])
		}

		fmt.Println("---")
		fmt.Println(bit, ": ", o2prefix, co2prefix)

		file, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			input := scanner.Bytes()
			if bit == 0 || string(input[:bit]) == o2prefix {
				fmt.Println("o2: ", input[bit])
				switch input[bit] {
				case '0':
					o2zeroes++
				case '1':
					o2ones++
				}
			}
			if bit == 0 || string(input[:bit]) == co2prefix {
				fmt.Println("co2: ", input[bit])
				switch input[bit] {
				case '0':
					co2zeroes++
				case '1':
					co2ones++
				}
			}
		}

		_ = file.Close()

		fmt.Println(o2zeroes, o2ones, co2zeroes, co2ones)
		if o2zeroes > o2ones {
			o2[bit] = '0'
		} else {
			o2[bit] = '1'
		}

		if co2zeroes > 0 && (co2zeroes <= co2ones || co2ones == 0) {
			co2[bit] = '0'
		} else {
			co2[bit] = '1'
		}
	}

	bitValue := 1
	var o2val, co2val int
	for i := inputWidth - 1; i >= 0; i-- {
		if o2[i] == '1' {
			o2val += bitValue
		}
		if co2[i] == '1' {
			co2val += bitValue
		}
		bitValue *= 2
	}

	fmt.Println("Day 3, part 2: ", o2val*co2val)
}
