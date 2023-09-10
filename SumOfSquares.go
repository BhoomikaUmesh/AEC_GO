package main
import ("fmt")
func main(){
	var n int
	var sum int
	sum=0
	fmt.Println("enter n")
	fmt.Scan(&n)
	for i:=1;i<=n;i++{
		sum=sum+(i*i)
	}
	fmt.Print(sum)
}