package filter

import (
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/spf13/viper"
	"github.com/pkg/errors"
)
var tokenState string

var AuthFilter = func(ctx *context.Context) {
	header := ctx.Input.Header("Authorization")
	s := beego.AppConfig.String("key")

	claims, ok := parseToken(header, s)
	if ok {
		/*
		oldT, _ := strconv.ParseInt(claims.(jwt.MapClaims)["exp"].(string), 10, 64)
		beego.Error("oldT: ",oldT)
		ct := time.Now().Unix()
		// 如果当前时间 大于 Token中的时间 = 过期
		if  ct > oldT{
			ok = false
			tokenState = "Token 已过期"
			ctx.Redirect(302,"/")
		} else {
			tokenState = "Token 正常"
		}
		*/
		tokenState = "Token 正常"
	}else {
		tokenState = "token无效"
		ctx.Redirect(302,"/")
	}
	fmt.Println(tokenState)
	fmt.Println(claims)
}
func parseToken(tokenString string, key string) (interface{}, bool){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println(err)
		return "", false
	}
}


func secret()jwt.Keyfunc{
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("token.secret")),nil
	}
}

var Auth = func(ctx *context.Context){
	header := ctx.Input.Header("Authorization")
	token,err := jwt.Parse(header,secret())
	if err != nil{
		beego.Error("Token解析失败")
		ctx.Redirect(302,"/")
		return
	}

	_,ok := token.Claims.(jwt.MapClaims)
	if !ok{
		err = errors.New("token 失效")
		ctx.Redirect(302,"/")
		return
	}
	//验证token，如果token被修改过则为false
	if  !token.Valid {
		err = errors.New("token 无效")
		ctx.Redirect(302,"/")
		return
	}
}
