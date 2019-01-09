package main

import (
    "os"
    "fmt"
    "bufio"
)

func main() {
    // Scanner.Spit() 默认的函数就是 SplitLines
    // 如果使用默认的分割函数，那么可以省略 Split
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        fmt.Println(s.Text())
    }
    if err := s.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "reading standard input:", err)
    }
}
