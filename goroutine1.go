package main

import (
	"fmt"
	"runtime"
	"sync"
)


var arr [][]int
var wg sync.WaitGroup
var mutex sync.Mutex

func main(){
	cpuNums := runtime.NumCPU()
	fmt.Println("CPU核心数：",cpuNums)
	runtime.GOMAXPROCS(cpuNums)

	nums := 50000
	wg.Add(nums)
	for i:=0;i<nums;i++{
		go work(i)
	}
	wg.Wait()
	fmt.Println("总数=",len(arr))
}

func work(id int){
	fmt.Println(id)
	mutex.Lock()
	{
		arr = append(arr,sortArr())
	}
	mutex.Unlock()
	defer wg.Done()
}

func sortArr() []int{
	var row [] int
	for i:=0;i<4;i++{
		row = append(row, i)
	}
	return row
}

func getArray2Dict(arrSlice []string) map[string]int{
	var arrMap map[string]int
	arrMap = make(map[string]int)
	for idx,c := range arrSlice {
		key := string(c)
		arrMap[key] = idx
		//fmt.Println(c)
	}
	return arrMap
}
