package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//func main() {
//	fmt.Println("Please visit http://127.0.0.1:12345/")
//	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
//		s := fmt.Sprintf("你好，世界!-- Time:%s", time.Now().String())
//		fmt.Fprintf(w, "%v\n", s)
//		log.Printf("%v\n", s)
//	})
//	if err := http.ListenAndServe(":12345", nil); err != nil{
//		log.Fatal("ListenAndServe:", err)
//	}
//}

//// 具名函数
//func Add(a, b int) int {
//	return a + b
//}

//// 匿名函数
//var Add = func(a, b int) int {
//	return a + b
//}

//// 多个参数和多个返回值
//func Swap(a, b int) (int , int ){
//	return b , a
//}

//// 可变数量的参数
//// more 对应 []int 切片类型
//func Sum(a int, more ...int) int {
//	for _, v := range more{
//		a += v
//	}
//	return a
//}
//
//func main() {
//	var a = []interface{}{123, "abc"}
//
//	Print(a...)
//	Print(a)
//}
//
//func Print(a ...interface{}){
//	fmt.Println(a...)
//}

//
//func Find(m map[int]int, key int) (value int, ok bool) {
//	value, ok = m[key]
//	return
//}
//
//func Inc() (v int){
//	defer func() {v ++ }()
//	return 42
//}

////数组定义的方式
//var a [3]int
//var b = [...]int{1, 2, 3}
//var c = [...]int{2: 3, 1: 2}
//var d = [...]int{1, 2, 4: 5, 6}
//
//func main() {
//	fmt.Println(a)
//	fmt.Println(b)
//	fmt.Println(c)
//	fmt.Println(d)
//}
//
//func main() {
//	var a = [...]int{1, 2, 3}
//	var b = &a
//	fmt.Println(a[0], a[1])
//	fmt.Println(b[0], b[1])
//
//	for i ,v := range b {
//		fmt.Println(i, v)
//	}
//
//	for i := range a {
//		fmt.Printf("a[%d]: %d\n", i, a[i])
//	}
//
//	for i ,v := range b {
//		fmt.Printf("b[%d]: %d\n", i, v)
//	}
//
//	for i := 0; i < len(b) ; i ++ {
//		fmt.Printf("c[%d]: %d\n", i, b[i])
//	}
//
//
//}
//
//func main() {
//	var times [5][0]int // 二维数组
//	for range times {
//		fmt.Println("hello")
//	}
//}

//var d [0]int
//var e = [0]int{}
//var f = [...]int{}
//
//
//func main() {
//	//s := "hello, world"
//	//hello := s[:5]
//	//world := s[7:]
//	//
//	//s1 := "hello, world"[:5]
//	//s2 := "hello, world"[7:]
//	//fmt.Println(hello)
//	//fmt.Println(world)
//	//fmt.Println(s1)
//	//fmt.Println(s2)
//
//	//fmt.Printf("%#v\n", []byte("hello, 世界"))
//	for i, c := range "\xe4\x00\x00\xe7\x95\x8cabc" {
//		fmt.Println(i, c)
//	}
//
//	fmt.Println("_____________________")
//
//	for i, c := range []byte("世界abc") {
//		fmt.Println(i, c)
//	}
//
//	fmt.Println("_____________________")
//
//	fmt.Printf("%#v\n", []rune("世界"))
//
//	fmt.Printf("%#v\n", string([]rune{'世', '界'}))
//}

//var (
//	a []int  // nil切片, 和nil相等， 一般用来表示一个不存在的切片
//
//	b = []int{} // 空切片， 和nil不相等， 一般用来表示一个空的集合
//
//	c = []int{1, 2, 3} //有3个元素的切片，len和cap都为3
//
//	d = c[:2] // 有两个元素的切片， len为2 cap为3
//
//	e = c[0:2:cap(c)] // 有两个元素的切片， len为2 cap为3
//
//	f = c[:0] // 有0个元素切片，len为0 cap为3
//
//	g = make([]int, 3) //有3个元素的切片。 len和cap都为3
//
//	h = make([]int, 2, 3) //有2个元素的切片， len为2， cap为3
//
//	i = make([]int, 0, 3) //有0个元素的切片， len为0 cap为3
//	// len函数返回切片中有效的元素长度， cap函数返回切片的容量大小， 容量必须大于或等于切片的长度
//)

