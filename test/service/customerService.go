package service

import (
	"../model"
)

//CustomerService  完成对Customer的操作,包括
//增删改查
type CustomerService struct {
	customers []model.Customer
	//声明一个字段，表示当前切片含有多少个客户
	//该字段后面，还可以作为新客户的id+1
	customerNum int
}

//NewCustomerService 编写一个方法，可以返回*CustomerService
func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "0101-123213", "sd@suho.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//List 返回客户切片
func (t *CustomerService) List() []model.Customer {
	return t.customers
}

//Add 添加客户到customers
func (t *CustomerService) Add(customer model.Customer) bool {
	t.customerNum++
	customer.ID = t.customerNum
	t.customers = append(t.customers, customer)
	return true
}

//Delete 根据id删除客户(从切片中删除)
func (t *CustomerService) Delete(id int) bool {
	index := t.FindByID(id)
	if index == -1 {
		return false
	}
	t.customers = append(t.customers[:index], t.customers[index+1:]...)
	return true
}

//FindByID 根据id查找客户在切片中对应下标，如果没有该客户，返回-1
func (t *CustomerService) FindByID(id int) int {
	index := 1
	for i := 0; i < len(t.customers); i++ {
		if t.customers[i].ID == id {
			index = i
		}
	}
	return index
}
