package controllers

import (
	"net/http"

	"github.com/ankit-k56/Repelit/models"
	awss3 "github.com/ankit-k56/Repelit/utils"

	"github.com/ankit-k56/Repelit/initializers"
	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context){
	var body models.Project
	
	c.BindJSON(&body)
	project := models.Project{Name: body.Name,UserID: body.UserID,Tech: body.Tech }
	result :=initializers.Db.Create(&project)
	if result.Error != nil{
		c.JSON(400, gin.H{"message": "Error creating project"})
		return
	}
	_,err := awss3.Uploader("repelit-iam","directory","./test.html")
	if err != nil{
		c.JSON(400, gin.H{"message": "Error uploading file"})
		return
	}

	// _, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
	// 	Bucket: aws.String("my-bucket"),
	// 	Key:    aws.String("my-object-key"),
	// 	Body:   "uploadFile",

	// })
	
	c.JSON(200, project)




	
}

func InialiseProject(c *gin.Context){
	var body struct{
		SourceString string `json:"srcString"`
		DestinationString string `json:"dstString"`
	}
	// fmt.Println(body.SourceString)
	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	err := awss3.CopyS3Folder(body.SourceString, body.DestinationString)

	if err != nil{
		c.JSON(400, gin.H{"message": "Error copying folder"})
		return
	}
	c.JSON(200, gin.H{"message": "Folder copied successfully"})

}