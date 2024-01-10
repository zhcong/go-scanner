package main

import (
	"fmt"
	"testing"
)

func TestTask_ScanTcp(t *testing.T) {
	testTask := &Task{"192.168.199.1", 80, 0, false}
	err := testTask.ScanTcp()
	if err != nil {
		fmt.Println("error: " + err.Error())
		return
	}
	fmt.Println(testTask)
}
