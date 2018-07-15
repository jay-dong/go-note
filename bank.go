package bank
var deposits=make(chan int)
var balances=make(chan int)
func Desposit(amount int)
func Balance() int
func teller(){

	var balance int
	for{
		select{
		case amount:=<-deposits:
			balance+=amount
		case balances<-balance
		}
	}
}
func init(){

	go teller()

}