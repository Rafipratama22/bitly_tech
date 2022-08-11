package middleware

import (
	"bitly_backend/app/models"
	"fmt"
	"net/http"

	// "strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gomodul/envy"
	"gorm.io/gorm"
)

type md struct {
	db *gorm.DB
}

// type tokenMap struct {
// 	UserID string `json:"user_id"`
// 	Authorized bool `json:"authorized"`
// }

func NewMiddleware(db *gorm.DB) Middlewared {
	return &md{
		db,
	}
}

func (c *md) ValidatedToken(ctx *gin.Context) {
	tokenMap := map[string]float64{
		"user_id": 0,
	}
	secret_jwt := envy.Get("JWT_KEY", "jwt")
	tokenString := ctx.GetHeader("authorization")
	user := models.User{}
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, map[string]string{
			"message": "your not allowed 1",
		})
		ctx.Abort()
	}
	fmt.Println(tokenString)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) { //Parse token
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", "secret")
		}
		return []byte(secret_jwt), nil
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, map[string]string{
			"message": "your not allowed 2",
		})
		ctx.Abort()
	}
	if claims, err := token.Claims.(jwt.MapClaims); err && token.Valid {
		//Validate token
		fmt.Println("claims", claims)
		for key, val := range claims {
			if key == "user_id" {
				s, ok := val.(float64)
				if ok {
					tokenMap[key] = s
				}
			} else {
				continue
			}
		}
	}
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, map[string]string{
			"message": "your not allowed 3",
		})
		ctx.Abort()
	}
	fmt.Println(tokenMap["user_id"])
	// inte, _ := strconv.Atoi(tokenMap["user_id"])
	if err = c.db.Model(&user).Where("id = ?", tokenMap["user_id"]).First(&user).Error; err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, map[string]string{
			"message": "your not allowed 4",
		})
		ctx.Abort()
	}

	ctx.Set("user_id", tokenMap["user_id"])
	ctx.Next()
}
