package database

import (
	"fmt"
	"github.com/mariumfaheem/PE-GO-PROJECT/models"
	"sync"
)

var (
	db []*models.User
	mu sync.Mutex
)

func Connect() {
	db = make([]*models.User, 0)
	fmt.Println("Connected with database")
}

func Insert(user *models.User) {
	mu.Lock()
	db = append(db, user)
	mu.Unlock()

}

func Get() []*models.User {
	return db
}
