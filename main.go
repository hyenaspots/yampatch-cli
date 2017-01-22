package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`YAML`)
		os.Exit(0)
	}

	inputYAMLFile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Println(string(inputYAMLFile))
}
