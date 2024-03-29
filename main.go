package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		println("param empty.")
		return
	}
	ip := os.Args[1]
	maxPost, err := strconv.Atoi(os.Args[2])
	if err != nil {
		println("port error.")
		return
	} // 通信
	toMonitor := make(chan *Task, 10)
	defer close(toMonitor)
	// 计数
	count := sync.WaitGroup{}
	// 锁
	lock := sync.Mutex{}
	for port := 1; port != maxPost; port++ {
		count.Add(1)
		port := port
		go func() {
			defer count.Done()
			task := &Task{ip, port, 0, false}
			err := task.ScanTcp()
			if err != nil {
				fmt.Println("error: " + err.Error())
				return
			}
			lock.Lock()
			toMonitor <- task
			lock.Unlock()
		}()
		time.Sleep(10 * time.Millisecond)
	}
	// 结果集合
	result := make([]*Task, 0, maxPost)
	// 监控线程
	go func() {
		for {
			temp, err := <-toMonitor
			if !err {
				continue
			}
			result = append(result, temp)
			if len(result) == maxPost {
				return
			}
		}
	}()
	count.Wait()
	// 遍历结果
	for _, task := range result {
		if task.IsLink {
			fmt.Printf("port[%d] open, cost [%d]s \n", task.Port, task.LinkSec)
		}
	}
}
