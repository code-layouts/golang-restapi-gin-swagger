package apiuser

import (
	"encoding/json"
	"errors"
	f "fmt"
	"io/ioutil"
)

const JsonFileUser = "./data/users.json"

func toUsers() []User {
	file, err := ioutil.ReadFile(JsonFileUser)
	if err != nil {
		f.Print("err:", err)
	}

	var users []User
	if err := json.Unmarshal(file, &users); err != nil {
		f.Println("toUsers()--err:", err)
	}
	return users
}

type UserRepository interface {
	GetAll() []User
	GetById(id int) (User, error)
	NewId(users []User) int
	Add(user User) (int, error)
	Delete(id int) error
	Update(user User) (User, error)
	Exists(id int) bool
}

type JsonUserRepository struct {
	// nextId int
	users []User
	// s      sync.RWMutex
}

func NewRepository(users []User) *JsonUserRepository {
	return &JsonUserRepository{users}
}

func (repo *JsonUserRepository) Exists(id int) bool {
	for _, e := range repo.users {
		if e.Id == id {
			return true
		}
	}
	return false
}

func (repo *JsonUserRepository) NewId() int {
	var result int
	for _, e := range repo.users {
		if e.Id > result {
			result = e.Id
		}
	}
	return result + 1
}

func (repo *JsonUserRepository) GetAll() []User {
	return repo.users
}

func (repo *JsonUserRepository) GetById(id int) *User {
	for _, e := range repo.users {
		if e.Id == id {
			return &e
		}
	}
	return nil
}

func (repo *JsonUserRepository) Add(user User) int {
	var newId = repo.NewId()
	user.Id = newId
	var list = repo.users
	list = append(list, user)
	result, err := json.MarshalIndent(list, "", "  ")
	if err != nil {
		f.Println("AddToJson.err: ", err)
	}
	if fileErr := ioutil.WriteFile(JsonFileUser, result, 0644); fileErr != nil {
		f.Println("fileErr:", fileErr)
	}
	return newId
}

func (repo *JsonUserRepository) Delete(id int) error {
	var idx = -1
	for i, e := range repo.users {
		if e.Id == id {
			idx = i
			break
		}
	}
	if idx < 1 {
		return errors.New(f.Sprintf("Can not find UserID '%s'", id))
	}
	users := repo.users

	copy(users[idx:], users[idx+1:])
	users = users[:len(users)-1]
	result, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return errors.New(f.Sprintln("Json Marshalling Error ", err.Error()))
	}
	if fileErr := ioutil.WriteFile(JsonFileUser, result, 0644); fileErr != nil {
		return errors.New(f.Sprintln("File Write Error ", err.Error()))
	}
	return nil
}

func (repo *JsonUserRepository) Update(user User) (User, error) {
	var users = repo.users
	var idx int = -1
	for i, e := range users {
		if e.Id == user.Id {
			idx = i
			break
		}
	}
	users[idx] = user
	result, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		f.Println("Update.err: ", err)
	}
	if fileErr := ioutil.WriteFile(JsonFileUser, result, 0644); fileErr != nil {
		f.Println("fileErr:", fileErr)
	}
	return user, nil
}
