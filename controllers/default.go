package controllers

import (
	"HelloBeego06/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller //匿名字段
}

func (c *MainController) Get() {
	//name1 :=c.GetString("name")
	//age1,err :=c.GetInt("age")
	//获取get类型请求的请求参数
	name := c.Ctx.Input.Query("name")
	age := c.Ctx.Input.Query("age")
	sex := c.Ctx.Input.Query("sex")
	fmt.Println(name, age, sex)
	//以admin,18为正确数据进行验证
	if name != "admin" || age != "18" {
		c.Ctx.ResponseWriter.Write([]byte("数据验证错误"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("数据验证成功"))
}

/*
该post方法是处理post类型的请求的时候，要调用的方法
*/
/*func(c *MainController) Post(){
	fmt.Println("post类型的请求")
	user :=c.Ctx.Request.FormValue("user")
	fmt.Println("用户名为：",user)
	psd :=c.Ctx.Request.FormValue("psd")
	fmt.Println("密码是:",psd)


	//与固定值比较  用户名为：admin  密码是123456
	if user != "admin" || psd != "123456"{
		//失败页面
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据不正确"))
		return
	}

	c.Ctx.ResponseWriter.Write([]byte("成功了，数据正确"))

	//request:请求  response响应

}*/
/*func (c *MainController) Post(){
	dataByes,err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据接收失败")
		return
	}
	//JSON包解析
	var person models.Person
	err = json.Unmarshal(dataByes,&person)
	if err != nil{
		c.Ctx.WriteString("数据解析失败,请重试")
		return
	}
	fmt.Println("用户名是:",person.User,",年龄",person.Age)
	c.Ctx.WriteString("用户名:"+person.User)
	c.Ctx.WriteString("性别是:"+person.Sex)
}*/
func (c *MainController) Post() {
	name := c.Ctx.Request.FormValue("name")
	birthday := c.Ctx.Request.FormValue("birthday")
	address := c.Ctx.Request.FormValue("address")
	nick := c.Ctx.Request.FormValue("nick")
	fmt.Println(name, birthday, address, nick)
	//用得到的数据进行对比
	if name != "wujiali" && address != "dongyang" {
		c.Ctx.WriteString("数据接收错误，重试啊！")
		return
	}
	c.Ctx.WriteString("数据接收成功，恭喜!!")

	var person models.Person
	data, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据接收错误,请重试尝试")
		return
	}
	err = json.Unmarshal(data, &person)
	if err != nil {
		c.Ctx.WriteString("数据解析错误")
		return
	}
	fmt.Println("姓名:", person.Name)
	fmt.Println("生日:", person.Birthday)
	fmt.Println("地址:", person.Address)
	fmt.Println("昵称", person.Nick)
	c.Ctx.WriteString("数据请求成功")
}