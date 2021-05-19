package creation

import (
	"sync"
	"sync/atomic"
)

type Singleton struct {
	name string
}

// ------------------------ 饿汉式 start --------------------
var singletonHungry *Singleton

func init() {
	singletonHungry = &Singleton{name: "hungry"}
}

func GetInstanceHungry() *Singleton {
	return singletonHungry
}

// ------------------------ 饿汉式  end --------------------

// ------------------------ 懒汉式  start ------------------
var (
	singletonLazy *Singleton
	epoch         int64
)

func GetInstanceLazy() *Singleton {
	if singletonLazy == nil && atomic.AddInt64(&epoch, 1) == 1 {
		singletonLazy = &Singleton{name: "lazy"}
	}
	return singletonLazy
}

// ------------------------ 懒汉式  end ------------------

// ------------------------ 懒汉式 sync.Once实现 start ------------------
var (
	singletonLazyII *Singleton
	once            sync.Once
)

func GetInstanceLazyII() *Singleton {
	if singletonLazyII == nil {
		once.Do(func() {
			singletonLazyII = &Singleton{name: "lazy II"}
		})
	}
	return singletonLazyII
}

// ------------------------ 懒汉式 sync.Once实现 end ------------------
