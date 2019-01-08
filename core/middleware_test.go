package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	helper "github.com/zokypesch/ceria/helper"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	rt "github.com/zokypesch/ceria/route"
)

func TestMiddleWare(t *testing.T) {
	middle := NewServiceMiddleWare()
	midSet := middle.GetMiddleWareAPI()
	// test gin jwt
	t.Run("Test all function is already in gin-jwt", func(t *testing.T) {

		assert.NotEmpty(t, midSet.Authenticator)
		assert.NotEmpty(t, midSet.Unauthorized)
		assert.NotEmpty(t, midSet.Authorizator)
	})

	// test a gin working
	t.Run("Test is middleware use in real case", func(t *testing.T) {
		r := rt.NewRouteService(true, "../templates/*", true)
		newHelper := helper.NewServiceHTTPHelper()

		rc, _ := r.Register(false)

		// add new route path
		rc.POST("/login", midSet.LoginHandler)

		rc.GET("/refresh_token", midSet.RefreshHandler)

		auth := rc.Group("/auth")
		auth.Use(midSet.MiddlewareFunc())
		{
			auth.GET("/loginSuccess", showIndexPageAPI)

		}

		var response map[string]string

		// test the wrong pass
		jsonFailParams := map[string]string{"username": "udin", "password": "123456"}
		jsonValueFail, _ := json.Marshal(jsonFailParams)

		w := newHelper.TestAPI(rc, "POST", "/login", jsonValueFail, nil)
		json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Equal(t, http.StatusUnauthorized, w.Code)

		// test with empty pasword
		jsonFailParams = map[string]string{"username": "admin"}
		jsonValueFail, _ = json.Marshal(jsonFailParams)

		w = newHelper.TestAPI(rc, "POST", "/login", jsonValueFail, nil)
		json.Unmarshal([]byte(w.Body.String()), &response)

		assert.Equal(t, http.StatusUnauthorized, w.Code)

		// success
		jsonParams := map[string]string{"username": "admin", "password": "admin"}
		jsonValue, _ := json.Marshal(jsonParams)

		var responseSuccess map[string]string

		s := newHelper.TestAPI(rc, "POST", "/login", jsonValue, nil)
		json.Unmarshal([]byte(s.Body.String()), &responseSuccess)

		// set a token
		token := responseSuccess["token"]

		assert.Equal(t, http.StatusOK, s.Code)

		aS := newHelper.TestAPI(
			rc,
			"GET",
			"/refresh_token",
			nil,
			map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)},
		)

		json.Unmarshal([]byte(aS.Body.String()), &response)
		assert.Equal(t, http.StatusOK, aS.Code)

		aF := newHelper.TestAPI(
			rc,
			"GET",
			"/refresh_token",
			nil,
			map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token)},
		)

		json.Unmarshal([]byte(aF.Body.String()), &response)
		assert.Equal(t, http.StatusOK, aF.Code)

		// fmt.Println(response["token"] == token, response["token"])
		aH := newHelper.TestAPI(
			rc,
			"GET",
			"/auth/loginSuccess",
			nil,
			map[string]string{"Authorization": fmt.Sprintf("Bearer %s", response["token"])},
		)
		json.Unmarshal([]byte(aH.Body.String()), &response)
		assert.Equal(t, http.StatusOK, aH.Code)

	})

}

func TestSession(t *testing.T) {
	middle := NewServiceMiddleWare()

	r := rt.NewRouteService(true, "../templates/*", true)
	newHelper := helper.NewServiceHTTPHelper()
	var handler gin.HandlerFunc
	rc, _ := r.Register(true)

	t.Run("Test Get Session and register it and return nil", func(t *testing.T) {
		handler = middle.InitialSessionMiddleWare("")
		assert.Empty(t, handler)
	})

	t.Run("Test Get Session and register it and return handler func", func(t *testing.T) {
		handler = middle.InitialSessionMiddleWare("myUsername")
		assert.NotEmpty(t, handler)

		rc.Use(handler)
		myUser := "triadi"

		auth := rc.Group("auth")

		auth.Use(middle.GetFullSessionMiddleWare)
		{
			auth.GET("/getsession", func(c *gin.Context) {
				session := sessions.Default(c)
				var userFromSession string

				userFromSession = session.Get("username").(string)

				c.HTML(
					http.StatusOK,
					"index.html",
					gin.H{
						"title": "hello world ! " + userFromSession,
					},
				)
			})
		}

		rc.GET("/setsession", func(c *gin.Context) {
			session := sessions.Default(c)

			session.Set("username", myUser)
			session.Save()

			c.HTML(
				http.StatusOK,
				"index.html",
				gin.H{
					"title": "success",
				},
			)
		})

		req, _ := http.NewRequest("GET", "/setsession", nil)

		err := newHelper.TestHTTPResponse(t, rc, req, func(w *httptest.ResponseRecorder) (bool, error) {
			statusOK := w.Code == http.StatusOK
			p, err := ioutil.ReadAll(w.Body)
			pageOK := err == nil && strings.Index(string(p), "<title>success</title>") > 0

			if statusOK && pageOK {
				return true, nil
			}
			return false, fmt.Errorf("Error when fetch a title")
		})

		assert.NoError(t, err)

		// check session cannot get session because it's gin setmode testmode
		reqGetSession, _ := http.NewRequest("GET", "/auth/getsession", nil)
		errGetSession := newHelper.TestHTTPResponse(t, rc, reqGetSession, func(w *httptest.ResponseRecorder) (bool, error) {
			statusOK := w.Code == http.StatusTemporaryRedirect // 307 status code temporary redirect
			if statusOK {
				return true, nil
			}
			return false, fmt.Errorf("Error when fetch a title")
		})

		assert.NoError(t, errGetSession)

	})
}

// example api test
func showIndexPageAPI(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		gin.H{
			"title": "hello World",
		},
	)
}
