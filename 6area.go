package main
import(
    "fmt"
    "math"
)
type Shape interface{
    Area()float64
}
type Circle struct{
    radius float64
}
func(c Circle)Area()float64{
    return math.Pi*c.radius*c.radius
}
type Rectangle struct{
    height float64
    width float64
}
func(r Rectangle)Area()float64{
    return r.height*r.width
}
func print(s Shape){
    fmt.Println("area:",s.Area())
}
func main(){
    var r,h,w float64
    fmt.Println("enter radius of circle")
    fmt.Scan(&r)
    circle:=Circle{radius:r}
    print(circle)
    fmt.Println("enter h and w")
    fmt.Scan(&h,&w)
    rectangle:=Rectangle{height:h,width:w}
    print(rectangle)
}
