package main

import (
	"fmt"
	"log"

	"github.com/kavirajk/go-patterns/store"
	_ "github.com/lib/pq"
)

func main() {
	st := store.NewStore()
	user, err := st.User().GetByUsername("user1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user, user.Friends)
}
