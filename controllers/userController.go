package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	database "github.com/1shubham7/jwt/database"
	helper "github.com/1shubham7/jwt/helpers"
	models "github.com/1shubham7/jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func Hashpassword()

func VerifyPassword()

func SignUp() {
	
}

func Login()

func GetUsers()

func GetUserById() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id") // we are taking the user_id given by the user in json
		// with the help of gin.context we can access the json data send by postman or curl or user
		
		if err := helper.MatchUserTypeToUserId(c, userId); err != nil{

			//checking if the user in admin or not.
			// we will create that func in helper package.

				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return 
			}

	}

		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId}).Decode(&user)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// if everything goes ok, pass the data of the user (UserModel.go)
		c.JSON(http.StatusOK, user)

	}
}