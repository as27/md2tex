package main

import (
	"fmt"
	"os"

	"github.com/as27/md2tex"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("second parameter needs to be the md file")
		os.Exit(1)
	}
	inFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("error opening md file:", err)
		os.Exit(1)
	}
	defer inFile.Close()
	outFile, err := os.OpenFile(os.Args[1]+".tex", os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("error creating output file:", err)
	}
	defer outFile.Close()

	err = md2tex.Convert(inFile, outFile, md2tex.SimpleConf)
	if err != nil {
		fmt.Println("error converting file:", err)
	}
}
