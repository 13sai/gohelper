package gohelper_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/13sai/gohelper"
)

//主函数
func TestNewTask(te *testing.T) {
	//创建一个Task
	t := gohelper.NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})

	//创建一个协程池,最大开启个协程worker
	p := gohelper.NewPool(10)

	//开一个协程 不断的向 Pool 输送打印一条时间的task任务
	go func() {
		i := 0
		for i < 5 {
			p.EntryChannel <- t
			time.Sleep(time.Duration(1) * time.Second)
			i++
		}
		p.Close()
	}()

	//启动协程池p
	p.Run()
}
