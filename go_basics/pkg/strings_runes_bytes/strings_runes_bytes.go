package strings_runes_bytes

import "strings"

func Init() {
	var myString = "résumé"
	var index = myString[1]
	println("%v, %T\n", index, index)
	// range keyword does extra work to notice 2nd index is 2 byte char and decode it to 233 and skips the index 2
	for i, v := range myString {
		println(i, v)
	}

	// len will be +2 because it returns the bytes length
	// not the char length
	println("length of the string is ", len(myString))

	// to get the actual char length convert string into rune
	// rune is char representation of the string and are alias of int32
	// while string is a read only slice of bytes
	var myString2 = []rune("résumé")
	println("length of the string is ", len(myString2))

	// string appending
	var strSlice = []string{"s", "u", "b", "w", "a", "y"}
	var concatString = ""

	// strings are immutable
	// thus when i do ++ every time a new string will be created
	for _, value := range strSlice {
		concatString += value
	}

	// above is pretty inefficient - instead do below
	var strBuilder strings.Builder
	// array is ceated and values are appended into it
	for _, value := range strSlice {
		strBuilder.WriteString(value)
	}

	// only at the end the string is created by appending above array values
	println(strBuilder.String())
}
