package main

import (
	"fmt"
	//"time"
	//"log"
	"sync"
)

// 实现协程安全的map
type SyncMap struct{
	data map[string]int
	//lock sync.Mutex
	lock sync.RWMutex
}

// 结构体中包含map 自定义构造函数
func NewSyncMap() * SyncMap{
	return &SyncMap{
		data : make(map[string]int),
		lock : sync.RWMutex{},
	}
}

func(m * SyncMap) GetKey(key string) (int, bool){
	m.lock.RLock()
	defer m.lock.RUnlock()
	if _,ok := m.data[key]; ok {
		return m.data[key],true
	}else{
		fmt.Printf("key=%s not exists\n",key)
		return 0,false
	}
}


func(m * SyncMap) InsertKey(key string, value int ){
	m.lock.Lock()
	defer m.lock.Unlock()
	if _,ok := m.data[key]; !ok {
		m.data[key] = value
	}
}

func(m * SyncMap) DeleteKey(key string ){
	m.lock.Lock()
	defer m.lock.Unlock()
	if _,ok := m.data[key]; ok {
		delete(m.data,key)
	}
}

func PrintMap(m map[string]int){
	for key,value := range m{
		fmt.Printf("key=%s,value=%d\n",key,value)
	}
}

func main3() {
	m := NewSyncMap()
	vec := []string{"darwin","ross","rachel"}

	wg:= sync.WaitGroup{}

	for i:= 0 ; i< 3 ;i++{
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			m.InsertKey(vec[i],i)
		}(i)
	}

	for i:= 0; i < 3 ; i++{
		wg.Add(1)
		go func(i int){
			defer wg.Done()
			m.DeleteKey(vec[i])
		}(i)
	}
	wg.Wait()
	//time.Sleep(time.Second)
	
	PrintMap(m.data)

	// fmt.Print("-------\n")

	// m.DeleteKey("darwin")
	// PrintMap(m.data)

}