package businessmiddlewares

import (
	businessapi "apps/api/business"
	businesshelpers "apps/business/helpers"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(f businessapi.StrictHandlerFunc, operationID string) businessapi.StrictHandlerFunc {
    return func(ctx echo.Context, i interface{}) (interface{}, error) {
		if !needsAuthenticate(ctx) {
			// NOTE: 認証が不要なURIは認証をスキップ
			return f(ctx, i)
		}

        // NOTE: Cookieからtokenを取得し、JWTの復号
		tokenString, _ := ctx.Cookie("token")
		if tokenString == nil {
			return nil, echo.ErrUnauthorized
		}

		token, _ := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_TOKEN_KEY")), nil
		})

		// NOTE: ログイン種別に応じ、IDをContextにセットする
		c, err := newWithAuthenticateContext(token, ctx)
		if err != nil {
			return nil, echo.ErrUnauthorized
		}

		ctx.SetRequest(ctx.Request().WithContext(c))
        return f(ctx, i)
    }
}

func needsAuthenticate(ctx echo.Context) (bool) {
	spec, _ := businessapi.GetSwagger()
	security := spec.Paths.Value(ctx.Request().RequestURI).Operations()[ctx.Request().Method].Security
	
	return len(*security) > 0
}

func newWithAuthenticateContext(token *jwt.Token, ctx echo.Context) (context.Context, error) {
	var authenticateID int
	var authenticateType string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["supporter_id"] != nil {
			authenticateID = int(claims["supporter_id"].(float64))
			authenticateType = "supporter"
		} else if claims["company_id"] != nil {
			authenticateID = int(claims["company_id"].(float64))
			authenticateType = "company"
		} else {
			return nil, errors.New("invalid signed type")
		}
	}

	var c context.Context
	switch authenticateType {
	case "supporter":
		c = businesshelpers.NewWithSupporterIDContext(ctx.Request().Context(), authenticateID)
	case "company":
		c = businesshelpers.NewWithCompanyIDContext(ctx.Request().Context(), authenticateID)
	}
	return c, nil
}
