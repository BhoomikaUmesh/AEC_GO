package main
import(
    "fmt"
    "math"
)
func lpf(n int)int{
    l:=1
    for n%2==0{
        l=2
        n/=2
    }
    for i:=3;i<int(math.Sqrt(float64(n)));i+=2{
        for n%i==0{
         l=i
         n/=i
        }
    }
    if n>2{
        l=n
    }
    return l
}
func main(){
    var num int
    fmt.Println("enter n: ")
    fmt.Scan(&num)
    ans:=lpf(num)
    fmt.Println(ans)
}