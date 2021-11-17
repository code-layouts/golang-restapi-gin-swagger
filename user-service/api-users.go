package apiuser

import (
	f "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	Id        int        `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Title     string     `json:"title"`
	Email     string     `json:"email"`
	Role      string     `json:"role"`
	Usercode  string     `json:"usercode"`
	CreateDts *time.Time `json:"createDts"`
	UpdateDts *time.Time `json:"updateDts"`
}

func GetUsers(c *gin.Context) {
	f.Println("GetUsers")
	var users = toUsers()
	var repo = NewRepository(users)
	c.JSON(http.StatusOK, repo.GetAll())
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	idVal, _ := strconv.Atoi(id)
	f.Println("GetUser id:", idVal)

	var users = toUsers()
	var repo = NewRepository(users)
	var user = repo.GetById(idVal)
	if user.Id > -1 {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "no value"})
	}
}

func AddUser(c *gin.Context) {
	f.Println("AddUser()")
	var user User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	f.Println("user: ", user)
	var users = toUsers()
	var repo = NewRepository(users)

	f.Println("users.len: ", len(repo.users))

	repo.Add(user)

	c.JSON(http.StatusOK, gin.H{"message": "no value"})
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	idVal, _ := strconv.Atoi(id)
	f.Println("DeleteUser By Id:", idVal)
	var repo = NewRepository(toUsers())
	if repo.Exists(idVal) == false {
		c.JSON(http.StatusNotFound, gin.H{"message": "resource not found"})
		return
	}
	repo.Delete(idVal)
	c.JSON(http.StatusNoContent, gin.H{"message": "Resource deleted successful."})
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	idVal, _ := strconv.Atoi(id)
	f.Println("UpdateUser()", id)
	if idVal < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Can not find User ID"})
		return
	}
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Json UnMarshalling Error"})
		return
	}
	if idVal != user.Id {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User ID dose not match."})
		return
	}
	var repo = NewRepository(toUsers())
	oldUser := repo.GetById(idVal)
	if oldUser == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Resource not found. {idVal}"})
		return
	}

	updatedUser, err := repo.Update(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, updatedUser)

}
