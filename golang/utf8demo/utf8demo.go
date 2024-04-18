package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "日 本 \x80 語"
	fmt.Println(s)
	fmt.Println("Number of characters: ", utf8.RuneCountInString(s))
	for i, char := range s {
		fmt.Printf("character %#U starts at byte position %d\n", char, i)
	}
	fmt.Println()

	fmt.Println("Number of bytes: ", len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			fmt.Printf("<space> ")
			continue
		}
		fmt.Printf("%X ", s[i])
	}
	fmt.Println()
	fmt.Println()

	r, size := utf8.DecodeRuneInString(s)
	fmt.Printf("First character: %#U (%d bytes)\n", r, size)
	var u uint32
	for i := 0; i < size; i++ {
		u <<= 8
		fmt.Printf("%b ", s[i])
		u |= uint32(s[i])
	}
	fmt.Printf("-> 0x%X\n", u)

	u = 0
	for i := 0; i < size; i++ {
		mask := uint8(0b00111111)
		if i == 0 {
			mask = 0b00001111
		}
		b := s[i] & mask
		if i == 0 {
			fmt.Printf("____%04b ", b)
		} else {
			fmt.Printf("__%06b ", b)
		}
		u <<= 6
		u |= uint32(b)
	}
	fmt.Printf("-> 0x%X\n", u)
}
