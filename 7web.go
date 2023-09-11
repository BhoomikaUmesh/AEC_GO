package main
import(
	"fmt"
	"net/http"
	"log"
)
func main(){
	var port string
	fmt.Print("Enter the port number")
	fmt.Scan(&port)
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintln(w,"hello world")
	})
	port=":3000"
	fmt.Printf("Server is running on localhost:%s\n",port)
	log.Fatal(http.ListenAndServe(port,nil))
}

