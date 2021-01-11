package main

import (
	"fmt"
	"strings"
)

func upperName(names ...string) []string {
	var nameArray []string
	for _, name := range names {
		newName := strings.ToUpper(name)
		nameArray = append(nameArray, newName)
	}

	return nameArray
}

func main() {
	newArray := upperName("jack", "jasper", "kenny", "elly", "joon", "raphael", "awesome")
	fmt.Println(newArray)

}
