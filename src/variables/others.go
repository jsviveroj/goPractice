package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var State bool
var Salary float32
var Date time.Time

func OtherVariables() {
	Name = "Pedro"
	State = true
	Salary = 3000000
	Date = time.Now()
	fmt.Println(Name)
	fmt.Println(State)
	fmt.Println(Salary)
	fmt.Println(Date)
}

func ConvertToInt(number int) (bool, string) {
	texto := strconv.Itoa(number)
	return true, texto
}
