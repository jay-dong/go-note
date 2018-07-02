package main
import(
	"fmt"
	"net/http"
	"net/url"
	"time"
	"log"
)

func main(){

	// res,err:=http.Get("https://www.baidu.com")
	//发送post请求
	res,err:=http.PostForm("https://bys.ncwc.com.cn/admin/login",url.Values{"logname":{"jay_dong@outlook.com"},"logpass":{"1234"}})
	if err!=nil{

		fmt.Println(err)

	}
	syshello:=func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q")
	}

	s:=&http.Server{

		Addr:":8080",
		Handler:syshello,
		ReadTimeout:10 * time.Second,
		WriteTimeout:10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())


	fmt.Println(res.Location)
	defer func() {fmt.Println("关闭")} ()
	defer res.Body.Close()


}