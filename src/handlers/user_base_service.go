package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/laithrafid/news_fetch/src/cassandra"
	"github.com/laithrafid/news_fetch/src/domain"
	"github.com/laithrafid/news_fetch/src/repository"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	
// subscribed user if new user func
func Subscribed(c *gin.Context) {
	var respList []domain.UserDetails
	if results, err := .GetAllUser(); err != nil {
		fmt.Errorf("Error while fetching subscribers. Returning empty response")
	} else {
		respList = results
	}
	c.JSON(http.StatusOK, respList)
}

func Subscribe(c *gin.Context) {
	if body, err := ioutil.ReadAll(c.Request.Body); err != nil {
		fmt.Printf("Error encountered : %v", err.Error())
	} else {
		usrDet := new(domain.UserDetails)
		err := json.Unmarshal(body, &usrDet)
		if err != nil {
			fmt.Printf("Could not process the subscription request. Error encountered : %v", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"Error: ": err.Error()})
			return
		} else {
			if err := u.CheckAndPersist(usrDet); err != nil {
				c.JSON(http.StatusOK, gin.H{"Error: ": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"Success": "true"})
		}
	}
}

func CheckAndPersist(usrDet *domain.UserDetails) error {
	if subscriber, err := u.UserDbClient.GetUserByTgDetils(int(usrDet.ID), usrDet.Name); err != nil || subscriber.ID == 0 {
		fmt.Printf("New subscriber with Id : %d and Username : %s", usrDet.ID, usrDet)
		m := domain.UserDetails{
			ID:   usrDet.ID,
			Name: usrDet.Name,
		}
		if err := .InsertUser(m); err != nil {
			fmt.Printf("Failure while persisting subscriber with Id : %d and Username : %s - %s", subscriber.ID, subscriber.Name, err.Error())
			return err
		}
	} else {
		fmt.Printf("Subscriber found with Id : %d and Username : %s", subscriber.ID, subscriber.Name)
		return err
	}
	return nil
}