//func main() {
//	var a []int
//	a = append(a, 1)
//	a = append(a, 1,2,3,4,5)
//	a = append(a, []int{1, 2, 3}...)
//	// 容量不足的情况，append的操作会重新分配内存，可能导致巨大的内存分配和复制数据代价
//	fmt.Println(a)
//
//	a = append([]int{0}, a...)
//	a = append([]int{-3, -2, -1}, a...)
//	fmt.Println(a)
//}

//func main() {
//
//
//	//// 在切片中间插入数据
//	//var a = []int {5,6,7,8,9}
//	//a = append(a[:1], append([]int{10}, a[1:]...)...)
//	//fmt.Println(a)
//	//a = append(a[:1], append([]int{1,2,3}, a[1:]...)...)
//	//fmt.Println(a)
//
//	//var a = []int {5,6,7,8,9}
//	//a = append(a, 0)
//	//copy(a[1+1:], a[1:])
//	//a[1] = 10
//	//
//	//fmt.Println(a)
//
//	// 删除切片元素
//
//
//
//}

// 删除切片元素
//func main() {
//	// 删除尾部元素
//	a := []int{1,2,3}
//	a = a[:len(a)-1]
//	fmt.Println(a)
//	a = a[:len(a)-2]
//	fmt.Println(a)
//
//	// 删除开头的元素
//	b := []int{1,2,3}
//	//b = b[1:]
//	fmt.Println(b)
//	b = b[2:]
//	fmt.Println(b)
//
//
//
//}

//// 不导致内存空间结构变化 删除元素
//func main() {
//	a := []int{1, 2, 3}
//	//a = append(a[:0], a[1:]...) // 删除开头1个元素
//	fmt.Println(a)
//	a = append(a[:0], a[2:]...) // 删除开头任意一个元素
//}

//// 使用copy完成删除开头的元素
//func main() {
//	a := []int{1, 2, 3}
//	a = a[:copy(a, a[1:])] // 删除开头1个元素
//	a = a[:copy(a, a[N:])] // 删除开头N个元素
//}

//func TrimSpace(s []byte) []byte {
//	b := s[:0]
//	for _, x := range s {
//		if x != ' ' {
//			b = append(b, x)
//		}
//	}
//	return b
//}

//var a = []float64{4, 2, 5,7,2,1,88, 1}
//
//func SortFloat64FastV1(a []float64){
//
//	// 强制类型转换
//	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]
//
//	// 以int方式给float64排序
//	sort.Ints(b)
//
//}
//
//func SortFloat64FastV2(a []float64){
//
//	// 通过reflect.SliceHeader 更新切片头部信息实现转换
//	var c []int
//	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
//	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
//	*cHdr = *aHdr
//	// 以int方式给float64排序
//	sort.Ints(c)
//}

//var a = []float64{4, 2, 5, 7, 2, 1, 88, 1}
//
//func SortFloat64FastV1(a []float64){
//
//	// 强制类型转换
//	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]
//
//	// 以int方式给float64排序
//	sort.Ints(b)
//
//}
//func SrtFloat64FastV2(a []float64){
//	// 通过reflect.SliceHeader 更新切片头部信息实现转换
//	var c []int
//	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
//	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))
//	*cHdr = *aHdr
//	sort.Ints(c)
//
//}

//func main() {
//	for i := 0; i < 3; i ++ {
//		i := i
//		defer func(){ println(i)}()
//	}
//}

//func main() {
//	for i := 0; i < 3; i ++ {
//		defer func(i int){println(i)}(i)
//
//	}
//}

