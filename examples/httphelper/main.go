package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"encoding/json"

	"github.com/zokypesch/ceria/helper"
	"github.com/zokypesch/ceria/route"
)

func main() {
	httpHelper := helper.NewServiceHTTPHelper()

	newRoute := route.NewRouteService(true, "../../templates/*", true)

	r, _ := newRoute.Register(true)

	r.GET("/hello", showIndexPageAPI)
	w := httpHelper.TestAPI(r, "GET", "/", nil, nil)
	var response map[string]interface{}

	err := json.Unmarshal([]byte(w.Body.String()), &response)

	if err != nil {
		panic(err)
	}

	fmt.Println(w.Code, response)

}

func showIndexPageAPI(context *gin.Context) {
	httpHelper := helper.NewServiceHTTPHelper()

	httpHelper.EchoResponse(context, 200, true, "success", "", nil)

}
func showIndexPageAPICreated(context *gin.Context) {
	httpHelper := helper.NewServiceHTTPHelper()

	httpHelper.EchoResponseCreated(context, nil)
}

func showIndexPageAPISuccess(context *gin.Context) {
	httpHelper := helper.NewServiceHTTPHelper()

	httpHelper.EchoResponseSuccess(context, nil)
}

func showIndexPageAPIFailed(context *gin.Context) {
	httpHelper := helper.NewServiceHTTPHelper()

	httpHelper.EchoResponseBadRequest(context, "failed get", "error description")
}

func showIndexPageAPIPagination(context *gin.Context) {
	httpHelper := helper.NewServiceHTTPHelper()

	httpHelper.EchoResponseWithPagination(context, nil, "1", "10")
}
