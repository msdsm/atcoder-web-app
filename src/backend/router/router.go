package router

import (
	"net/http"
	"os"

	"atcoder-web-app/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, rc controller.IRivalController) *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true, // cookieの送受信を可能にする
	}))

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		// CookieSameSite: http.SameSiteNoneMode,
		CookieSameSite: http.SameSiteDefaultMode, // secure modeをfalseにしないとpostmanで動作確認できないから
		CookieMaxAge:   60,
	}))

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

	t := e.Group("/user")
	// middleware
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("/rival", rc.GetAllRivals)       // ライバルユーザー全聚徳
	t.POST("/rival", rc.CreateRival)       // ライバルユーザー追加
	t.DELETE("/rival/:id", rc.DeleteRival) // ライバルユーザー削除
	t.GET("/table", rc.GetTable)           // ライバルユーザー情報取得
	t.GET("/submission", rc.GetSubmission) // 今日の提出取得
	t.GET("/profile", uc.GetAtcoderId)     // 自分のAtcoder ID取得
	t.POST("/profile/:id", uc.Update)      // 自分のAtcoder ID 変更
	return e
}
