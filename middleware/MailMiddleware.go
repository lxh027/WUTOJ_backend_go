package middleware

import (
	"OnlineJudge/app/common"
	"OnlineJudge/app/helper"
	"OnlineJudge/config"
	"OnlineJudge/db_server"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
)

func SendMail(EmailAddress string) (helper.ReturnType, error) {

	mailConfig := config.GetMailConfig()

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	VerifyCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	keyValue := "VerifyCode" + EmailAddress
	// save to redis
	db_server.PutToRedis(keyValue, VerifyCode, 1000*60*15)

	SendTime := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	html := fmt.Sprintf(`<div>
		<div>
			尊敬的%s, 您好!
		</div>
		<div style="padding: 8px 40px 8px 50px;">
			<p> 您于 %s 提交了邮箱验证，本次验证码为 %s，为了保证账号安全，验证码有效期为15分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解和使用。 </p>
		</div>
		<div>
			<p> 次邮箱为系统邮箱，请勿回复。</p>
		</div>
	<div>`, EmailAddress, SendTime, VerifyCode)

	message := gomail.NewMessage()
	message.SetAddressHeader("From", mailConfig["from"].(string), mailConfig["from_name"].(string))
	message.SetHeader("To", EmailAddress)
	message.SetHeader("Subject", "[我的验证码]邮箱验证")
	message.SetBody("text/html", html)

	dia := gomail.NewDialer(mailConfig["host"].(string), mailConfig["port"].(int), mailConfig["username"].(string), mailConfig["password"].(string))

	if err := dia.DialAndSend(message); err != nil {
		return helper.ReturnType{Status: common.CodeError, Msg: "邮件发送失败", Data: err.Error()}, err
	}
	return helper.ReturnType{Status: common.CodeSuccess, Msg: "邮件发送成功，请注意查收", Data: ""}, nil
}

func GetUserIdFromSession(c *gin.Context) uint {
	session := sessions.Default(c)
	if id := session.Get("user_id"); id != nil {
		return id.(uint)
	}
	return 0
}
