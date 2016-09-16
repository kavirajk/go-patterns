package main

import (
	"fmt"
	"log"

	"github.com/kavirajk/go-patterns/models"
	"github.com/kavirajk/go-patterns/store"
	_ "github.com/lib/pq"
)

func main() {
	st := store.NewStore()
	// user, err := st.User().GetByUsername("user1")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(user, user.Friends)

	// Album store
	// album := models.Album{Title: "Second Album", OwnerId: 7, IsActive: true}
	// if err := st.Album().Save(&album); err != nil {
	// 	log.Fatal(err)
	// }
	// user := models.User{Username: "kaviraj", Password: "kaviraj."}
	// if err := st.User().Save(&user); err != nil {
	// 	log.Fatal(err)
	// }
	album := models.Album{Title: "No idea", OwnerId: 2}
	if err := st.Album().Save(&album); err != nil {
		log.Fatal(err)
	}
	pic := models.Picture{Caption: "Andrew and Sarah", AlbumId: album.Id}
	if err := st.Picture().Save(&pic); err != nil {
		log.Fatal(err)
	}
	fmt.Println(album.Pictures)
}
