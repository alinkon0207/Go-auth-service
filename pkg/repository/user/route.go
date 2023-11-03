package user

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, c controller) {
	// public routes (no auth required)
	public := r.Group("/api/auth")
	public.POST("/checkemail", c.CheckEmail)
	public.POST("/upload", c.Upload)
	public.POST("/verify", c.Verify)
	public.POST("/verifycode", c.VerifyCode)
	public.POST("/register", c.Register)
	public.POST("/login", c.Login)
	public.POST("/resetpassword", c.ResetPassowrd)

	// private routes (auth required)
	private := r.Group("/api/auth")
	private.Use(c.AuthMiddleware)
	private.GET("/user", c.Get)
	private.PUT("/user", c.Update)
	private.DELETE("/user", c.Delete)
}
