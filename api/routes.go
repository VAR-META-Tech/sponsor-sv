package api

import (
	// "sponsor-sv/services/account"
	"net/http"
	"sponsor-sv/services/account"
	"sponsor-sv/services/sponsor"
	"sponsor-sv/services/transfer"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Routers struct {
	Pattern  string
	Method   string
	HandlerF func(*gin.Context)
	Name     string
}

var routerList = []Routers{
	// not use yet
	{
		Pattern:  "/account/details",
		Method:   "GET",
		HandlerF: account.GetAccountHandler,
		Name:     "get account detail in std.BaseAccount",
	},

	// not use yet
	// {
	// 	Pattern:  "/sponsor/list",
	// 	Method:   "GET",
	// 	HandlerF: sponsor.ListAllHandler,
	// 	Name:     "get list sponsor",
	// },

	{
		Pattern:  "/sponsor/transfer",
		Method:   "POST",
		HandlerF: transfer.TransferHandler,
		Name:     "transfer the information to backend",
	},

	{
		Pattern:  "/user/balance/:id",
		Method:   "GET",
		HandlerF: sponsor.GetBalance,
		Name:     "get endpoint",
	},
}

func NewHandler() http.Handler {
	baseRoute := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowMethods("GET", "POST", "PUT", "DELETE")
	corsMW := cors.New(config)
	v1Feature := baseRoute.Group("/v1")
	gin.SetMode(gin.DebugMode)
	v1Feature.Use(corsMW)

	for _, router := range routerList {
		switch router.Method {
		case http.MethodGet:
			{
				v1Feature.GET(router.Pattern, router.HandlerF)
			}
		case http.MethodPost:
			{
				v1Feature.POST(router.Pattern, router.HandlerF)
			}
		case http.MethodPut:
			{
				v1Feature.PUT(router.Pattern, router.HandlerF)
			}
		case http.MethodDelete:
			{
				v1Feature.DELETE(router.Pattern, router.HandlerF)
			}
		}
	}
	return baseRoute
}
