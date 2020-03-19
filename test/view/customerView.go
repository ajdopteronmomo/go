package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"../model"
	"../service"
)

type customerView struct {

	//定义必要字段
	key  string //接收用户输入...
	loop bool   //表示是否循环的显示主菜单
	//增加一个字段customerService
	customerService *service.CustomerService
}

//显示所有的客户信息
func (this *customerView) list() {

	//首先，获取到当前所有的客户信息(在切片中)
	customers := this.customerService.List()
	//显示
	fmt.Println("---------------------------客户列表---------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		//fmt.Println(customers[i].Id,"\t", customers[i].Name...)
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Printf("\n-------------------------客户列表完成-------------------------\n\n")
}

//得到用户的输入，信息构建新的客户，并完成添加
func (this *customerView) add() {
	fmt.Println("---------------------添加客户---------------------")
	fmt.Print("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("邮箱:")
	email := ""
	fmt.Scanln(&email)
	//构建一个新的Customer实例
	//注意: id号，没有让用户输入，id是唯一的，需要系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	//调用
	if this.customerService.Add(customer) {
		fmt.Println("---------------------添加完成---------------------")
	} else {
		fmt.Println("---------------------添加失败---------------------")
	}
}

//得到用户的输入id，删除该id对应的客户
func (this *customerView) delete() {
	fmt.Println("---------------------删除客户---------------------")
	fmt.Print("请选择待删除客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}
	fmt.Println("确认是否删除(Y/N)：")
	//这里同学们可以加入一个循环判断，直到用户输入 y 或者 n,才退出..
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		//调用customerService 的 Delete方法
		if this.customerService.Delete(id) {
			fmt.Println("---------------------删除完成---------------------")
		} else {
			fmt.Println("---------------------删除失败，输入的id号不存在----")
		}
	}
}

//退出软件
func (this *customerView) exit() {

	fmt.Print("确认是否退出(Y/N)：")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}

		fmt.Print("你的输入有误，确认是否退出(Y/N)：")
	}

	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}

}

//显示主菜单
func (this *customerView) mainMenu() {

	for {
		fmt.Println("-----------------客户信息管理软件-----------------")
		fmt.Println("                 1 添 加 客 户")
		fmt.Println("                 2 修 改 客 户")
		fmt.Println("                 3 删 除 客 户")
		fmt.Println("                 4 客 户 列 表")
		fmt.Println("                 5 退       出")
		fmt.Print("请选择(1-5)：")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			fmt.Println("修 改 客 户")
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}

		if !this.loop {
			break
		}

	}
	fmt.Println("已退出了客户关系管理系统...")
}

func running(name string) {
	var times int
	for {
		times++
		fmt.Println(name, "tick", times)
		time.Sleep(time.Second)
	}
}

func asyncFunc(index int) {
	sum := 0
	for i := 0; i < 10000; i++ {
		sum += 1
	}
	fmt.Println(index, sum)
}

var wg sync.WaitGroup

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
	// wg.Add(5)
	// fmt.Println("xxx")
	// for i := 0; i < 5; i++ {
	// 	go asyncFunc(i)
	// }
	// fmt.Println("xxx1")
	// wg.Wait()
	// //在main函数中，创建一个customerView,并运行显示主菜单..
	// customerView := customerView{
	// 	key:  "",
	// 	loop: true,
	// }
	// //这里完成对customerView结构体的customerService字段的初始化
	// customerView.customerService = service.NewCustomerService()
	// //显示主菜单..
	// customerView.mainMenu()

	// s := "a"
	// go running(s)
	// s = "b"
	// go running(s)

	// var input string
	// fmt.Scanln(&input)
	// fmt.Println(input)
	// ch:=make(chan int)
	// var chSendOnly chan<-int =ch
	// var chReciveOnly chan<-int ch

	// ch := make(chan int, 2)
	// ch <- 1
	// ch <- 2
	// close(ch)
	// for i := 0; i < cap(ch)+1; i++ {
	// 	result, ok := <-ch
	// 	fmt.Println(result, ok)
	// }

}

type client chan<- string //对外发送消息的通道

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) //所有连接的客户端
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			//把所有接收到的消息广播给所有客户端
			//发送消息通道
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) //对外发送客户消息的通道
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "欢迎" + who
	messages <- who + "上线"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ":" + input.Text()
	}
	leaving <- ch
	messages <- who + "下线"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // 注意：忽略网络层面的错误
	}
}
