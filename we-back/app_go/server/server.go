package server

import (
	"app_go/controller"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Init サーバを起動する
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	// CORS対応
	r.Use(cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		MaxAge: 24 * time.Hour,
	}))

	// タスクルータ
	task := r.Group("/task")
	{
		// タスクコントローラー
		tc := controller.TaskController{}
		task.GET("", tc.GetTaskList)
		task.POST("", tc.RegistTask)
		task.PUT("", tc.UpdateTask)
		task.DELETE("", tc.DeleteTask)
	}

	// 計画ルータ
	plan := r.Group("/plan")
	{
		// 計画コントローラー
		pc := controller.PlanController{}
		plan.GET("", pc.GetPlanList)
		plan.GET("/summary", pc.GetPlanSummary)
		plan.POST("/regist", pc.UpdateItemRegist)
	}

	// ユーザールータ
	user := r.Group("/user")
	{
		// ユーザーコントローラー
		uc := controller.UserController{}
		user.POST("/auth", uc.CertifyUser)
		user.POST("/regist", uc.RegistUser)
	}

	return r
}
