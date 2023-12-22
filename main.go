// main
package main

import (
	"fmt"
	"os"
)

func main() {

	data, err := os.ReadFile("testdata/server.crt")
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(string(data))

}
