package controllers

import "github.com/gin-gonic/gin"

/*
APIController is interface defined to guarantee API separation of context in various modules
*/
type APIController interface {
	Register(e *gin.Engine)
}
