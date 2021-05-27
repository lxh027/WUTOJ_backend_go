## 项目目录

    ├─app
    │  ├─api
    │  │  ├─controller
    │  │  │  └─IndexController.go
    │  │  ├─model
    │  │  │  └─IndexModel.go
    │  │  └─validate
    │  │     └─IndexValidate.go
    │  └─common
    ├─config
    ├─db_server
    ├─judger
    ├─log
    ├─routes
    ├─server
    ├─web
    ├─go.mod
    ├─main.go

## ORM

使用GORM库，文档地址: [gorm doc](https://gorm.io/docs/) [gorm文档翻译](https://jasperxu.com/#/Programming/Golang/GORM/)

## Gin

[gin docs](gin-gonic.com/docs)

## Session

使用`gin-contrib\sessions`库，使用方法如下：

```go
// 初始化session
session := sessions.Default(c)
// 保存session
session.Set("user_id", userInfo.UserID)
session.Set("nick", userInfo.Nick)
session.Set("identity", userInfo.Identity)
session.Set("data", string(jsonData))
session.Save()
// 取出session
if session.Get("user_id") != nil {
	//ops
}
```

## Redis

使用`redigo`库，使用方法如下：

```go
// 放入redis, (key, value, timeout)
_ = db_server.PutToRedis("k", "132", 3600)
// 取出
k, err := GetFromRedis("m")
// 取出并定义类型
mpBytes, _ := redis.String(k, err)
```

## Validate

使用`gookit/validate`库，使用方法如下：

```go
if res, err := userValidate.ValidateMap(userMap, "login"); !res {
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "输入信息不完整或有误", err.Error()))
	return
}
```

数据验证定义在`app/common/validate`内

数据的校验需要在数据绑定后校验，否则会出现校验失败

完整的校验如下：

```go
var userJson model.User

if err := c.ShouldBind(&userJson); err != nil {
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "数据绑定模型错误", err.Error()))
	return
}

userMap := helper.Struct2Map(userJson)
if res, err := userValidate.ValidateMap(userMap, "login"); !res {
	c.JSON(http.StatusOK, helper.ApiReturn(common.CodeError, "输入信息不完整或有误", err.Error()))
	return
}
```

## Panel

后台使用`layui`模板，放于`web/`内

后台api模块为`app/panel`

## Todo List
- [ ] 增加数据库操作队列，优化数据库性能，减少数据丢失
- [ ] 后台增加队伍账号导入模块，导入队名csv文件，生成比赛账号密码，自动注册并选择要报名的比赛。