package jwt

import (
	"net/http"
	"os"
	"time"

	response "github.com/NicholasLiem/GoLang_Microservice/http"
	"github.com/golang-jwt/jwt/v5"
)

type JWTToken struct {
	Token string `json:"token"`
}

func CreateJWT() (JWTToken, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour).Unix(),
		"name": "Testing",
	})

	tokenString, err := token.SignedString(signingKey)

	return JWTToken{Token: tokenString}, err // Fixed token structure and variable name
}

func VerifyJWT(tokenStr string) (jwt.Claims, error) {
	signingKey := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims, nil // Return claims directly
}

func GetJwt(rw http.ResponseWriter, r *http.Request) {
	if apiKey := r.Header.Get("API_KEY"); apiKey != "" {
		if apiKey == os.Getenv("API_KEY") {
			token, err := CreateJWT()
			if err != nil {
				response.ErrorResponse(rw, http.StatusBadRequest, "Fail to generate JWT token")
				return
			}

			responseJson := map[string]string{
				"token": token.Token,
			}

			response.SuccessResponse(rw, http.StatusCreated, "Successfully generated JWT token", responseJson)
			return
		} else {
			response.ErrorResponse(rw, http.StatusBadRequest, "Wrong API key")
			return
		}
	} else {
		response.ErrorResponse(rw, http.StatusBadRequest, "API key header not found")
		return
	}
}
