package injector

import (
	"fmt"
	"reflect"
	"sync"
)

type Injector struct {
	registry map[reflect.Type]any
	lock     sync.RWMutex
}

var (
	instance *Injector
	once     sync.Once
)

func GetInjector() *Injector {
	once.Do(func() {
		instance = &Injector{
			registry: make(map[reflect.Type]any),
		}
	})

	return instance
}

func (i *Injector) Register(interfaceType any, implementation any) {
	i.lock.Lock()
	defer i.lock.Unlock()

	ifaceType := reflect.TypeOf(interfaceType).Elem()
	i.registry[ifaceType] = implementation
}

func (i *Injector) Resolve(interfaceType any) any {
	i.lock.RLock()
	defer i.lock.RUnlock()

	ifaceType := reflect.TypeOf(interfaceType).Elem()
	impl, exists := i.registry[ifaceType]
	if !exists {
		panic(fmt.Sprintf("No registered implementation for %v", ifaceType))
	}
	return impl
}

func Register[T any](impl T) {
	GetInjector().Register((*T)(nil), impl)
}

func Resolve[T any]() T {
	return GetInjector().Resolve((*T)(nil)).(T)
}
