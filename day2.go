package main

import "fmt"

//func recv(c chan int) {
//	ret := <-c
//	fmt.Println("receive data from chan", ret)
//	<-c
//}
//
//// 管道chan
//func main() {
//	ch := make(chan int) //无缓冲通道,接受和发送同步通信,发送->切换换到接受线程
//	//有缓冲make(chan in,1),主线程非守护线程,直接执行结束退出
//	go recv(ch)
//	ch <- 10
//	ch <- 20
//	fmt.Println("send data into chan")
//}

//func reflectType(inter interface{}) {
//	v := reflect.TypeOf(inter)
//	fmt.Printf("type: %v\n", v)
//	//类型Type和种类Kind
//	fmt.Printf("type: %v, Kind: %v\n", v.Name(), v.Kind())
//
//}
//
//// 反射
//func main() {
//	var a float32 = 3.14
//	reflectType(a)
//}

// WashingMachine 洗衣机
type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct{}

// 实现WashingMachine接口的dry()方法
func (d dryer) dry() {
	fmt.Println("甩一甩")
}

// 海尔洗衣机
type haier struct {
	dryer //嵌入甩干器
}

// 实现WashingMachine接口的wash()方法
func (h haier) wash() {
	fmt.Println("洗刷刷")
}
func main() {
	var inter WashingMachine = haier{} //interface接受struct(类),执行interface方法(其实条用struct绑定方法),实现多态
	inter.wash()
}
