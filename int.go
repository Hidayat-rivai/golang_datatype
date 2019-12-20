package main

import "fmt"

func main(){
	var integ8 int8 = 127
	var integ16 int16 = 32767
	var integ32 int32 = 2147483647
	var integ64 int64 = 9223372036854775807
	var int_def int = 9223372036854775807
	var rune_def rune = 2147483647

	fmt.Println(integ8)
	fmt.Println(integ16)
	fmt.Println(integ32)
	fmt.Println(integ64)
	fmt.Println(int_def)
	fmt.Println(rune_def)
}