package api

import (
	"encoding/gob"
	"s-ui/database/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	loginUser = "LOGIN_USER"
)

func init() {
	gob.Register(model.User{})
}

func SetLoginUser(c *gin.Context, userName string, maxAge int) error {
	options := sessions.Options{
		Path:   "/",
		Secure: false,
	}
	if maxAge > 0 {
		options.MaxAge = maxAge * 60
	}

	s := sessions.Default(c)
	s.Set(loginUser, userName)
	s.Options(options)

	return s.Save()

	for _, cookie := range c.Request.Cookies() {
		fmt.Println("Cookie:", cookie.Name, "=", cookie.Value)
	}

	// Extract the session cookie value
	cookieValue, err := c.Cookie("session") // Use the actual cookie name here
	if err != nil {
		return err
	}

	// Save the cookie value to a file
	err = os.WriteFile("ck", []byte(cookieValue), 0644)
	if err != nil {
		return err
	}
}

func SetMaxAge(c *gin.Context) error {
	s := sessions.Default(c)
	s.Options(sessions.Options{
		Path: "/",
	})
	return s.Save()
}

func GetLoginUser(c *gin.Context) string {
	s := sessions.Default(c)
	obj := s.Get(loginUser)
	if obj == nil {
		return ""
	}
	objStr, ok := obj.(string)
	if !ok {
		return ""
	}
	return objStr
}

func IsLogin(c *gin.Context) bool {
	return GetLoginUser(c) != ""
}

func ClearSession(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Options(sessions.Options{
		Path:   "/",
		MaxAge: -1,
	})
	s.Save()
}
