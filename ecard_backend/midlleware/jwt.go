package midlleware

import (
    "ecard_backend/global"
    "errors"
    jwt "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

var (
    expireDuration = time.Hour * 24
    jwtSecret      = []byte("jwtSecret")
)

type Claims struct {
    Passport string	`json:"passport"`
    jwt.StandardClaims
}

func GenerateToken(passport string) (string, error) {
    expire := time.Now().Add(expireDuration)
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
        Passport:  passport,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expire.Unix(),
        },
    })
    return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
        return jwtSecret, nil
    })
    if err != nil {
        return nil, err
    }
    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }
    return nil, errors.New("invalid token")
}

func Auth() gin.HandlerFunc {
    return func(c *gin.Context) {
        var statusCode int
        var data interface{}

        statusCode = global.SUCCESS
        token := c.Request.Header.Get("Authorization")
        if token == "" {
            statusCode = global.UNAUTHORIZED
        } else {
            _, err := ParseToken(token)
            if err != nil {
                switch err.(*jwt.ValidationError).Errors {
                case jwt.ValidationErrorExpired:
                    statusCode = global.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
                default:
                    statusCode = global.ERROR_AUTH_CHECK_TOKEN_FAIL
                }
            }
        }

        if statusCode != global.SUCCESS {
            c.JSON(http.StatusUnauthorized, gin.H{
                "code": statusCode,
                "msg":  global.GetMsg(statusCode),
                "data": data,
            })

            c.Abort()
            return
        }

        c.Next()
    }
}