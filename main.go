package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/average/utils"
)

func main() {
	numbers, err := utils.GetFloat(`average.txt`)
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println(sum / float64(len(numbers)))
}

/* go install can be executed without the path if you are already in the directory, with the .go file to be installed.
go install
instead of
go install github.com/headfirstgo/average
same with go init

for go init to work :
export PATH=$PATH:/usr/local/go/bin

execute in folder with main.go
go init
name packages the same as their folders. go files can be named differently.

*/
