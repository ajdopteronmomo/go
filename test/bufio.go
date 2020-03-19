package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// name := "demo.txt"
	// content := "http://sdfdsfdsf"

	// fileObj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	// if err != nil {
	// 	fmt.Println("文件打开失败", err)
	// }

	// defer fileObj.Close()
	// writeObj := bufio.NewWriterSize(fileObj, 4096)

	// buf := []byte(content)
	// if _, err := writeObj.Write(buf); err == nil {
	// 	if err := writeObj.Flush(); err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("数据写入成功")
	// }

	fileObj, err := os.Open("demo.txt")
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}

	defer fileObj.Close()
	reader := bufio.NewReader(fileObj)
	buf := make([]byte, 1024)
	info, err := reader.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("读取的字节数:" + strconv.Itoa(info))
	fmt.Println("读取的文件内容:", string(buf))
}
