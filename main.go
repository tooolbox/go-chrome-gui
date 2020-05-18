package main

import (
	"log"
	"os"
	"strconv"

	"github.com/zserge/lorca"
)

func main() {
	// Create UI with basic HTML passed via data URI
	width, height := 800, 450
	if len(os.Args) >= 3 {
		if w, err := strconv.Atoi(os.Args[2]); err == nil {
			width = w
		}
	}
	if len(os.Args) >= 4 {
		if h, err := strconv.Atoi(os.Args[3]); err == nil {
			height = h
		}
	}
	ui, err := lorca.New(os.Args[1], "", width, height)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Wait until UI window is closed
	<-ui.Done()
}
