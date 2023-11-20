package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/devsamahd/go-jwt-api/database"
	helper "github.com/devsamahd/go-jwt-api/helpers"
	"github.com/devsamahd/go-jwt-api/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func Signup() gin.HandlerFunc{
	return func (c *gin.Context){
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User

		if err := c.BindJSON(&user); errr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":validationErr.Error()})
			return 
		}
		count, err := userCollection.CountDocuments(ctx, bson.M{"email":user.Email})
		defer cancel()
	}
}

func Login(){

}

func GetUsers() {}

func GetSingleUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userId:= c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		var ctx, cacel = context.WithTimeout((context.Background()) 100*time.Second)
		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id":userId}).Decode(&user)
		defer cancel()
		if err != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			err
		}
		c.JSON(http.StatusOk, user)
	}
}

func HashPassword(){

}