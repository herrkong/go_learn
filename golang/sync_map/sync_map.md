


sync.map  LoadOrStore()


#### 实现协程安全的map
rwmutex

#### 使用go map注意初始化 和并发读写
并发读写 
没有初始化
未初始化的 map 都是 nil，直接赋值会报 panic。
map 作为结构体成员的时候，很容易忘记对它的初始化。 


#### 两种方案 并发读写map
map + mutex 
sync.Map  // 大量读 少量写的 适合


#### sync.Map

Store() // set value
Load()  // 查找key 
LoadOrStore() // no found then store






##### sync.Map的底层实现原理

写：直写。 读：先读read，没有再读dirty。

把dirty提升到read



sync.Map是通过冗余的两个数据结构(read、dirty),实现性能的提升。为了提升性能，load、delete、store等操作尽量使用只读的read；为了提高read的key击中概率，采用动态调整，将dirty数据提升为read；对于数据的删除，采用延迟标记删除法，只有在提升dirty的时候才删除。


type Map struct {
	mu Mutex  //加锁作用。保护后文的dirty字段
	read atomic.Value // readOnly 存读的数据。因为是atomic.Value类型，只读，所以并发是安全的。实际存的是readOnly的数据结构。
	dirty map[interface{}]*entry   //包含最新写入的数据。当misses计数达到一定值，将其赋值给read。
	misses int  //计数作用。每次从read中读失败，则计数+1。
}


// readOnly is an immutable struct stored atomically in the Map.read field.
type readOnly struct {
	m       map[interface{}]*entry
	amended bool // true if the dirty map contains some key not in m.
}




