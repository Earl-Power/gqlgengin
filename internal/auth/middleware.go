package auth

import (
	"context"
	"github.com/Earl-Power/gqlgengin/internal/pkg/jwt"
	"github.com/Earl-Power/gqlgengin/internal/users"
	"net/http"
	"strconv"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			header := request.Header.Get("Authorization")
			// 允许未经身份验证的用户进入
			if header == "" {
				next.ServeHTTP(writer, request)
				return
			}

			// 验证jwt令牌
			tokenStr := header
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(writer, "Invalid token", http.StatusForbidden)
				return
			}

			// 创建用户并检查数据库中是否存在该用户
			user := users.User{Username: username}
			id, err := users.GetUserIdByUsername(username)
			if err != nil {
				next.ServeHTTP(writer, request)
				return
			}
			user.ID = strconv.Itoa(id)
			ctx := context.WithValue(request.Context(), userCtxKey, &user)

			// 并使用我们的新上下文调用下一个
			request = request.WithContext(ctx)
			next.ServeHTTP(writer, request)
		})
	}
}

// ForContext 从上下文中查找用户。 需要运行中间件。
func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)
	return raw
}
