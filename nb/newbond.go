package main

import (
	"fmt"
	"log"

	"github.com/jasonlvhit/gocron"
)

func task() {
	//定义收件人
	mailTo := []string{
		"702282623@qq.com",
		"380266713@qq.com",
		"zhouyang@gzrobot.com",
	}
	//邮件主题为"Hello"
	subject := "可转债申购通知"
	mailBody, needSend := reptile()
	if !needSend {
		return
	}
	err := SendMail(mailTo, subject, mailBody)
	if err != nil {
		log.Println("发送失败", err)
		return
	}
	fmt.Println("send over")
}

func test() {
	fmt.Println("xx")
}

func main() {
	// task()
	fmt.Println("start")
	gocron.Every(1).Day().At("09:30").Do(task)
	// gocron.Every(1).Second().Do(test)
	<-gocron.Start()
}
