package api

import (
	"sponsor-sv/services/account"
	"sponsor-sv/services/sponsor"
	"sponsor-sv/services/transfer"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Routers struct {
	Pattern  string
	Method   string
	HandlerF func(*gin.Context)
	Name     string
}

var routerList = []Routers{
	{
		Pattern:  "/transfer/sponsor",
		Method:   "POST",
		HandlerF: transfer.TransferHandler,
		Name:     "transfer the information to backend",
	},
	{
		Pattern:  "/account/details",
		Method:   "GET",
		HandlerF: account.GetAccountHandler,
		Name:     "get account detail in std.BaseAccount",
	},
	{
		Pattern:  "/sponsor/list",
		Method:   "GET",
		HandlerF: sponsor.ListAllHandler,
		Name:     "get list sponsor",
	},
	{
		Pattern:  "/sponsor/endpoint",
		Method:   "GET",
		HandlerF: sponsor.GetEndpoint,
		Name:     "get endpoint",
	},
}

func NewHandler() http.Handler {
	baseRoute := gin.Default()
	v1Feature := baseRoute.Group("/v1")
	gin.SetMode(gin.DebugMode)

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
