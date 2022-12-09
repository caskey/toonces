package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"technocage.com/weightconverter/pkg/units"
)

func panic(e error) {
	if e != nil {
		panic(e)
	}
}

var (
	projectName string
	// Our global logger for output
	logger = log.New(os.Stderr, projectName+" ", log.Lshortfile)
)

func main() {
	flag.Parse()

	wg := &sync.WaitGroup{}

	// BODY HERE
	//	web.Start(wg)
	// var inputGrams = 1000.0
	// var inputLbs = 1.0
	// fmt.Printf("%v grams in lbs is %v\n", inputGrams, units.ConvertGtoLb(inputGrams))
	// fmt.Printf("%v lbs in gs is %v\n", inputLbs, units.ConvertLbtoG(inputLbs))

	var converter = units.GetConverter(logger)

	for i, input := range flag.Args() {
		logger.Printf("Processing arg %v: \"%v\"", i, input)
		if i != 0 {
			fmt.Print("\n")
		}
		// fmt.Printf("Cmdline: %v\n", input)
		fmt.Printf("%v", converter.Convert(input))
		logger.Printf("Done rocessing arg %v: \"%v\"", i, input)
	}

	logger.Print("Waiting for all tasks to complete.")
	wg.Wait()
	logger.Print("Finished.")
}
