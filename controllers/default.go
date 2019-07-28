package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/spf13/viper"
	"strconv"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
func (c *MainController) Gen() {
	c.Data["json"] = map[string]interface{}{"msg":"Token失效"}
	c.ServeJSON()
}
func (c *MainController) Login() {
	var keys map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &keys); err != nil {
		c.Data["json"] = map[string]interface{}{"msg":"jsonErr"}
		return
	}
	userInfo := make(map[string]interface{})
	if keys["name"].(string) == "admin"{
		userInfo["exp"] = strconv.FormatInt(time.Now().Unix() + 604800, 10) // 1周
		userInfo["iat"] = time.Now().Unix()
		userInfo["iss"] = "SanJi_Fabric"
		token := c.createToken(beego.AppConfig.String("key"),userInfo)
		c.Data["json"] = map[string]interface{}{"msg":"ok","Authorization":token}

	} else {
		c.Data["json"] = "no"
	}
	c.ServeJSON()
}
func (c *MainController) XiuGai() {
	c.Data["json"] = "ok OK"
	c.ServeJSON()
}
func (c *MainController) createToken(key string, m map[string] interface{}) string{
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	for index, val := range m {
		claims[index] = val
	}
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString
}
func CreateToken()(tokenss string,err error){
	//自定义claim
	claim := jwt.MapClaims{
		"nbf":      int64(time.Now().Unix() - 1000),
		"exp":      int64(time.Now().Unix() + 3600),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	tokenss,err  = token.SignedString([]byte(viper.GetString("token.secret")))
	return
}
