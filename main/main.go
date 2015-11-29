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
	inputFile := inputFile()
	outputFile := outputFile()
	snr := snr()
	rand.Seed(time.Now().UTC().UnixNano())
	pttw.Algorithm(inputFile, outputFile, snr, rand.Float64)
}

func inputFile() *os.File {
	name := findParam("-in", os.Args)
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
	name := findParam("-out", os.Args)
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
	value := findParam("-snr", os.Args)
	if (value == "") {
		return 20
	}
	snr, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Fatal(err.Error())
	}
	return snr
}

func findParam(param string, args []string) string {
	for i := 0; i < len(args) - 1; i++ {
		if (args[i] == param) {
			return args[i + 1]
		}
	}
	return ""
}