//f, _ := OpenFile("foo.dat")
//
//// 绑定到f对象
//// func Close() error
//var Close = func() error {
//	return (*File).Close(f)
//}
//
//var Read = func(offset int64, data []byte) int {
//	return (*File).Read(f, offset, data)
//}
//
//func main() {
//	// 文件处理
//	Read(0, data)
//	Close()
//}

//f, _ := OpenFile("foo.dat")
//
//var Close = f.Close
//
//var Read = f.Read
//
//Read(0, data)
//Close()

//var (
//	a.io.ReadCloser = ( *os.File)(f)
//)

//func Fprintf(w io.Writer, format string, args ...interface{}) (i
//nt, error)

//
//var total struct{
//	sync.Mutex
//	value int
//}
//
//func worker(wg *sync.WaitGroup){
//	defer wg.Done()
//	for i :=0 ; i <=100; i++ {
//		total.Lock()
//		total.value += i
//		total.Unlock()
//	}
//}
//func main() {
//	var wg sync.WaitGroup
//	wg.Add(3)
//	go worker(&wg)
//	go worker(&wg)
//	go worker(&wg)
//	wg.Wait()
//
//	fmt.Println(total.value)
//}

//var total uint64
//
//func worker(wg *sync.WaitGroup) {
//	defer wg.Done()
//	var i uint64
//	for i = 0; i <= 100; i++ {
//		atomic.AddUint64(&total, i)
//	}
//}
//
//
//func main() {
//	var wg sync.WaitGroup
//	wg.Add(2)
//	go worker(&wg)
//	go worker(&wg)
//	wg.Wait()
//}

//var a string
//var done bool
//
//func setup() {
//	a = "hello world"
//	done = true
//}
//
//func main() {
//	go setup()
//	for !done {}
//	print(a)
//}

////不能确定是先打印还是main函数先执行结束
//func main() {
//	go println("世界が終わる")
//}

//func main() {
//	done := make(chan int)
//
//	go func(){
//		println("世界が終わる")
//		done <- 1
//	}()
//
//	<-done
//
//}

//// Mutex进行加锁，互斥锁
//func main() {
//	var mu sync.Mutex
//
//	mu.Lock()
//	go func(){
//		println("アスカ")
//		mu.Unlock()
//	}()
//
//	mu.Lock()
//}

//var a string
//
//func f (){
//	print(a)
//}
//
//func hello(){
//	a = "hello、アスカ"
//	go f()
//}
//func main() {
//	hello()
//}

//var done = make(chan bool)
//var msg string
//func aGoroutine(){
//	msg = "お早う、アスカ"
//	done <- true
//}
//func main() {
//	go aGoroutine()
//	<-done
//	println(msg)
//}

//?
//var limit = make(chan int, 3)
//var work = []func(){
//	func() { println("1"); time.Sleep(1 * time.Second) },
//	func() { println("2"); time.Sleep(1 * time.Second) },
//	func() { println("3"); time.Sleep(1 * time.Second) },
//	func() { println("4"); time.Sleep(1 * time.Second) },
//	func() { println("5"); time.Sleep(1 * time.Second) },
//}
//
//
//func main() {
//	for _, w := range work {
//		go func(w func()) {
//			limit <- 1
//			w()
//			<-limit
//		}(w)
//	}
//	select{}
//}

//func main() {
//	go println("お早う、世界、大好き、明日香ちゃん")
//	time.Sleep(time.Second)
//}

//func main() {
//	var mu sync.Mutex
//	mu.Lock()
//	go func() {
//		fmt.Println("こうにちは。世界")
//		mu.Unlock()
//	}()
//	mu.Lock()
//}

//func main() {
//	done := make(chan int, 1)
//
//	go func(){
//		fmt.Println("こうにちは。世界")
//		done <- 1
//	}()
//	<-done
//}

