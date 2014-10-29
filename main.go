package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {

	// http://golangcookbook.blogspot.ca/2012/11/generate-random-number-in-given-range.html
	rand.Seed(time.Now().Unix())

	byteArray, err := ioutil.ReadFile("./tests/pdfs/pdf-sample.pdf")

	if err != nil {
		fmt.Println("error reading file")
		panic(err)
	}

	fuzzFactor := 250
	numWrites := len(byteArray)/fuzzFactor + 1
	fmt.Printf("Writing %d random bytes...\n", numWrites)

	for i := 0; i < numWrites; i++ {
		rn := random(0, len(byteArray))
		rbyte := byte(random(0, 255))
		byteArray[rn] = rbyte
	}

	err = ioutil.WriteFile("./changed.pdf", byteArray, 0644)

	if err != nil {
		fmt.Println("error writing file")
		panic(err)
	} else {
		fmt.Println("Done!")
	}
}
