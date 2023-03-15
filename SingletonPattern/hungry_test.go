package SingletonPattern

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	instance1 := GetHungryInstance()
	instance2 := GetHungryInstance()
	if instance1 != instance2 {
		t.Fatal("instance is not equal")
	}
	fmt.Printf("instance1 = %p\n", instance1)
	fmt.Printf("instance2 = %p\n", instance2)
}

func TestParallelHungrySingleton(t *testing.T) {
	const parCount = 100
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]HungrySingleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			// 协程阻塞，等待channel被关闭才能继续运行
			<-start
			instances[index] = GetHungryInstance()
			wg.Done()
		}(i)
	}
	// 关闭channel，所有协程同时开始运行，实现并行
	close(start)
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
