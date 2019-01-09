package main

import (
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

func main() {
    input := "1234 5678 1234567901234567890"
    s :=  bufio.NewScanner(strings.NewReader(input))

    // 自定义分割函数
    s.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error){
        advance, token, err = bufio.ScanWords(data, atEOF)
        if err == nil && token != nil {
            // 如果超过 int 的返回，会报错。这个会被 Scan() 捕获到
            _, err = strconv.ParseInt(string(token), 10, 32) 
        }
        return
    })

    for s.Scan() {
        fmt.Println(s.Text())
    }

    if err := s.Err(); err != nil {
        fmt.Printf("Invalid input: %s", err)
    }
}
