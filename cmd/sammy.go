package main

import (
	"fmt"

	"github.com/kayex/sammy"
)

func main() {

	cs, err := sammy.GenerateChangeSet("./samples", sammy.ExtendMajor, sammy.ExtendMinor)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", cs)
	fmt.Printf("\nFound %d samples", len(cs))
}
