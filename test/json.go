package main

import (
	// "encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

func main() {

	//-------编码-------
	// info := []Website{{"Golang", "http://c.biancheng.net/golang/", []string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutorial/"}}, {"Java", "http://c.biancheng.net/java/", []string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/python/"}}}

	// //创建文件
	// filePtr, err := os.Create("info.json")
	// if err != nil {
	// 	fmt.Println("文件创建失败", err.Error())
	// }
	// defer filePtr.Close()

	// //创建json编码器
	// encoder := json.NewEncoder(filePtr)
	// err = encoder.Encode(info)
	// if err != nil {
	// 	fmt.Println("编码错误", err.Error())
	// } else {
	// 	fmt.Println("编码成功")
	// }

	// //-------解码-------
	// filePtr, err := os.Open("./info.json")
	// if err != nil {
	// 	fmt.Println("打开文件失败", err.Error())
	// 	return
	// }
	// defer filePtr.Close()
	// var info []Website
	// //创建解码器
	// decoder := json.NewDecoder(filePtr)
	// err = decoder.Decode(&info)
	// if err != nil {
	// 	fmt.Println("解码失败", err.Error())
	// } else {
	// 	fmt.Println("解码成功")
	// 	fmt.Println(info)
	// }

	//写xml文件
	//info := []Website{{"Golang", "http://c.biancheng.net/golang/", []string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutorial/"}}, {"Java", "http://c.biancheng.net/java/", []string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/python/"}}}
	// f, err := os.Create("./info.xml")
	// if err != nil {
	// 	fmt.Println("文件创建失败", err.Error())
	// 	return
	// }
	// defer f.Close()
	// //序列化到文件中
	// encoder := xml.NewEncoder(f)
	// err = encoder.Encode(info)
	// if err != nil {
	// 	fmt.Println("编码错误", err.Error())
	// 	return
	// } else {
	// 	fmt.Println("编码成功")
	// }

	//读xml文件
	f, err := os.Open("./info.xml")
	if err != nil {
		fmt.Println("打开文件失败", err.Error())
		return
	}
	defer f.Close()
	info := Website{}
	//创建xml解码器
	decoder := xml.NewDecoder(f)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Println("解码失败", err.Error())
		return
	} else {
		fmt.Println("解码成功")
		fmt.Println(info)
	}
}
