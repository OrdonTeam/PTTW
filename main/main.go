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
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err.Error())
	}
	return file
}

func outputFile() *os.File {
	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal(err.Error())
	}
	return file
}

func snr() float64 {
	snr, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		log.Fatal(err.Error())
	}
	return snr
}
