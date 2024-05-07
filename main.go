package main


import (
	"fmt"
	"os"
	"log"
	"encoding/hex"
	"strings"
)

const CharactersPerLine int = 13
var Output strings.Builder

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Printf("usage: %v <file-to-dump>\n", args[0])
		os.Exit(1)
	}

	filename := args[1]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	hexSlice := []string{}
	hexString := hex.EncodeToString(fileContents)
	
	count := len(hexString) / 2
	for i := range count {
		bottom := i * 2
		top := bottom + 2
		hex := hexString[bottom:top]

		hexSlice = append(hexSlice, hex)
	}
	
	top := 0
	bottom := 0
	lineLength := 0
	
	replaceBadCharacters(&fileContents)
	count = len(hexSlice) / CharactersPerLine
	for i := range count {
		top = i * CharactersPerLine
		bottom = top + CharactersPerLine
		
		region := hexSlice[top:bottom]
		strToPrint := string(fileContents[top:bottom])
		
		lineLength = printLine(&region, strToPrint, i * CharactersPerLine, 0)
	}

	region := hexSlice[bottom:]
	strToPrint := string(fileContents[bottom:])
	
	printLine(&region, strToPrint, count * CharactersPerLine, lineLength)
}

func printLine(region *[]string, str string, address int, lineLength int) int {
	Output.Reset()
	
	printAddress(address)
	for i, v := range *region {
		if i % 2 == 0 && i != 0 {
			fmt.Fprint(&Output, " ")
		}
		fmt.Fprint(&Output, v)
	}

	charCount := len(*region)
	if charCount < CharactersPerLine {
		spaceNeeded := lineLength - Output.Len()
		for _ = range spaceNeeded {
			fmt.Fprint(&Output, " ")
		}
	}

	lineLength = Output.Len()
	
	fmt.Fprint(&Output, " ", str)
	fmt.Println(Output.String())
	
	return lineLength
}

func printAddress(add int) {
	hexString := hex.EncodeToString([]byte{byte(add)})
	zeroCount := 8 - len(hexString)
	for range zeroCount {
		fmt.Fprint(&Output, "0")
	}
	
	fmt.Fprint(&Output, hexString)
	fmt.Fprint(&Output, ": ")
}

func replaceBadCharacters(arr *[]byte) {
	for i := 0; i < len(*arr); i++ {
		curr := (*arr)[i]
		if curr != 32 && (curr < 33 || curr > 126) {
			(*arr)[i] = 46
		}
	}
}
