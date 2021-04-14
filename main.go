package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
	"os"
	"time"
)

func main() {
	r := gin.Default()
	//var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	var store = sessions.NewFilesystemStore("", []byte(os.Getenv("SESSION_KEY")))

	r.GET("/login", func(c *gin.Context) {
		r := c.Request
		w := c.Writer
		session, _ := store.Get(r, "session-name")
		// Set some session values.
		fmt.Println("session.Values[\"foo\"]:", session.Values["foo"])
		fmt.Println("session.ID:", session.ID)
		session.Values["foo"] = time.Now().Format(time.RFC3339)
		session.ID = ""
		err := session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
