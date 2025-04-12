
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("/root/flag.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("FLAG:", string(data))
}
