package main
import(
	"fmt"
	"net/http"
)
func main(){
	var port string
	fmt.Print("Enter the port number")
	fmt.Scan(&port)
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintln(w,"hello world")
	})
	fmt.Printf("Server is running on localhost:%s\n",port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Server error: %s\n", err)
	}
}

