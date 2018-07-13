package main

import (
	"time"
	"os"
	"fmt"
	"net"
)
type addinfo struct{

	ip string
	name string

}
/*
	tcp
*/

func main(){

	x:=addinfo{":8090","hello wolrd"}
	ip,err:=net.ResolveTCPAddr("tcp4",x.ip)
	listen,err:=net.ListenTCP("tcp",ip)
	// listen.SetDeadline(deal)
	if err!=nil{

		fmt.Println(err)

	}
	i:=0  //初始化服务端接入数量
	//阻塞  等待连接接入
	for {

		con,err:=listen.Accept()   // 同意接受链接请求
		if err!=nil{

			fmt.Fprintln(os.Stdout,"连接失败的错误原因:%v",err)
			continue		//跳出本次连接请求

		}
		// defer con.Close()		//关闭链接
		//声明一个用存储的链接地址的map
		go hanlerRes(con)
		i++
	}

}

func hanlerRes(c net.Conn){

	date:="北京时间"+time.Now().Format("2006-01-02 15:04:05")
	c.Write([]byte(date))
	buf := make([]byte,1024)  //声明读取返回数据容器 
	fmt.Println("接入地址:",c.RemoteAddr())  //返回远程链接地址返回远程链接地址
	
	//读取用户输入信息
	for{
		msg,err:=c.Read(buf)
		if err!=nil{

			panic(err)

		}

		fmt.Println(c.RemoteAddr(),"发送的信息:",string(buf[:msg]))
		
	}

}
/*

	udp连接

*/
func udp(c net.Conn){



}


/*
	面向数据包的网络连接
*/
func pack(c net.PacketConn){

	


}





