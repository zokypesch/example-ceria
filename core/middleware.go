package core

import (
	"reflect"
	"time"

	"github.com/appleboy/gin-jwt"
	// "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// MiddleWareHandler for handler of middleware
type MiddleWareHandler struct{}

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// User struct get info for middleware
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

var identityKey = "id"

var middle *MiddleWareHandler

// NewServiceMiddleWare for get middleware
func NewServiceMiddleWare() *MiddleWareHandler {
	if middle == nil {
		middle = &MiddleWareHandler{}
	}
	return middle
}

// GetMiddleWareAPI for get middleware api using gin JWT
func (mid *MiddleWareHandler) GetMiddleWareAPI() *jwt.GinJWTMiddleware {
	jwt := &jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
				return &User{
					UserName:  userID,
					LastName:  "Bo-Yi",
					FirstName: "Wu",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {

			var usrnm string

			if reflect.TypeOf(data).Kind() == reflect.String {
				usrnm = reflect.ValueOf(data).String()
			}

			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}

			if usrnm == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}

	return jwt

}

// InitialSessionMiddleWare for initial session web using redis
func (mid *MiddleWareHandler) InitialSessionMiddleWare(name string) gin.HandlerFunc {
	store, err := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))

	if name == "" || err != nil {
		return nil
	}

	return sessions.Sessions(name, store)
}

// GetFullSessionMiddleWare for get session full of middleware
func (mid *MiddleWareHandler) GetFullSessionMiddleWare(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("username")

	if v == nil {
		session.Set("username", "")
		session.Save()
	}

	if session.Get("username") == "" {
		c.Redirect(307, "/")
	}

	c.Next()
}
