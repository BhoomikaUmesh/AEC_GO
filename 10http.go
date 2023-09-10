package main
import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)
type Post struct{
	userId int `json:"userId"`
	Id int `json:"Id"`
	Title string `json:"Title"`
	Body string `json:"Body"`
}
func main(){
	var postId int
	fmt.Printf("enter the post id")
	fmt.Scan(&postId)
	url:=fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d",postId)
	response,err:=http.Get(url)
	if err !=nil{
		fmt.Printf("lala")
		return
	}
	defer response.Body.Close()
	body, err:= ioutil.ReadAll(response.Body)
	if err !=nil{
		fmt.Printf("lala")
		return
	}
	var post Post
	err= json.Unmarshal(body,&post)
	if err!=nil{
		fmt.Printf("lala")
		return
	}
	fmt.Printf("user id:%d\n",post.userId)
	fmt.Printf("Id:%d\n",post.Id)
	fmt.Printf("Title:%s\n",post.Title)
	fmt.Printf("Body:%s\n",post.Body)

}
 