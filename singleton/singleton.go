package singleton

import "sync"

var singleton *Singleton
var once sync.Once

type Singleton struct {
	data int
}

func GetSingleton() *Singleton {
	once.Do(func() {
		singleton = &Singleton{100}
	})
	return singleton
}
