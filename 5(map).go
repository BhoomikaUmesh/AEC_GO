package main
import "fmt"
func main(){
    mymap:=make(map[string]int)
    for{
        var ch int
        fmt.Println("Enter your choice:\n1.Create\n2.Delete\n3.Update\n4.Display\n5.Eixt\n")
        fmt.Scan(&ch)
        switch ch{
            case 1:
            var key string
            var val int
            fmt.Println("Enter key and val:")
            fmt.Scan(&key,&val)
            mymap[key]=val
            
            case 2:
            var delkey string
            fmt.Println("Enter key to delete:")
            fmt.Scan(&delkey)
            delete(mymap,delkey)
            
            case 3:
            var upkey string
            fmt.Println("Enter key to up:")
            fmt.Scan(&upkey)
            if _,exists:=mymap[upkey];exists{
                var newval int
                fmt.Println("Enter val:")
            fmt.Scan(&newval)
                mymap[upkey]=newval
            }else{
                fmt.Println("not possible")
            }
			
            case 4:
            for key, val := range mymap {
                fmt.Printf("Key: %s, Value: %d\n", key, val)
            }

            default:fmt.Println("exiting")
            return
            
        }
    }
    
}