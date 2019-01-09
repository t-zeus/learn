package main

import (
	"bufio"
	"fmt"
	"strings"
    "os"
)

func scan(input string) {
	scanner := bufio.NewScanner(strings.NewReader(input)) // 第一步：创建

	// 定义 bufio.SplitFunc
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		// 当入参是空的时候，返回 bufio.ErrFinalToken，告诉 Scan 没有多余的 token 了，不会报错
		return 0, data, bufio.ErrFinalToken
	}

	scanner.Split(onComma) // 第二步：指定分割

	// 第三步：扫描
	for scanner.Scan() {
		fmt.Printf("%q ", scanner.Text())
	}

    fmt.Println()

	// 第四部：验证错误
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input: \n", err)
	}

}

func main() {
    scan("1,2,3,4,") // 最后是空

    println("\n-----------\n")

    scan("1, 2 ,3")
}
