package main

import (
	"fmt"
	"math/rand"
	"time"
	"runtime"
	"sync"
)

const MaxLen = 3
const MaxLen9 = 9

//班级数据
var cls9 = [][]int{
	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	{1, 2, 3, 4, 5, 6, 7, 8, 9},
}
var cls = [][]int{
	{1,2,3},
	{1,2,3},
	{1,2,3},
}
//冲突数据
var chongtu9 = [MaxLen9][MaxLen9]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0},
}
var chongtu = [3][3]int{
	{0,0,0},
	{0,0,0},
	{0,0,0},
}

func main() {

	fmt.Println("开始交换=======")

	runtime.GOMAXPROCS(runtime.NumCPU())
	proess1()
	fmt.Println(cls)
}

func proess1(){
	wg := sync.WaitGroup{}
	wg.Add(MaxLen)

	//cls1 := cls[0:1]
	for i:=0 ;i<MaxLen;i++{
		go autoWork(cls[i],i,&wg)
	}

	wg.Wait()
}

func process2(){
	c := make(chan bool)
	for i:=0 ;i<MaxLen;i++{
		go autoWork2(cls[i],i,c)
	}
	<-c
}

func autoWork(arr []int,idx int,wg *sync.WaitGroup)  {

	for{
		nums := check()
		fmt.Println("冲突数量:", nums)
		if nums == 0 {
			fmt.Println("完成排课=====")
			fmt.Println(cls)
			wg.Done()
			return
		}
		startIdx := -1
		endIdx := -1
		rand.Seed(time.Now().UnixNano())
		for{

			startIdx = rand.Intn(MaxLen)
			endIdx = rand.Intn(MaxLen)
			if startIdx != endIdx{
				break
			}
		}
		arr[startIdx], arr[endIdx] = arr[endIdx], arr[startIdx]
		fmt.Println("交换",idx,startIdx,endIdx)
	}
}
func autoWork2(arr []int,idx int,ch chan bool)  {

	for{
		nums := check()
		fmt.Println("冲突数量:", nums)
		if nums == 0 {
			fmt.Println("完成排课=====")
			fmt.Println(cls)
			ch <- true

			return
		}
		startIdx := -1
		endIdx := -1
		rand.Seed(time.Now().UnixNano())
		for{

			startIdx = rand.Intn(MaxLen)
			endIdx = rand.Intn(MaxLen)
			if startIdx != endIdx{
				break
			}
		}
		arr[startIdx], arr[endIdx] = arr[endIdx], arr[startIdx]
		fmt.Println("交换",idx,startIdx,endIdx)
	}


}

//冲突检查
func check() int {
	//冲突数
	for i := 0; i < MaxLen; i++ {
		for j :=0; j < MaxLen; j++ {
			for k:=0;k<MaxLen;k++{
				if k != i{
					if cls[i][j] == cls[k][j]{
						chongtu[i][j] = 1
						break
					}
				}
			}
		}

	}
	return total()
}

func total() int {
	//冲突数
	nums := 0
	for i := 0; i < MaxLen; i++ {
		for j :=0; j < MaxLen; j++ {
			if chongtu[i][j] > 0 {
				nums += 1
			}
		}

	}
	return nums
}

