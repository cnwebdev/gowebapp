// Basic string literal
package main

import "fmt"

func main() {

	var a string
	var b string
	a = "Hồ Chủ Tịch Sống Mãi Trong Quần Chúng Ta!"
	b = `Hồ Chủ Tịch Sống Mãi Trong Quần Chúng Ta!`
	fmt.Println(a)
	fmt.Println(b)

	htmlTmpl := "<html>\n\t<body>\n\t<h1>Hello World!</h1>\n\t</body>\n</html>"
	fmt.Println(htmlTmpl)

	// Raw string literal allow to print multiple lines text
	htmlTmpl = `
	<html>
		<body>
			<h1>Hello World!</h1>
		</body>
	</html>`
	fmt.Println(htmlTmpl)

	// Print files and directories example
	fmt.Println("c:\\home\\user\\files")
	fmt.Println(`c:\home\user\files`)
}