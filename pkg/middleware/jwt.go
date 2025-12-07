package middleware

import (
	"context"
	"log"
	"strings"
	myutils"todo-list/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func JWTAuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		authHeader := string(c.GetHeader("Authorization"))
		if authHeader == "" {
			log.Println("Authorization header is missing")
			c.JSON(consts.StatusUnauthorized, utils.H{
				"error": "Authorization header is missing",
			})
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == authHeader {
			log.Println("Authorization header format is wrong")
			c.JSON(consts.StatusUnauthorized, utils.H{
				"error": "Authorization header format must be Bearer {token}",
			})
			c.Abort()
			return
		}

		claims, err := myutils.ParseJWT(tokenStr)
		if err != nil {
			log.Println(err)
			c.JSON(consts.StatusUnauthorized, utils.H{
				"error": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		c.Set("userId", claims.UserId)

		c.Next(ctx)
	}
}