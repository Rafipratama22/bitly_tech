package transport

import (
	"bitly_backend/app/usecase"
	"bitly_backend/app/usecase/request"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type transport struct {
	usecase usecase.Usecase
}

func NewTransport(usecase usecase.Usecase) *transport {
	return &transport{
		usecase,
	}
}



func (t transport) Register(ctx *gin.Context) {
	var request request.Register
	err := ctx.ShouldBindJSON(&request)
	fmt.Println(request, "request")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "You have still missing request",
		})
		return
	}
	
	res, err := t.usecase.Register(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Please check your register",
		})
		return
	}
	
	ctx.JSON(http.StatusOK, res)
}

func (t transport) Login(ctx *gin.Context) {
	var request request.Login
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "You have still missing request",
		})
		return
	}
	
	res, err := t.usecase.Login(request)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Please check your request",
		})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"access_token": res,
	})
}

func (t transport) GetListProductByUserID(ctx *gin.Context) {
	user_id := ctx.MustGet("user_id")
	s, ok := user_id.(float64)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Your not allowed",
		})
		return
	}
	res, err := t.usecase.GetListProductByUserID(s)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Please check your request",
		})
		return
	}
	
	ctx.JSON(http.StatusOK, res)
}

func (t transport) CreateUrl(ctx *gin.Context) {
	var request request.CreateUrl
	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "You have still missing request",
		})
		return
	}	
	user_id, _ := ctx.MustGet("user_id").(float64)
	request.UserID = user_id
	res, err := t.usecase.CreateUrl(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Please check your request",
		})
		return
	}
	
	ctx.JSON(http.StatusOK, res)
}

func (t transport) UpdateUrl(ctx *gin.Context) {
	var request request.UpdateUrl
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Sorry there is trouble on machine",
		})
		return
	}
	request.ID = id
	err = ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "You have still missing request",
		})
		return
	}
	res, err := t.usecase.UpdateUrl(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Please check your request",
		})
		return
	}
	
	ctx.JSON(http.StatusOK, res)
}

func (t transport) PatchGetClick(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Sorry there is trouble on machine",
		})
		return
	}
	res, err := t.usecase.PatchGetClick(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Please check your request",
		})
		return
	}
	
	ctx.JSON(http.StatusOK, res)
}

func (t transport) CreateUrlWithoutUser(ctx *gin.Context) {
	var request request.CreateUrlWithoutUser
	err := ctx.ShouldBindJSON(&request)
	fmt.Println(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Sorry there is trouble on machine",
		})
	}
	res, err := t.usecase.CreateUrlWithoutUser(request.DestinationUrl)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Sorry there is trouble on machine",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"code": res,
	})
}

func (t transport) GetInfoUser(ctx *gin.Context) {
	user_id := ctx.MustGet("user_id")
	s, ok := user_id.(float64)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Your not allowed",
		})
		return
	}
	res, err := t.usecase.GetInfoUser(s)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, res)	
}

func (t transport) GetOneProduct(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Sorry there is trouble on machine",
		})
		return
	}
	res, err := t.usecase.GetOneProduct(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, res)
}