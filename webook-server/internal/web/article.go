package web

import (
	"log/slog"
	"net/http"
	"webook-server/internal/domain"
	"webook-server/internal/service"
	"webook-server/internal/web/middleware"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	svc service.ArticleService
}

func NewArticleHandler(svc service.ArticleService) ArticleHandler {
	return ArticleHandler{svc: svc}
}

func (h *ArticleHandler) InitRouter(r *gin.Engine, auth *middleware.AuthMiddleware) {
	base := r.Group("/api/article").Use(auth.JwtAuthMiddleware())
	base.POST("", h.CreateOrUpdate)
	base.DELETE("", h.Delete)
}

func (h *ArticleHandler) CreateOrUpdate(c *gin.Context) {
	type Req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	var req Req
	if err := c.ShouldBindJSON(&req); err != nil {
		slog.Error("请求参数有误", "Error", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	userIdRaw, exist := c.Get("user_id")
	authorId, ok := userIdRaw.(int64)
	if !exist || !ok {
		slog.Error("用户登录状态有误")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := h.svc.CreateOrUpdate(c, domain.Article{
		Title:    req.Title,
		Content:  req.Content,
		AuthorId: authorId,
	})
	if err != nil {
		slog.Error("CreateOrUpdate fail", "Error", err.Error())
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, Response{
		Code: 0,
	})
}

func (h *ArticleHandler) Delete(c *gin.Context) {

}
