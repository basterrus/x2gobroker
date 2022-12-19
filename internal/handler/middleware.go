package handler

import "github.com/gin-gonic/gin"

type QParams struct {
	user     string
	login    string
	password string
	task     string
	sid      string
	hwid     string
}

func (h *Handler) ParsQueryParams(c *gin.Context) {
	user := c.Query("user")
	login := c.Query("login")
	password := c.Query("password")
	task := c.Query("task")
	sid := c.Query("sid")
	hwid := c.Query("hwid")

}
