package main
import "fmt"
func isp(n int)bool{
    or:=n
    rev:=0
    for n>0{
        rem:=n%10
        rev=rev*10+rem
        n=n/10
    }
    return or==rev
}
func main(){
    var limit int
    fmt.Println("enter limit")
	fmt.Scan(&limit)
	lp:=0
    var m1,m2 int
    for i:=limit;i>=0;i--{
        for j:=i;j>=0;j--{
            p:=i*j
            if p<lp{
                break
            }
            if isp(p) && p>lp{
                lp=p
                m1=i
                m2=j
            }
        }
    }
    fmt.Println("largest palindrome product:",lp)
    fmt.Println("multiplicand 1:",m1)
    fmt.Println("multiplicand 2:",m2)
}
