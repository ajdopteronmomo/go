package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

func main() {
	//写gob文件
	// info := map[string]string{
	// 	"name":    "C语言中文网",
	// 	"website": "http://c.biancheng.net/golang/",
	// }
	// name := "demo.gob"
	// File, _ := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0777)
	// defer File.Close()
	// enc := gob.NewEncoder(File)
	// if err := enc.Encode(info); err != nil {
	// 	fmt.Println(err)
	// }

	// //读gob文件
	var M map[string]string
	File, _ := os.Open("demo.gob")
	D := gob.NewDecoder(File)
	D.Decode(&M)
	fmt.Println(M)

	// var M map[string]string
	// File, _ := os.Open("demo.gob")
	// D := gob.NewDecoder(File)
	// D.Decode(&M)
	// fmt.Println(M)
}
