package singleton

// 单例模式
// 懒汉式
//type singleton struct {
//}
//var instance *singleton
//func GetSingleton() *singleton {
//	if instance == nil {
//		instance = &singleton{}
//	}
//	return instance
//}

// 饿汉式
//type singleton struct {
//}
//var instance *singleton = &singleton{}
//func GetSingleton() *singleton {
//	return instance
//}

// 并发，加锁
//type singleton struct{}
//
//var instance *singleton
//var mu sync.Mutex
//
//func GetSingleton() *singleton {
//	mu.Lock()
//	defer mu.Unlock()
//	if instance == nil {
//		instance = &singleton{}
//	}
//	return instance
//}

// 双重锁定 check-lock-check
import "sync"

type singleton struct {
}

var instance *singleton
var mu sync.Mutex

func GetSingleton() *singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = &singleton{}
		}
	}
	return instance
}

// once的写法--双重检查锁
//import "sync"
//
//type singleton struct{}
//
//var instance *singleton
//var once sync.Once
//
//func GetSingleton() *singleton {
//	once.Do(func() {
//		instance = &singleton{}
//	})
//	return instance
//}
