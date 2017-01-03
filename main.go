package main
import (
"os"
service "github.com/dkassab/gogo-service/service"
"fmt"
//"reflect"
)
 
func main(){
    port := os.Getenv("PORT")
	fmt.Println(port)
	if len(port)==0 {
		port="8000"
	}

	server := service.NewServer()
	server.Run(":"+port)
}
