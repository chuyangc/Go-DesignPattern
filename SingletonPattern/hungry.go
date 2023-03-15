package SingletonPattern

// 饿汉模式

// HungrySingleton 通过该接口避免 GetInstance返回一个包私有类型的指针
type HungrySingleton interface {
	hungry()
}

type hungrysingleton struct {
}

func (s *hungrysingleton) hungry() {

}

var singletonInstance = &hungrysingleton{}

// GetHungryInstance 用于获取单例模式对象
func GetHungryInstance() HungrySingleton {
	return singletonInstance
}
