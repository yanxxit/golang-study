package main

import (
	"fmt"
)

// 薪资计算器接口
type SalaryCalculator interface {
	CalculateSalary() int
}
// 普通挖掘机员工
type Contract struct {
	empId  int
	basicpay int
}
// 有蓝翔技校证的员工
type Permanent struct {
	empId  int
	basicpay int
	jj int // 奖金
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.jj
}

func (p *Permanent)Work()  {
	fmt.Println(p.empId,"working!")
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}
// 总开支
func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("总开支 $%d", expense)
}

func main() {
	pemp1 := Permanent{1,3000,10000}
	pemp2 := Permanent{2, 3000, 20000}
	cemp1 := Contract{3, 3000}
	pemp2.Work()
	employees := []SalaryCalculator{pemp1, pemp2, cemp1}//SalaryCalculator节点下，必须要满足都有CalculateSalary方法，是否有其他方法，不限制，统一行为
	totalExpense(employees)
}