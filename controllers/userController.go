package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/ankit-k56/Repelit/initializers"
	"github.com/ankit-k56/Repelit/models"
	"github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) { 

	//Retieving email and password from the request body
	var body struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	// Hashing the password
	hashedPassword , err := bcrypt.GenerateFromPassword([]byte(body.Password),10)

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Internal server error"})
		return
	}

	//Creating a new user
	user := models.User{Email : body.Email, Password: string(hashedPassword)}

	result := initializers.Db.Create(&user)

	if result.Error != nil{
		c.JSON(http.StatusInternalServerError,gin.H{"message":"Internal server error"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"email": user.Email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Internal server error"})
		return
	}

	c.SetCookie("Authorization", tokenString, 60*60*24*30,"", "", false, true)
	c.JSON(http.StatusCreated, gin.H{"message":"User created successfully"})

}

func Login(c *gin.Context){
	var body struct{
		Email string `json:"email"`
		Password string `json:"password"`

	}
	if c.Bind(&body)!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid Request"})
		return
	}
	//Serach User
	var user models.User
	initializers.Db.First(&user, "email = ?", body.Email)

	if user.ID == 0{
		c.JSON(http.StatusBadRequest,gin.H{"message":"User not found"})
		return
	}


	// Compare Password
	err :=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(body.Password))

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Password"})
		return
	}

	//Generate and send Jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"email": user.Email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Internal server error"})
		return
	}

	c.SetCookie("Authorization", tokenString, 60*60*24*30,"", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Login Succesfull"})

}

func Validated(c *gin.Context){
	var user1 models.User

	user,_:= c.Get("user")

	user1, ok :=user.(models.User)
	if !ok{
		c.JSON(http.StatusBadRequest, gin.H{"message":"User not found"})
		return
	
	}


	c.JSON(http.StatusOK, gin.H{"message":"Validated", "user": user1.Email})


}
