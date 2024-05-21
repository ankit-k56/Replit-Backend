package router

import (
	// "github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	"fmt"
	"log"
	"time"

	"github.com/ankit-k56/Repelit/controllers"
	"github.com/ankit-k56/Repelit/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}


func Webs(c *gin.Context){
	fmt.Println("Dope2")

	conn , err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil{

		// c.JSON(500, gin.H{"message": "Error in websocket connection"})
		log.Fatal(err)
		fmt.Println("Jo")
		return
	
	}
	fmt.Println("Dope")
	defer conn.Close()
	i :=0;
	for{
		i++;
		conn.WriteMessage(6, []byte("Hello world "+ string(i) ))
		time.Sleep(10 * time.Second)
	}

	
	// c.JSON(200, gin.H{"message": "Folder copied successfully"})
}
func NewRouter () *gin.Engine{
	r := gin.Default()
	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{"message": "Hello world"})
	})
	r.GET("/ws", Webs)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/check",middleware.CheckAuth, controllers.Validated)
	r.POST("/createProject", controllers.CreateProject)
	r.POST("/init", controllers.InialiseProject)
	r.GET("/download", controllers.DownloadProject)
	return r
}