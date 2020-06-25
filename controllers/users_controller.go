package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"net/http"

	orm "github.com/hongyukeji/easy-go/datasource"
	"github.com/hongyukeji/easy-go/models"
)

type UsersController struct{}

// 用户列表
func (c *UsersController) GetUsers(ctx iris.Context) {
	var users []models.User
	orm.Eloquent.Find(&users)
	ctx.JSON(iris.Map{"message": "pong", "data": users})
}

// 新增用户
func (c *UsersController) PostUsers(ctx iris.Context) {
	/*//获取请求path
	path := ctx.Path()
	//日志
	ctx.Application().Logger().Info(path)
	//获取请求数据字段
	postForm := ctx.FormValues()
	name := ctx.PostValue("username")
	pwd, err := ctx.PostValueInt("password")
	if err != nil {
		ctx.Application().Logger().Fatalf("created one record failed: %s", err.Error)
		ctx.JSON(iris.Map{
			"code":  http.StatusBadRequest,
			"error": err.Error,
		})
		return
	}
	ctx.Application().Logger().Info(name, "  ", pwd)
	//返回
	ctx.JSON(
		iris.Map{
			"code": http.StatusOK,
			"data": postForm,
		})*/
	var user models.User
	tx := orm.Eloquent.Begin()

	var body putParam
	ctx.ReadJSON(&body)
	ctx.Application().Logger().Println(body)
	user = models.User{
		Username:  body.Data.Username,
		Salt:      body.Data.Salt,
		Password:  body.Data.Password,
		Languages: body.Data.Languages,
	}
	if err := tx.Create(&user); err == nil {
		ctx.Application().Logger().Fatalf("created one record failed: %s", err.Error)
		ctx.JSON(iris.Map{
			"code":  http.StatusBadRequest,
			"error": err.Error,
		})
		return
	}
	ctx.JSON(
		iris.Map{
			"code": http.StatusOK,
			"data": user.Serializer(),
		})
}

// 查看用户
func (c *UsersController) GetUsersBy(id uint, ctx iris.Context) {
	var user models.User
	if err := orm.Eloquent.Where("id = ?", int(id)).First(&user).Error; err != nil {
		ctx.JSON(iris.Map{
			"code":  http.StatusBadRequest,
			"error": err.Error,
		})
		return
	}
	ctx.JSON(iris.Map{
		"code": http.StatusOK,
		"data": user.Serializer(),
	})
}

// 修改用户
func (c *UsersController) PutUsersBy(id uint, ctx iris.Context) {
	//id, _ := ctx.Params().GetUint("id")
	if id == 0 {
		ctx.JSON(iris.Map{
			"code":   http.StatusOK,
			"detail": "query param id should not be nil",
		})
		return
	}
	var user models.User
	tx := orm.Eloquent.Begin()
	if err := tx.Where("id = ?", id).First(&user).Error; err != nil {
		ctx.Application().Logger().Fatalf("record not found")
		ctx.JSON(iris.Map{
			"code":   http.StatusOK,
			"detail": err.Error,
		})
		return
	}

	var body putParam
	ctx.ReadJSON(&body)
	ctx.Application().Logger().Println(body)
	if err := tx.Model(&user).Updates(map[string]interface{}{"username": body.Data.Username, "password": body.Data.Password}).Error; err != nil {
		ctx.Application().Logger().Fatalf("update record failed")
		tx.Rollback()
		ctx.JSON(iris.Map{
			"code":  http.StatusBadRequest,
			"error": err.Error,
		})
		return
	}
	tx.Commit()
	ctx.JSON(iris.Map{
		"code": http.StatusOK,
		"data": user.Serializer(),
	})
}

// 删除用户
func (c *UsersController) DeleteUsersBy(id uint, ctx iris.Context) {
	//id, _ = ctx.Params().GetUint("id")
	if id == 0 {
		ctx.JSON(iris.Map{
			"code":   http.StatusOK,
			"detail": "query param id should not be nil",
		})
		return
	}
	var user models.User
	if err := orm.Eloquent.Where("id = ?", id).First(&user).Error; err != nil {
		ctx.Application().Logger().Fatalf("record not found")
		ctx.JSON(iris.Map{
			"code":   http.StatusOK,
			"detail": err.Error,
		})
		return
	}
	orm.Eloquent.Delete(&user)
	ctx.JSON(iris.Map{
		"code": http.StatusOK,
		"data": user.Serializer(),
	})
}

type putParam struct {
	Data struct {
		Username  string `json:"username" form:"username"`
		Password  string `json:"password" form:"password"`
		Salt      string `json:"salt" form:"salt"`
		Languages string `json:"languages" form:"languages"`
	} `json:"data"`
}

// 激活后，所有依赖项都被设置为只读访问
func (c *UsersController) AfterActivation(a mvc.AfterActivation) {}

// 在激活前调用一次，在控制器适应主应用程序之前
func (c *UsersController) BeforeActivation(b mvc.BeforeActivation) {
	// b.Dependencies().Add/Remove
	// b.Router().Use/UseGlobal/Done
	// 和已知的任何标准 API 调用

	// 1-> 方法
	// 2-> 路径
	// 3-> 将控制器的函数名称解析为处理程序
	// 4-> 应该在 MyCustomHandler 之前运行的任何处理程序
	//b.Handle("GET", "/users/{id:uint}", "GetUser")
}
