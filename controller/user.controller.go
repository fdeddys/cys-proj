package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fdeddys/tes/dto"
	"github.com/fdeddys/tes/service"
	"github.com/gin-gonic/gin"
)

// UserController ...
type UserController struct {
}

// UserService ...
var UserService = new(service.UserService)

// SaveDataUser ...
func (h *UserController) SaveDataUser(c *gin.Context) {
	req := dto.RequestUser{}
	res := dto.Response{}

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)

	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		log.Println("failed unmarshal body request ")
		res.StatusCode = 400
		res.StatusDesc = "failed unmarshal body request "
		c.JSON(http.StatusOK, res)
		return
	}

	res.StatusCode = 200
	res.StatusDesc = "OK"
	errResp := UserService.SaveUser(&req)
	if errResp != nil {
		res.StatusCode = 500
		res.StatusDesc = errResp.Error()
	}

	c.JSON(http.StatusOK, res)
}

// SaveDataUser ...
func (h *UserController) ListResources(c *gin.Context) {
	req := dto.RequestGather{}
	res := dto.Response{}

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)
	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		log.Println("failed unmarshal body request ")
		res.StatusCode = 400
		res.StatusDesc = "failed unmarshal body request "
		c.JSON(http.StatusOK, res)
		return
	}

	token := c.Request.Header.Get("Authorization")
	if token == "" {
		res.StatusCode = 401
		res.StatusDesc = "Token not provided "
		c.JSON(http.StatusOK, res)
		return
	}

	code, desc, data := UserService.ListResourceByUsername(req.Username, token)
	res.StatusCode = code
	res.StatusDesc = desc
	res.Content = data
	c.JSON(http.StatusOK, res)
}

// SaveDataUser ...
func (h *UserController) GatherData(c *gin.Context) {
	req := dto.RequestGather{}
	res := dto.Response{}

	body := c.Request.Body
	dataBodyReq, _ := ioutil.ReadAll(body)
	if err := json.Unmarshal(dataBodyReq, &req); err != nil {
		log.Println("failed unmarshal body request ")
		res.StatusCode = 400
		res.StatusDesc = "failed unmarshal body request "
		c.JSON(http.StatusOK, res)
		return
	}

	token := c.Request.Header.Get("Authorization")
	if token == "" {
		res.StatusCode = 401
		res.StatusDesc = "Token not provided "
		c.JSON(http.StatusOK, res)
		return
	}

	code, desc := UserService.GatherService(req.Username, token)
	res.StatusCode = code
	res.StatusDesc = desc

	c.JSON(http.StatusOK, res)
}
