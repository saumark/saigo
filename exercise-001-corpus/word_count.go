package main

import (
	"fmt"
	"github.com/saumark/saigo/exercise-001-corpus/corpus"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Requires an argument")
		os.Exit(1)
	}

	dat, err := ioutil.ReadFile(os.Args[1])
	check(err)

}
