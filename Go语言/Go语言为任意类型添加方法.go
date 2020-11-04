package Go语言

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//// 将int定义为MyInt类型
//type MyInt int
//
//// 为MyInt添加IsZero（）方法
//func (m MyInt) IsZero() bool {
//	return m == 0
//}
//// 为MyInt添加Add()方法
//func (m MyInt) Add(other int) int {
//	return other + int(m)
//}
//
//func main() {
//
//	var b MyInt
//
//	fmt.Println(b.IsZero())
//
//	b = 1
//
//	fmt.Println(b.Add(2))
//}

func main() {
	client := &http.Client{}

	// 创建一个http请求
	req, err := http.NewRequest("POST", "http://163.com/", strings.NewReader("key=value"))

	// 发现错误就打印并退出
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	// 为标头添加信息
	req.Header.Add("User-Agent", "myClient")

	// 开始请求
	resp, err := client.Do(req)

	// 处理请求的错误
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	defer resp.Body.Close()
}

//func main() {
//	fmt.Println(time.Hour.String())
//}

//func (d Duration) String() string{
//	// 一系列生成buf的代码
//	return string(buf[w:])
//}
