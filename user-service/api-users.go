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
		c.JSON(http.StatusOK, gin.H{"status": "no value"})
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

	c.JSON(http.StatusOK, gin.H{"status": "no value"})
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	idVal, _ := strconv.Atoi(id)
	f.Println("DeleteUser By Id:", idVal)
	var repo = NewRepository(toUsers())
	repo.Delete(idVal)
	c.JSON(http.StatusOK, gin.H{"status": "no value"})
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	idVal, _ := strconv.Atoi(id)
	f.Println("UpdateUser By Id:", idVal)
	var user User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	if idVal != user.Id {
		f.Println("Does not match User ID")
		return
	}
	var repo = NewRepository(toUsers())
	updatedUser, err := repo.Update(user)
	if err != nil {
		f.Println("Updated ERR", err)
		return
	}
	c.JSON(http.StatusOK, updatedUser)

}
