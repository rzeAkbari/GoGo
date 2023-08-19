package cmd

import (
	"github.com/rzeAkbari/GoGo/gin/server"
	"github.com/rzeAkbari/GoGo/gin/store"
	"log"
	"os"
)

func Run() error {
	dir, _ := os.Getwd()
	store, err := store.NewStore(dir + "/data/gin.db")
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
		return err
	}
	defer store.DB.Close()

	server := server.NewServer(store)
	os.Setenv("JWT_SIGN_KEY", "AllYourBase")

	s := server.SetupRouter()
	err = s.Run(":5000")
	if err != nil {
		log.Fatalf("error starting the server: %v", err)
		return err
	}
	return nil
}
