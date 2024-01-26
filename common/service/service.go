package service

import "sync"

var (
	once sync.Once
	name string
)

func SetName(n string) {
	once.Do(func() {
		name = n
	})
}

func GetName() string {
	return name
}
