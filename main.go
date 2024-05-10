package main

import (
	"fmt"

	"github.com/ankit-k56/Repelit/initializers"
	"github.com/ankit-k56/Repelit/router"
)

func init(){
	initializers.LoadEnv()
	initializers.ConnecttoDb()
	// initializers.SyncDatabase()
}
func main(){


	


	r := router.NewRouter();
	fmt.Println("Repelit backend")
	r.Run(":8080") 
}