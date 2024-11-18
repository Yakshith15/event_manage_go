package utils

import (
    "fmt"
    "time"  
    // "log"
	"os"
    "github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte(os.Getenv("JWT_SECRET")) // Use a more secure secret in production
// log.Println("JWT_SECRET:", secretKey)
// GenerateJWT generates a new JWT token
func GenerateJWT(email string, role string) (string, error) {
    claims := jwt.MapClaims{
        "email": email,
        "role":  role,
        "exp":   time.Now().Add(time.Hour * 24).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString(secretKey)
    if err != nil {
        return "", fmt.Errorf("could not generate token: %v", err)
    }

    return tokenString, nil
}

// ParseJWT parses the JWT token
func ParseJWT(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
        }
        return secretKey, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }
    return nil, fmt.Errorf("invalid token")
}
