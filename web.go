package main

import (
	"log"
	"fmt"
	"net/http"
	"strings"
	"html/template"
	"os"
	"io"
	"net"
	_ "github.com/Go-SQL-Driver/MySQL"
	"database/sql"
	"time"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/mgo.v2"
	// "runtime"
)
func ser(w http.ResponseWriter,r *http.Request){

	r.ParseForm()		//参数解析 默认是不解析参数的
	fmt.Println("请求类型",r.Method)
	fmt.Println("表单信息",r.Form)  //打印参数信息到服务端
	fmt.Println("请求url",r.URL.Path) 
	fmt.Println("请求url前缀",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for _, v := range r.Form {

		fmt.Println("value:", strings.Join(v,""))

	}
	// r.ParseMultipartForm(32 << 20)
	// file, _, _ := r.FormFile("file")
	// file.Close()
	// f,_:=os.Create("./test.jpg")  //声明图片名称
	// defer f.Close()
	// io.Copy(f,file)  //将提交的文件数据写入f文件
	// r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file") 
	if err != nil {

		log.Fatal(err)

	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f,err:=os.Create(handler.Filename)  //创建将要写入的文件名称
	defer f.Close()
	io.Copy(f,file)
	
}

func form(w http.ResponseWriter,r *http.Request){

	t,_:=template.ParseFiles("login.gtpl")
	t.Execute(w,nil)

}
/*
	数据库操作
*/
func m(){
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	db,err:=sql.Open("mysql","bys:Xuxian0628.@tcp(116.196.72.23:3306)/bysedu")
	if err!=nil{

		log.Fatal(err)

	}
	//准备一个写入的sql
	stm,err:=db.Prepare("INSERT bys_test SET username=?,real_name=?,age=?,address=?,created_at=?")
	if err!=nil{
		
		log.Fatal(err)

	}
	//添加要写入的内容，并执行
	res, err :=stm.Exec("小白","董杰波",23,"北京市海淀区苏州街长远天地",time.Now().Unix())
	if err!=nil{
		
		log.Fatal(err)

	}
	id,err:=res.LastInsertId()
	fmt.Println(id)
	if err!=nil{
		
		log.Fatal(err)

	}

	rows,err:=db.Query("select * from bys_test")
	if err!=nil{

		log.Fatal(err)

	}
	defer rows.Close()
	for rows.Next(){
		// var id int
		var username string
		var real_name string
		var age string
		var id string
		var address string
		var created_at int
		if err:=rows.Scan(&id,&username,&real_name,&address,&age,&created_at);err!=nil{

			log.Fatal(err)

		}
		fmt.Println(username,real_name,age,address)
	}



}
func r(){

	conn,err:=redis.Dial("tcp","60.205.221.158:6379")

	if err!=nil{

		log.Fatal(err)

	}
	defer conn.Close()
	auth,err:=conn.Do("AUTH","bjbys")
	if err!=nil{

		log.Fatal(err)

	}
	fmt.Println(auth)
	conn.Send("Set","age",2)
	u,err:=conn.Do("Get","age")
	if err!=nil{

		log.Fatal(err)
	}
	fmt.Println(u)
	
}
/*
*	mongo
*/
type Person struct {

	name string
	phone string

}
func mo(){


	// session, err := mgo.Dial("127.0.0.1:27017") 
	// if err != nil {
	// 	panic(err) 
	// }
	// defer session.Close() 
	// c := session.DB("test").C("people")
	// err = c.Insert(&Person{"Ale", "+55 53 8116 9639"},&Person{"Cla", "+55 53 8402 8510"}) 
	// if err != nil {
	// 	panic(err) 
	// }

	session,err:=mgo.Dial("127.0.0.1:27017")
	
	if err!=nil{
		panic(err)
	}
	defer session.Close()
	member:=session.DB("test").C("member")
	err=member.Insert(&Person{"Ale", "+55 53 8116 9639"},&Person{"Cla", "+55 53 8402 8510"})
	if err != nil {
		panic(err) 
	}
	// fmt.Println(mg)
}

type class struct{

	name string
	age int 
	address string

}

func beida(c class) string{

	return "我叫小白"

}

func (c class) beida(b class) string{

	return "我叫小黑"

}


func main(){

	// http.HandleFunc("/hello",ser)
	// http.HandleFunc("/login",form)

	// err:=http.ListenAndServe(":8090",nil)
	// if err!=nil{

	// 	log.Fatal(err)

	// }
	// m()
	// r()
	yiban:=class{"董力",23,"北京市海淀区"}
	y:=beida(yiban)
	x:=yiban.beida(yiban)
	
	mo()
	service := "127.0.0.1:7777"
	tcpAddr, err := net.ResolveUDPAddr("udp", service)
	if err!=nil{
		panic(err)
	}
	fmt.Println(tcpAddr)
	// fmt.Println(runtime.NumCPU())

	fmt.Println(x)
	fmt.Println(y)



}