package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hyenaspots/yampatch"
)

func main() {

	var (
		inputYAMLBytes  []byte
		inputYAMLString string
		opsBytes        []byte
		err             error
	)

	switch {
	case len(os.Args) < 2:
		fmt.Println(`YAML`)
		os.Exit(1)

	case len(os.Args) == 2:
		inputYAMLBytes, err = ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}

		fmt.Println(string(inputYAMLBytes))

	case len(os.Args) > 2:
		inputYAMLBytes, err = ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}

		inputYAMLString = string(inputYAMLBytes)

		for i := 2; i < len(os.Args); i++ {
			opsBytes, err = ioutil.ReadFile(os.Args[i])
			if err != nil {
				panic(err)
			}

			inputYAMLString, err = yampatch.ApplyOps(inputYAMLString, string(opsBytes))
			if err != nil {
				fmt.Println(err)
				panic(err)
			}

		}

		fmt.Println(`---`)
		fmt.Println(inputYAMLString)
	}

}
