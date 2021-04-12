package gohelper

import (
	"fmt"
)

// 定义任务Task类型,每一个任务Task都可以抽象成一个函数
type Task struct {
	f func() error
}

func NewTask(f func() error) *Task {
	t := Task{
		f: f,
	}

	return &t
}

// 执行Task任务的方法
func (t *Task) Execute() {
	t.f()
}

// 定义池类型
type Pool struct {
	// 协程池最大worker数量,限定Goroutine的个数
	size int

	// 对外接收Task的入口
	EntryChannel chan *Task

	// 协程池内部的任务就绪队列
	JobsChannel chan *Task
}

// 创建一个协程池
func NewPool(size int) *Pool {
	p := Pool{
		size:         size,
		EntryChannel: make(chan *Task),
		JobsChannel:  make(chan *Task),
	}

	return &p
}

// 协程池创建一个worker并且开始工作
func (p *Pool) worker(work_ID int) {
	// worker不断的从JobsChannel内部任务队列中拿任务
	for task := range p.JobsChannel {
		// 如果拿到任务,则执行task任务
		task.Execute()
		fmt.Println("worker ID ", work_ID, " 执行完毕任务")
	}
}

// 让协程池Pool开始工作
func (p *Pool) Run() {
	// 首先根据协程池的worker数量限定,开启固定数量的Worker,
	for i := 0; i < p.size; i++ {
		go p.worker(i)
	}

	// 从EntryChannel协程池入口取外界传递过来的任务
	// 并且将任务送进JobsChannel中
	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}
}

func (p *Pool) Close() {
	close(p.EntryChannel)
	close(p.JobsChannel)
}
