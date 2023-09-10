package main
import "fmt"
func spt(n int)(int,int,int){
    for a:=1;a<=n/3;a++{
        for b:=a;b<=((n-a)/2);b++{
            c:=n-a-b
            if a*a+b*b==c*c{
                return a,b,c
            }
        }
    }
    return -1,-1,-1
}
func main(){
    var n int
    fmt.Println("enter n: ")
    fmt.Scan(&n)
    a,b,c:=spt(n)
    if a!=-1 && b!=-1 && c!=-1{
        fmt.Printf("%d,%d,%d\n", a,b,c)
    }else{
        fmt.Println("not present")
    }
}
