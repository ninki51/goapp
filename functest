package main

import "fmt"

func test1(s1 [] int)  {
	fmt.Println("传递切片参数")
	fmt.Println(s1)
	fmt.Println("修改=====")
	s1[0] = 100
}

func test2(m1 map[string]int){
	fmt.Println("传递map参数")
	fmt.Println(m1)
	fmt.Println("修改=====")
	m1["age"] = 28
}

func test3(a,b int) int{
	return a+b
}

//闭包
func test4(a,b int) func() int{
	c := a + b
	return func() int {
		return c * 2
	}
}


func main()  {
	s := []int {1,2,3,4}
	test1(s)
	fmt.Println(s)

	m := make(map[string] int)
	m["age"] = 18
	test2(m)
	fmt.Println(m)

	sum := test3(1,2)
	fmt.Println(sum)

	sum2 := test4(2,3)
	fmt.Println(sum2())
}