//func main() {
//	done := make(chan int, 10) // 带10个缓存
//
//	// 开启n个后台打印线程
//	for i := 0; i < cap(done); i ++ {
//		go func () {
//			fmt.Println("お早う、世界")
//			done <- 1
//		}()
//	}
//
//	// 等待N个后台线程完成
//	for i := 0; i< cap(done); i++{
//		<-done
//	}
//}
//
//func main() {
//	var wg sync.WaitGroup
//
//	// 开启N个后台打印线程
//	for i := 0; i < 10 ; i ++ {
//		wg.Add(1)
//		go func(){
//			fmt.Println("お早う、世界")
//			wg.Done()
//		}()
//	}
//	wg.Wait()
//}

//func main() {
//	var s , sep string
//	for i := 1; i < len(os.Args); i ++ {
//		s += sep + os.Args[i]
//		sep = " "
//	}
//	fmt.Println(s)
//}

//func main() {
//	s, sep := "", " "
//	for _, arg := range os.Args[1:]{
//		s += sep +arg
//		sep = " "
//	}
//	fmt.Println(s)
//}
//
//func main() {
//	fmt.Println(strings.Join(os.Args[1:], " "))
//}

//package main
//
//import (
//"fmt"
//"os"
//)
//
//func main() {
//	var s, sep string
//	for i := 1; i < len(os.Args); i++ {
//		s += sep + os.Args[i]
//		sep = " "
//	}
//	fmt.Println(s)
//}

//func main() {
//	counts := make(map[string]int)
//	input := bufio.NewScanner(os.Stdin)
//	for input.Scan() {
//		counts[input.Text()]++
//	}
//	// NOTE: ignoring potential errors from input.Err()
//	for line, n := range counts {
//		if n > 1 {
//			fmt.Printf("%d\t%s\n", n, line)
//		}
//	}
//}

//func main() {
//	counts := make(map[string]int)
//	files := os.Args[1:]
//	if len(files) == 0{
//		countLines(os.Stdin, counts)
//	} else{
//		for _,args := range files {
//			f, err := os.Open(args)
//			if err != nil {
//				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
//				continue
//			}
//			countLines(f, counts)
//			f.Close()
//		}
//	}
//	for line, n := range counts {
//		if n > 1 {
//			fmt.Printf("%d\t%s\n", n, line)
//		}
//	}
//}
//
//func countLines(f *os.File, counts map[string]int){
//	input := bufio.NewScanner(f)
//	for input.Scan(){
//		counts[input.Text()]++
//	}
//}

//func main() {
//	counts :=make(map[string]int)
//	for _, filename := range os.Args[1:]{
//		data ,err := ioutil.ReadFile(filename)
//		if err != nil {
//			fmt.Fprintf(os.Stderr, "dups:%v\n", err)
//			continue
//		}
//		for _, line := range strings.Split(string(data), "\n"){
//			counts[line]++
//		}
//	}
//	for line, n := range counts {
//		if n > 1 {
//			fmt.Printf("%d\t%s\n", n, line)
//		}
//	}
//}

//var palette = []color.Color{color.White, color.Black}
//
//const (
//	whiteIndex = 0
//	blackIndex = 1
//)
//
//func main() {
//	rand.Seed(time.Now().UTC().UnixNano())
//	lissajous(os.Stdout)
//}
//
//func lissajous(out io.Writer){
//	const (
//		cycles = 5
//		res = 0.001
//		size = 100
//		nframes = 64
//		delay = 8
//	)
//	freq := rand.Float64() * 3.0
//	anim := gif.GIF{LoopCount:nframes}
//	phase := 0.0
//	for i := 0;i <nframes ; i ++ {
//		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
//		img := image.NewPaletted(rect, palette)
//		for t := 0.0; t < cycles *2 * math.Pi; t += res {
//			x := math.Sin(t)
//			y := math.Sin(t*freq + phase)
//			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),blackIndex)
//		}
//		phase += 0.1
//		anim.Delay = append(anim.Delay, delay)
//		anim.Image = append(anim.Image, img)
//	}
//	gif.EncodeAll(out, &anim)
//}

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}

}
