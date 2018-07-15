package main
import ( 
	"fmt" 
	"net"
	"os"
 )
func main() {
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service) 
	checkError(err)
	conn, err := net.ListenUDP("udp4", udpAddr) 
	checkError(err)
	for {
		handleClient(conn) 
	}
}
func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return 
	}
	conn.WriteToUDP([]byte("最近好吗")) 
}
func checkError(err error) { 
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1) 
	}
}
