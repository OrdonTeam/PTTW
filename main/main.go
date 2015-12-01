package main

import (
	"os"
	"github.com/OrdonTeam/PTTW/pttw"
	"log"
	"time"
	"math/rand"
	"strconv"
)

func main() {
	if (pttw.ShouldDisplayHelp(os.Args)) {
		pttw.DisplayHelpScreen(os.Args)
	}else {
		inputFile := inputFile()
		outputFile := outputFile()
		snr := snr()
		rand.Seed(time.Now().UTC().UnixNano())
		pttw.Algorithm(inputFile, outputFile, snr, rand.Float64)
	}
}

func inputFile() *os.File {
	name := pttw.FindParam("-in", os.Args)
	if (name == "") {
		return os.Stdin
	}
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err.Error())
	}
	return file
}

func outputFile() *os.File {
	name := pttw.FindParam("-out", os.Args)
	if (name == "") {
		return os.Stdout
	}
	file, err := os.Create(name)
	if err != nil {
		log.Fatal(err.Error())
	}
	return file
}

func snr() float64 {
	value := pttw.FindParam("-snr", os.Args)
	if (value == "") {
		return 20
	}
	snr, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatal(err.Error())
	}
	return snr
}