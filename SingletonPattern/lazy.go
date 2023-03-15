package SingletonPattern

import "sync"

// 懒汉模式

// LazySingleton 通过该接口避免 GetInstance返回一个包私有类型的指针
type LazySingleton interface {
	lazy()
}

type lazysingleton struct {
}

func (s *lazysingleton) lazy() {

}

var (
	instance *lazysingleton
	once     sync.Once
)

// GetLazyInstance 用于获取单例模式对象
func GetLazyInstance() LazySingleton {
	once.Do(func() {
		instance = &lazysingleton{}
	})
	return instance
}
