package api

import (
	"fmt"

	"github.com/coderajay94/microservice1/model"
	cmap "github.com/orcaman/concurrent-map"
)

type MemoryUserAccounts struct {
	AllUserAccounts cmap.ConcurrentMap
}

//initalize the memory user accounts
func InitMemoryUserAccounts() MemoryUserAccounts{
	m := MemoryUserAccounts{
		AllUserAccounts: cmap.New(),
	}
	m.StoreUserAccounts()
	return m 
}


//store
func (m MemoryUserAccounts) StoreUserAccounts(){
	m.AllUserAccounts.Set("ajaykumar@gmail.com", model.UserResponseDB{
		Email : "ajaykumar@gmail.com",
		Name : "ajaykumar",
		Balance: 123455.33,
		AccountNumber: "ABS1233",

	})

	m.AllUserAccounts.Set("raghu@gmail.com", model.UserResponseDB{
		Email : "raghu@gmail.com",
		Name : "raghu",
		Balance: 3455.33,
		AccountNumber: "ASDF1234",

	})
	fmt.Println(m.AllUserAccounts)
}

func (m MemoryUserAccounts) GetAccountDetails(req model.UserRequestDB) (model.UserResponseDB, error) {
	  user, found := m.AllUserAccounts.Get(req.Email)
	count :=  m.AllUserAccounts.Count()
	  fmt.Println("if user found:", found, count)
	if found{
		userRes, _ := user.(model.UserResponseDB)
		return userRes, nil
	}
	return model.UserResponseDB{}, nil
}
