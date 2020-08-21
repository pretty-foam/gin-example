package controller

import (
	"time"
	"github.com/dchest/captcha"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//User 用户
type User struct{
	UserID int  `json:"id"`
	Name string  `json:"name"`
	jwt.StandardClaims
}


// token key
var jwtKey =[]byte("http://www.baidu.com")
var tokenInfo string 

type userDTO struct{
	Name string `form:"name" binding:"required"` //名称
	Password string `form:"password" binding:"required"` //密码
} 




// Test 测试接口 
// @Tags 用户
// @Summary 测试接口
// @Accept  mpfd
// @Success 200
// @Router /user/test [get]
func (u *User) Test(ctx *gin.Context){
	ctx.JSON(200,gin.H{
		"status":200,
		"msg":"num",
	})
}


//Get 获取用户信息
// @Tags 用户
// @Summary 获取用户信息
// @Param info body userDTO  true "--"
// @Accept  json
// @Success 200
// @Router /user/get [post]
func (u *User) Get(ctx  *gin.Context){	
	params := userDTO{}
	if err := ctx.Bind(&params);err !=nil {
		ctx.String(400,err.Error())
		ctx.Abort()
		return 
	}
	ctx.JSON(200,gin.H{
		"status":200,
		"msg":"获取用户信息",
	})
}

//Captcha 获取验证码
// @Tags 用户
// @Summary 获取验证码
// @Success 200
// @Router /user/captcha [get]
func (u *User) Captcha(ctx *gin.Context)  {
	//验证长度生成ID
	captchaID := captcha.NewLen(4)
	session := sessions.Default(ctx)
	session.Set("captcha",captchaID)
	err := session.Save()
	if err != nil {
		ctx.JSON(500,gin.H{
			"message":err.Error(),
		})
		return
	}
	captcha.WriteImage(ctx.Writer,captchaID,200,80)
}

//CaptchaVerify 验证码验证
// @Tags 用户 
// @Summary 验证码验证
// @param value query number true "验证值"
// @Success 200
// @Router /user/captchaVerify [get]
func (u *User) CaptchaVerify(ctx *gin.Context)  {
	session := sessions.Default(ctx)
	captchaID := session.Get("captcha")
	session.Delete("captcha")
	if captchaID == nil {
		ctx.JSON(200,gin.H{
			"status":500,
			"msg":"session id为空",
		})
		ctx.Abort()
		return
	}
	val := ctx.Query("value")
	if captcha.VerifyString(captchaID.(string),val) {
		ctx.JSON(200,gin.H{
			"status":200,
			"message":"验证成功",
		})
	}else {
		ctx.JSON(200,gin.H{
			"status":500,
			"message":"验证码错误",
		})
		ctx.Abort()
	}
}

// Setjwt 设置jwt
// @Tags 用户
// @Summary 设置jwt
// @Success 200
// @Router /user/login [get]
func (u *User) Setjwt(ctx *gin.Context)  {
	//配置
	config := &User{
		UserID: 123213,
		Name: "测试",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60*time.Second).Unix(), //过期时间
			IssuedAt: time.Now().Unix(),
			Issuer: "1334875424@qq.com", //签名颁发者'
			Subject: "user token",  //签名主题
		},
	}
	//token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,config)
	tokenString,err := token.SignedString(jwtKey)
	if err != nil {
		ctx.JSON(200,gin.H{
			"status":500,
			"message":err.Error(),
		})
		ctx.Abort()
		return 
	}
	info,_ := jwt.Parse(tokenString,func (token *jwt.Token)(info interface{} ,err error)  {
		return info,nil
	})
	ctx.JSON(200,gin.H{
		"status":200,
		"message":info,
	})
}