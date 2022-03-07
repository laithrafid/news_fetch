package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/laithrafid/news_fetch/src/cassandra"
	"github.com/laithrafid/news_fetch/src/handlers"
)

func mapUrls() {

	router.Use(guidMiddleware())
	hClient := GetHttpClient()
	usr := handlers.NewUserBaseService(cassandra.GetSession()).UserServ.(*handlers.UserBaseService)
	gnews := handlers.NewGNews(hClient, cassandra.GetSession()).Service.(*handlers.GNewsService)
	ping := handlers.Ping
	router.GET("/ping", ping)
	router.POST("/subscribers", usr.Subscribe)
	router.GET("/subscribers/all", usr.Subscribed)
	router.GET("/sources", gnews.GetSources)
	router.GET("/headlines", gnews.GetHeadlines)
	router.GET("/news", gnews.GetNews)
}
func guidMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New()
		c.Set("uuid", uuid)
		fmt.Printf("The request with uuid %s is started \n", uuid)
		c.Next()
		fmt.Printf("The request with uuid %s is served \n", uuid)
	}
}
func GetHttpClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	return &http.Client{
		Transport: tr,
	}
}
