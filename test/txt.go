package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// //创建一个新文件,写入内容
	// filePath := "./output.txt"
	// file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	// if err != nil {
	// 	fmt.Printf("打开文件错误=%v \n", err)
	// 	return
	// }
	// defer file.Close()
	// str := "xx11dd\n"
	// writer := bufio.NewWriter(file)
	// for i := 0; i < 3; i++ {
	// 	writer.WriteString(str)
	// }
	// writer.Flush()

	//读纯文本文件
	file, err := os.Open("./output.txt")
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}
	fmt.Println("文件读取结束")
}
