package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/goquery"
	"golang.org/x/mahonia"
)

//BondInfo 债券信息
type BondInfo struct {
	stockCode    string //债券代码
	stockName    string //债券简称
	purchaseDate string //申购日期
}

func reptile() (string, bool) {
	url := "http://bond.jrj.com.cn/data/kzz.shtml?to=pc"
	// 生成client客户端
	client := &http.Client{}
	//生成Request对象
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	//发起请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//关闭响应体
	defer resp.Body.Close()
	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//解决中文乱码
	bodystr := mahonia.NewDecoder("gbk").ConvertString(string(body))

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(bodystr))
	if err != nil {
		log.Fatal(err)
	}

	stockList := make([]BondInfo, 0)
	newStockCode := ""
	newStockName := ""
	newPurchaseDate := ""

	dom.Find(".tableSw tbody tr td").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		//代码
		if i%13 == 0 {
			newStockCode = selection.Text()
		}
		//名称
		if i%13 == 1 {
			newStockName = selection.Text()
		}
		//申购日期
		if i%13 == 3 {
			newPurchaseDate = selection.Text()
			//小于今天
			arr := strings.Split(newPurchaseDate, "-")
			month, err := strconv.Atoi(arr[0])
			if err != nil {
				log.Fatal(err)
			}
			day, err := strconv.Atoi(strings.Split(arr[1], "(")[0])
			if err != nil {
				log.Fatal(err)
			}
			if month != int(time.Now().Month()) || day < time.Now().Day() {
				return false
			}
			stockList = append(stockList, BondInfo{
				stockCode:    newStockCode,
				stockName:    newStockName,
				purchaseDate: strconv.Itoa(month) + "-" + strconv.Itoa(day),
			})
		}
		return true
	})

	// 邮件正文
	mailBody := ""
	needSend := false
	for i, value := range stockList {
		arr := strings.Split(value.purchaseDate, "-")

		month, err := strconv.Atoi(arr[0])
		if err != nil {
			log.Fatal(err)
		}
		day, err := strconv.Atoi(arr[1])
		if err != nil {
			log.Fatal(err)
		}
		if month != int(time.Now().Month()) || day != time.Now().Day() {
			continue
		}
		mailBody += "第" + strconv.Itoa(i) + "只:[证券代码]" + value.stockCode + "[证券名称]" + value.stockName + "[申购日期]" + value.purchaseDate + "<br>"
		needSend = true
	}
	return mailBody, needSend
}

// func get() {
// 	url := "http://data.eastmoney.com/kzz/default.html"
// 	// http://data.eastmoney.com/kzz/default.html
// 	// 生成client客户端
// 	client := &http.Client{}
// 	//生成Request对象
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	//发起请求
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	//关闭响应体
// 	defer resp.Body.Close()
// 	// 读取响应
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	// fmt.Println(string(body))
// 	//解决中文乱码
// 	bodystr := mahonia.NewDecoder("gbk").ConvertString(string(body))
// 	fmt.Println(bodystr)
// }
