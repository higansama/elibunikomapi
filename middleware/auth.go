package middleware

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	fmt.Println("terpanggil")
	// var tokenString string
	// session := sessions.Default(c)
	// if v := session.Get(framework.Config("jwtName")); v == nil {
	// 	notfound(session, c)
	// 	c.Abort()
	// 	return
	// } else {
	// 	tokenString = v.(string)
	// }

	// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	if jwt.GetSigningMethod("HS256") != token.Method {
	// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	// 	}

	// 	return []byte(framework.Config("jwtKey")), nil
	// })
	// if token == nil {
	// 	notfound(session, c)
	// 	c.Abort()
	// 	return
	// }

	// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid && err == nil {
	// 	var data_jwt map[string]interface{}
	// 	data_jwt = map[string]interface{}{}
	// 	data_jwt["id"] = claims["id"]
	// 	data_jwt["nama"] = claims["nama"]
	// 	data_jwt["email"] = claims["email"]
	// 	data_jwt["role"] = claims["role"]
	// 	c.Set("jwt", data_jwt)
	// } else {
	// 	notfound(session, c)
	// 	c.Abort()
	// 	return
	// }
}

func notfound(session sessions.Session, c *gin.Context) {
	// session.Delete(framework.Config("jwtName"))
	// session.Save()
	// c.Redirect(http.StatusFound, "/")
}
