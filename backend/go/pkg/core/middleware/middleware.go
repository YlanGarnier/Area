package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lenismtho/area/pkg/core/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type dbMiddleware interface {
	GetUserByToken(token string) (*models.User, error)
}

type Middleware struct {
	db dbMiddleware
}

func NewMiddleware(db dbMiddleware) (*Middleware, error) {
	return &Middleware{
		db: db,
	}, nil
}

func (m *Middleware) IsValidToken(c *gin.Context) {
	if len(c.Request.Header["Authorization"]) == 0 {
		c.AbortWithStatusJSON(http.StatusForbidden, "No \"Authorization\" token provided")
		return
	}
	headerTab := strings.Split(c.Request.Header["Authorization"][0], " ")
	if len(headerTab) == 0 || headerTab[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusForbidden, "Invalid token")
		return
	}
	if len(headerTab) <= 1 {
		c.AbortWithStatusJSON(http.StatusForbidden, "Invalid token")
		return
	}
	token := headerTab[1]
	_, err := m.db.GetUserByToken(token)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(http.StatusForbidden, "Invalid token")
		return
	} else if err != nil {
		log.WithField("error", err).Error("failed to parse token")
		c.AbortWithStatusJSON(http.StatusForbidden, "Invalid token")
		return
	}
	c.Next()
}
