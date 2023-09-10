package main
import "fmt"
func main(){
    var prev,current,sum,limit int
    prev=0
    current=1
    sum=0
    fmt.Printf("Enter limit")
    fmt.Scan(&limit)
    for prev+current<limit{
        next:=prev+current
        if next%2==0{
            fmt.Println(next)
            sum+=next
        }
        prev=current
        current=next
    }
    fmt.Println("sum:",sum)
}