package main

import "fmt"

// 自定义错误结构体
type myErr struct {
	s string
}

// 实现 error 接口，指针实现的
func (e *myErr) Error() string {
	return e.s
}

// 产生错误
func NewErr(msg string) error {
	return &myErr{msg}
}

func openFile() ([]byte, error) {
	/**
	这种方式是直接调用包内的错误结构体，但是如果是别的包想要调用呢？
	定义个对外的函数，来包裹错误结构体引用。
	*/
	return nil, &myErr{"打开错误"}
}

func openDir() ([]string, error) {
	return nil, NewErr("打开目录失败")
}

func main() {
	var err error

	_, err = openFile()
	if err != nil {
		fmt.Println(err)
	}

	_, err = openDir()
	if err != nil {
		fmt.Println(err)
	}
}
