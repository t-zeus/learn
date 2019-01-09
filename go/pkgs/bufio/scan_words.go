package main

import(
    "bufio"
    "strings"
    "fmt"
)

func main() {
    input := "foo bar   baz"
    // strings.NewReader 创建 io.Reader 的实现
    scanner := bufio.NewScanner(strings.NewReader(input))
    // 设定分割函数，默认 bufio.ScanLine
    // bufio.ScanWords 会把空格都去掉
    scanner.Split(bufio.ScanWords)
    // Scan 扫描是否终止（到达文本末端、遇到错误）
    // 非终止，Scanner.Text() 有效
    // 终止，Scanner.Err() 错误原因（非 EOF）
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}
