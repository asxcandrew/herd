package middlewares

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/asxcandrew/herd/server/models"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

type basicAuthorizer struct {
	enforcer *casbin.Enforcer
	context  *gin.Context
}

var _authorizer *basicAuthorizer

func Authorizer() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _authorizer == nil {
			e := casbin.NewEnforcer("access_model.conf", "access_policy.csv")
			_authorizer = &basicAuthorizer{enforcer: e, context: c}
		}
		if !_authorizer.CheckPermission(c.Request) {
			_authorizer.RequirePermission(c)
		}
	}
}

// CheckPermission checks the role/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *basicAuthorizer) CheckPermission(r *http.Request) bool {
	role := a.getRole()
	method := r.Method
	path := r.URL.Path
	return a.enforcer.Enforce(role, path, method)
}

func (a *basicAuthorizer) getRole() string {
	claims := jwt.ExtractClaims(a.context)
	if claims["userID"] == nil {
		return models.GuestRole
	}
	return claims["userRole"].(string)
}

// RequirePermission returns the 403 Forbidden to the client
func (a *basicAuthorizer) RequirePermission(c *gin.Context) {
	c.AbortWithStatus(http.StatusForbidden)
}
