package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"technocage.com/hello/pkg/messagegen"
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
	fmt.Println(messagegen.GetGreeting())

	logger.Print("Waiting for all tasks to complete.")
	wg.Wait()
	logger.Print("Finished.")
}
