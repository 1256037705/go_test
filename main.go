package main

import "fmt"

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i:%v \n", i)
	}

	var (
		name  string
		age   int
		email string
	)

	fmt.Println("请输入姓名：")
	fmt.Scan(&name, &age, &email)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(email)
}

func test2() {
	var hello string
	fmt.Printf("请输入值:")
	fmt.Scan(&hello)

	switch hello {
	case "a":
		fmt.Println("123")
	case "c":
		fmt.Println("789")
	case "b":
		fmt.Println("456")
	default:
		fmt.Println("非法输入")
	}
}
func test3() {
	var a = [...]int{1, 2, 3}
	for _, i2 := range a {
		fmt.Printf(":%v   ", i2)
	}
}

func test4() {
	i := 2
	if i > 2 {
		fmt.Printf("2")
	} else {
		goto end
	}
end:
	fmt.Println("标签推出")
}

func test5() {
	va := []int{1, 3, 23}
	fmt.Printf("va:%T \n", va)

	vas := [...]string{"1", "2"}
	fmt.Printf("va:%T", vas)
}

func test6() {
	va := [...]int{1, 2, 3, 4, 5, 6, 7}
	vas := len(va)
	fmt.Printf("长度为：%v \n", vas)
	for i := 0; i < vas; i++ {

		fmt.Printf("数组：%v \n", va[i])
	}
}

func test7() {
	va := [...]int{1, 2, 3, 4, 5, 6, 7}
	for i, _ := range va {
		fmt.Printf("数  组：%v \n", va[i])
	}
}

func test8() {
	va := []int{1, 2, 3, 4, 5, 6, 7}
	va = append(va, 300, 200)
	for i, _ := range va {
		fmt.Printf("数组：%v \n", va[i])
	}
}

func test9() {
	var va = map[string]string{"a": "1", "b": "2", "c": "3"}
	vars := make(map[string]string)
	vars["name"] = "1111"
	fmt.Println(va, vars)
}

func test10(a int, b int) (c int) {
	c = a + b
	return c
}

func main() {
	//hello := user.Hello()
	//fmt.Println(hello)
	abs := test10(2, 3)
	fmt.Printf("rount:%v", abs)
}
