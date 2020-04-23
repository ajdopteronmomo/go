package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	startDate := time.Date(2004, 1, 1, 0, 0, 0, 0, time.Local)
	endDate := time.Date(2020, 4, 25, 0, 0, 0, 0, time.Local)
	genTradeDays(startDate, endDate, 2)
}

//生成开始到结束日期间 增加days天数后的交易日
func genTradeDays(start time.Time, end time.Time, days int) {
	f, err := os.OpenFile("result.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()
	for queryDay := start; queryDay.Before(end); queryDay = queryDay.AddDate(0, 0, 1) {
		if queryDay.Weekday() == 0 || queryDay.Weekday() == 6 {
			continue
		}
		queryDayStr := queryDay.Format("20060102")

		resultDay := addTradeDays(queryDayStr, days)

		//记录文件中
		f.Write([]byte(queryDayStr + "|" + resultDay + "\n"))
		fmt.Println(queryDayStr + "|" + resultDay)
		time.Sleep(time.Second * 10)
	}
}

//queryDate增加交易天数后的日期
func addTradeDays(queryDay string, dayCount int) (resultDay string) {
	if dayCount == 0 {
		return queryDay
	}

	tm, _ := time.Parse("20060102", queryDay)

	tm = tm.AddDate(0, 0, 1)

	strDate := tm.Format("20060102")

	if isTradeDay(strDate) == 1 {
		dayCount = dayCount - 1
	}
	return addTradeDays(strDate, dayCount)
}

//是否交易日  1:yes 0:no
func isTradeDay(queryDay string) (result int) {
	tm, _ := time.Parse("20060102", queryDay)

	if getDayType(queryDay) == "0" && tm.Weekday() != 0 && tm.Weekday() != 6 {
		return 1
	}
	return 0
}

//返回日期类型 0 workday 1 weekend 2 holiday -1 err
func getDayType(queryDay string) (result string) {

	url := "http://tool.bitefu.net/jiari/?d=" + queryDay

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	result = string(body[0])
	return result
}
