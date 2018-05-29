
package main
import (
    "fmt"
    "runtime"
)



//其实这就相当于一个线程池的作用;
func main() {
    const (
        GOROUTINE_COUNT = 10000
        TASK_COUNT      = 1000000
    )
    chReq := make(chan string, GOROUTINE_COUNT)
    chRes := make(chan int, GOROUTINE_COUNT)
   
    runtime.GOMAXPROCS(4)


	//20个线程的线程池;
	for i := 0; i < GOROUTINE_COUNT; i++ {
        go func() {
            
            //记录是第几个线程
            tmp := i  

            for {
                url := <-chReq
                //输出多次，是为了让一个线程占用的时间比较多而已，不要一下就跑完了;
                fmt.Println("这是第几个线程====",tmp, url)
                fmt.Println("这是第几个线程====",tmp, url)
                fmt.Println("这是第几个线程====",tmp, url)
                fmt.Println("这是第几个线程====",tmp, url)
                fmt.Println("这是第几个线程====",tmp, url)
                fmt.Println("这是第几个线程====",tmp, url)
                fmt.Println("这是第几个线程====",tmp, url)
                fmt.Println("这是第几个线程====",tmp, url)
                fmt.Println("这是第几个线程====",tmp, url)
                fmt.Println("这是第几个线程====",tmp, url)
                fmt.Println("这是第几个线程====",tmp, url)
                fmt.Println("这是第几个线程====",tmp, url)
                fmt.Println("这是第几个线程====",tmp, url)
                fmt.Println("这是第几个线程====",tmp, url)
                fmt.Println("这是第几个线程====",tmp, url)
                fmt.Println("这是第几个线程====",tmp, url)
                fmt.Println("这是第几个线程====",tmp, url)
                chRes <- 0
            }
        }()
	}
	

	//这个线程单独向通道中发送数据;
    go func() {
        urls := make([]string, TASK_COUNT)
        for i := 0; i < TASK_COUNT; i++ {
            urls[i] = fmt.Sprintf("http://www.%d.com", i)
        }
        // got urls
        for i := 0; i < TASK_COUNT; i++ {
            chReq <- urls[i]
        }
	}()
	

	//等待接收一百个任务的结束标志后才结束;
    for i := 0; i < TASK_COUNT; i++ {
        d := <-chRes
        // check error
        _ = d
    }
}