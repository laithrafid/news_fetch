package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/laithrafid/news_fetch/clients"
	"github.com/laithrafid/news_fetch/repository"
)

type GNewsService struct {
	Client   *clients.GClient
	DbClient *repository.DbSession
}

func GetNewGNews() *NewsService {
	return &NewsService{
		Service: &GNewsService{
			Client:   clients.InitGNewsClient(),
			DbClient: repository.GetNewDbSession(),
		},
	}
}

func (g *GNewsService) GetNews(c *gin.Context) {
	// fetch from db and return response
	srcs := ""
	for _, val := range qp["sources"] {
		sVal := strings.Split(val, ",")
		for _, v := range sVal {
			srcs = srcs + "'" + v + "',"
		}
	}
	srcs = strings.TrimRight(srcs, ",")
	if dbResp, err := g.DbClient.GetTopNewsBySource(srcs, lmt); err == nil {
		c.JSON(http.StatusOK, dbResp)
	}
}
