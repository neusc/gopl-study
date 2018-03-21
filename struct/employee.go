package main

import (
	"time"
	"fmt"
)

type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

func main() {
	var fleta Employee

	position := &fleta.Position
	*position = "Senior" + *position

	var employeeOfTheMonth *Employee = &fleta
	employeeOfTheMonth.Position += "(proactive team player)"
	// 等价于(*employeeOfTheMonth).Position += "(proactive team player)"

	fmt.Println(employeeOfTheMonth)
}
