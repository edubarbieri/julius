package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/edubarbieri/julius/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

func JwtTokenCheck(c *gin.Context) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := parseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, UnsignedResponse{
			Message: "bad jwt token",
		})
		return
	}

	_, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
			Message: "unable to parse claims",
		})
		return
	}
	c.Next()

}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad Authorization header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func parseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte(config.JwtSecret), nil
	})

	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}
