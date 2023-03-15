package SingletonPattern

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetLazyInstance(t *testing.T) {
	instance1 := GetLazyInstance()
	instance2 := GetLazyInstance()
	if instance1 != instance2 {
		t.Fatal("instance is not equal")
	}
	fmt.Printf("instance1 = %p\n", instance1)
	fmt.Printf("instance2 = %p\n", instance2)
}

func TestParallelLazySingleton(t *testing.T) {
	const parCount = 100
	start := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]LazySingleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			// 协程阻塞，等待channel被关闭才能继续运行
			<-start
			instances[index] = GetLazyInstance()
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
