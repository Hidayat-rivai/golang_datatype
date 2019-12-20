package main

import "fmt"

func main(){
	var uns8 uint8 = 255
	var uns16 uint16 = 65535
	var uns32 uint32 = 4294967295
	var uns64 uint64 = 18446744073709551615
	var uint_def uint = 18446744073709551615
	var byte_def byte = 255

	fmt.Println(uns8)
	fmt.Println(uns16)
	fmt.Println(uns32)
	fmt.Println(uns64)
	fmt.Println(uint_def)
	fmt.Println(byte_def)
}