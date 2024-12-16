package action

import "fmt"

const PI float64 = 3.14159264

var (
	BMI int = 1
)

// 定义变量
func DefineVariableStyle1() {
	var name string = "wt"
	age := 23
	height, weight := 178, 75

	fmt.Println(name, age, height, weight)
	fmt.Printf("BMI:%d\n", BMI)
	fmt.Printf("PI:%f\n", PI)
	fmt.Printf("PI:%v PI:%T\n", PI, PI)
}

func InputVariable() {
	var name string
	fmt.Println("请输入您的姓名：")
	fmt.Scan(&name)
	fmt.Printf("您的姓名是：%s\n", name)
}
