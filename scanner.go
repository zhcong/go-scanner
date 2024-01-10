package main

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"
)

const WAIT_TIME time.Duration = 5 * time.Second

type Task struct {
	Ip      string
	Port    int
	LinkSec int64
	IsLink  bool
}

func (task *Task) ScanTcp() (err error) {
	if task.Port <= 0 {
		return errors.New("port error.")
	}
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	connectTcp(task)
	return nil
}

func connectTcp(task *Task) {
	startTime := time.Now().Unix()
	listenStr := task.Ip + ":" + strconv.Itoa(task.Port)
	listen, err := net.DialTimeout("tcp", listenStr, WAIT_TIME)
	defer func() {
		err := listen.Close()
		if err != nil {
			fmt.Println("close error.")
		}
		return
	}()
	if err != nil {
		panic(err)
		return
	}
	task.IsLink = true
	task.LinkSec = time.Now().Unix() - startTime
}
