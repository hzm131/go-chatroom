package main

import "fmt"

func writeData(intChan chan int){
	for i:=1;i<=50;i++{
		intChan <- i
	}
	close(intChan) //写入50个后关闭管道  但是不影响读
}

func readData(intChan chan int,exitData chan bool){
	for {
		v,ok := <- intChan
		if !ok{
			break
		}
		fmt.Printf("readData 读取到数据=%v\n",v)
	}
	exitData<- true
	close(exitData)
}


func main(){
	intData := make(chan int,50)
	exitData := make(chan bool,1)

	go writeData(intData)
	go readData(intData,exitData)

	for {
		_,ok := <- exitData
		if !ok{
			break
		}
	}
}
