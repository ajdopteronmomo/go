package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type Website struct {
	Url int32
}

func main() {
	// file, err := os.Create("output.bin")
	// if err != nil {
	// 	fmt.Println("文件创建失败", err.Error())
	// }

	// defer file.Close()

	// for i := 1; i <= 10; i++ {
	// 	info := Website{
	// 		int32(i),
	// 	}

	// 	var bin_buf bytes.Buffer
	// 	binary.Write(&bin_buf, binary.LittleEndian, info)
	// 	b := bin_buf.Bytes()
	// 	_, err := file.Write(b)
	// 	if err != nil {
	// 		fmt.Println("编码失败", err.Error())
	// 		return
	// 	}
	// }
	// fmt.Println("编码成功")

	file, err := os.Open("output.bin")
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return
	}
	defer file.Close()

	m := Website{}
	for i := 1; i <= 10; i++ {
		data := readNextBytes(file, 4)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.LittleEndian, &m)
		if err != nil {
			fmt.Println("文件读取失败", err)
			return
		}
		fmt.Println("第", i, "值为：", m)
	}
}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {
		fmt.Println("解码失败", err)
	}
	return bytes
}
