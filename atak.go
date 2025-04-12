
package atak

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// )

// func init() {
// 	data, err := ioutil.ReadFile("/root/flag.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Fprintln(os.Stdout, string(data))
// }






import (
	"fmt"
	"os/exec"
	"strings"
)

func atak() {
	// تنفيذ أمر cat لقراءة الفلاج
	cmd := exec.Command("cat", "/root/flag.txt")
	output, err := cmd.CombinedOutput()

	// التعامل مع الخطأ
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// طباعة الفلاج
	fmt.Println("FLAG:", strings.TrimSpace(string(output)))
}













// package atak

// import (
// 	"fmt"
// 	"io/ioutil"
// )

// func init() string {
// 	data, err := ioutil.ReadFile("/root/flag.txt")
// 	if err != nil {
// 		return "Error: " + err.Error()
// 	}
// 	return "FLAG: " + string(data)
// }




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
