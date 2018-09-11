### Goroutines的定义

在Go中定义一个协程，只需要早方法前面加上关键字 `go`,那么这个方法就成为了一个协程方法，有下面几种方式  

* 在方法前面加上关键字 `go`

```
func main(){
    go sayHello()
}
func sayHello(){
    fmt.Println("hello")
}
```

* 在匿名函数前面加上关键字 `go` 

```
func main(){
    go func(){
        fmt.Println("hello")
    }() //匿名方法的协程，必须立即执行
}
```

* 将匿名方法复制给一个变量，并使用关键字 `go` 调用 

```
func main(){
    sayHello := func(){
        fmt.Println("hello")
    }

    go sayHello()
}
```

但是上面的这个段代码在实际运行时，并不会输出任何东西,因为`sayHello`方法并不会阻塞`main`方法的运行，`sayHello`会在自己的协程中运行，但是`sayHello`方法并没有获取机会运行，`main`方法就已经退出运行了，所以上面的代码并不能输出任何信息，为了能够使上面的代码输出正确的信息，使用`sync`包，改造一下  

```
func main() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello")
	}
	wg.Add(1)
	go sayHello()
	wg.Wait() //在main方法中设置一个阻塞点
}
```

#### 在协程中运行一个闭包，在闭包中操作的是变量的引用还是变量的copy?下面的代码会输出上面?

```
func main() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome" //变量的值改变
	}()
	wg.Wait()
	fmt.Println(salutation)
}
```

最终的输出是  

```
welcome
```

在考虑一下下面的代码会输出什么?

```
func main(){
	var wg1 sync.WaitGroup
	for _, val := range []string{"hello", "greetings", "good day"} {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			fmt.Println(val)
		}()
	}
	wg1.Wait()
}
```

输出结果是  

```
good day
good day
good day
```

改造一下，输数在slice中所有的字符串  

```
func main(){
    var wg1 sync.WaitGroup
	for _, val := range []string{"hello", "greetings", "good day"} {
		wg1.Add(1)
		go func(v string) { // 闭包是一个有参数的闭包
			defer wg1.Done()
			fmt.Println(v)
		}(val) // 传入需要的参数
	}
	wg1.Wait()
}
```

输出的结果  

```
good day
hello
greetings
```

#### 协程的大小

协程较于其他的线程的优势是轻量，在`Go`中一个协程的初始大小只有几KB，如果不够，会在运行时自动增加，这样可以运行成千上万个协程，通常的线程的数量会很少，怎么样测量一个协程的大小?

```
func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() {
		wg.Done()
		<-c
	}

	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}
```

#### sync包

#### waitGroup

如果你不关心并行操作的结果或者你有其他的方式获得并行操作的结果，那么`waitGroup`对一系列并行操作是非常有效的，基本的操作如下  

```
func main() {
	var wg sync.WaitGroup

	wg.Add(1) 
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutines sleeping...")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutines sleeping...")
		time.Sleep(2)
	}()

	wg.Wait()
	fmt.Println("All goroutimes complete.")
}
```

* Add()  //参数1说明有一个协程开始运行
* Done() // defer 保证在协程结束时，调用Done()方法
* Wait() // 阻塞主协程

可以认为`waitGroup`是一个线程安全的计数器，`Add`计数器会加上所传的参数，`Done`会是计数器减一，`Wait`阻塞等待计数器变为0  

#### Mutex和RWMutex

`Mutex`提供了一个中方式是协程能够同步的方位共享资源，基本的`Mutex`操作  

```
func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count += 1
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count -= 1
		fmt.Printf("Decrementing: %d\n", count)
	}

	var arithmetic sync.WaitGroup

	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("Arithmetic complete.",count)
}
```

* Lock()    // 对需要共享的资源加锁
* Unlock()  // 访问完共享资源之后，解锁，使用关键字 `defer`可以保证，总会被执行，即使发生了`panic`,如果没有正确的解锁，那么就有可能造成死锁

#### Cond

#### Once

先看一下下面的代码会输出什么？

```
func main() {
	var count int
	increment := func() {
		count += 1
	}

	var once sync.Once
	var increments sync.WaitGroup
	increments.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer increments.Done()
			once.Do(increment)
		}()
	}
	increments.Wait()
	fmt.Printf("Count is %d\n", count)
}
```

输出结果

```
Count is 1
```

`Once` 顾名思义就是只执行一次，看一下`Once`的源码  

`Once`结构  

```
type Once struct {
	m    Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	// Slow-path.
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
```

再看看下面的代码,会输出什么?

```
func main() {
	var count int
	increment := func() {
		count += 1
	}
	decrement := func() {
		count -= 1
	}

	var once sync.Once
	once.Do(increment)
	once.Do(decrement)
	fmt.Printf("Count: %d\n", count)
}
```

输出结果  

```
Count: 1
```

下面的代码  

```
	var onceA, onceB sync.Once
	var initB func()
	initA := func() {
		onceB.Do(initB)
	}
	initB = func() {
		onceA.Do(initA)
	}
	onceA.Do(initA)
```

`死锁`  


#### Pool