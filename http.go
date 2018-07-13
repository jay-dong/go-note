package main
import(
	"fmt"
	// "net/http"
	// "net"
	// "io"
	// "net/url"
	// "time"
	"runtime"
	"sync"
	// "log"
)

// func main(){

// 	// res,err:=http.Get("https://www.baidu.com")
// 	//发送post请求
// 	res,err:=http.PostForm("https://bys.ncwc.com.cn/admin/login",url.Values{"logname":{"jay_dong@outlook.com"},"logpass":{"1234"}})
// 	if err!=nil{

// 		fmt.Println(err)

// 	}
// 	syshello:=func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello, %q")
// 	}

// 	s:=&http.Server{

// 		Addr:":8080",
// 		Handler:,
// 		ReadTimeout:10 * time.Second,
// 		WriteTimeout:10 * time.Second,
// 		MaxHeaderBytes: 1 << 20,
// 	}
// 	log.Fatal(s.ListenAndServe())


// 	fmt.Println(res.Location)
// 	defer func() {fmt.Println("关闭")} ()
// 	defer res.Body.Close()
//}
func main(){

	runtime.GOMAXPROCS(2)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("y: ", i)
			fmt.Println("我先执行的")
			wg.Done()
		}(i)
	}

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i: ", i)
			fmt.Println("我先执行的")
			wg.Done()
		}(i)
	}
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("x: ", i)
			wg.Done()
		}()
	}
	wg.Wait()

	// one:=make(chan int)
	// two:=make(chan int)

	// go func(){

	// 	for x := 0; x < 10; x++{
	// 		one <- x
	// 	}
	// 	close(one)
		
	// }()

	// go func(){

	// 	for x:=range one{
	// 		two<- x*x
	// 	}
	// 	// fmt.Println(cap(one))
	// 	close(two)
	// }()

	// for x:= range two {

	// 	fmt.Println(x)

	// }



	// conn,err:=net.Dial("tcp","www.baidu.com:80")
	// if err!=nil {

	// 	log.Fatal(err)

	// }
	// fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	
	// status, err := bufio.NewReader(conn).ReadString('\n')


	// listener,err:=net.Listen("tcp",":8000")
	// if err!=nil{

	// 	log.Fatal(err)		

	// }

	// for{

	// 	conn,err:=listener.Accept()
	// 	if err!=nil{
	// 		log.Print(err)
	// 		continue
	// 	}
	// 	go handleConn(conn)

	// }
}
// func handleConn(c net.Conn){
// 	defer c.Close()
// 	for{
// 		_,err:=io.WriteString(c,time.Now().Format("15:04:06\n"))
// 		if err !=nil{
// 			return 
// 		}
// 		time.Sleep(1 * time.Second)
// 	}

// }