package utils

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strconv"
)

var (
	inFile  = "./student1.xlsx"
	outFile = "./out_student.xlsx"
)

type Student struct {
	Name   string
	age    int
	Phone  string
	Gender string
	Mail   string
}

func Export() {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("student_list")
	if err != nil {
		fmt.Printf(err.Error())
	}
	stus := getStudents()
	//add data
	for _, stu := range stus {
		row := sheet.AddRow()
		nameCell := row.AddCell()
		nameCell.Value = stu.Name

		ageCell := row.AddCell()
		ageCell.Value = strconv.Itoa(stu.age)

		phoneCell := row.AddCell()
		phoneCell.Value = stu.Phone

		genderCell := row.AddCell()
		genderCell.Value = stu.Gender

		mailCell := row.AddCell()
		mailCell.Value = stu.Mail
	}
	err = file.Save(outFile)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println("\n\nexport success")
}

func getStudents() []Student {
	students := make([]Student, 0)
	for i := 0; i < 10; i++ {
		stu := Student{}
		stu.Name = "name" + strconv.Itoa(i+1)
		stu.Mail = stu.Name + "@chairis.cn"
		stu.Phone = "1380013800" + strconv.Itoa(i)
		stu.age = 20
		stu.Gender = "男"
		students = append(students, stu)
	}
	return students
}
