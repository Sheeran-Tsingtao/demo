package controller

import (
	"github.com/gin-gonic/gin"
	"nano/data"
	"nano/model"
	"net/http"
)

func PostWebsite(c *gin.Context)  {
	DB := data.GetDB()
	name := c.PostForm("name")
	website := c.PostForm("website")
	introduction := c.PostForm("introduction")
	if website == " "{
		c.JSON(http.StatusOK,gin.H{"msg":"请输入网站"})
		return
	}

	u :=model.Webcount{
		Name: name,
		Website: website,
		Introduction: introduction,

	}
	DB.Create(&u)
	var o model.Allweb
	DB.Debug().Where("website = ?",website).First(&o)

	if o.ID == 0{
		DB.Create(&model.Allweb{Website: website,Num: 1})
		c.JSON(http.StatusOK,gin.H{"msg":"添加成功，恭喜你是第一个添加网站的"})
		return
	}
	o.Num=o.Num+1
	DB.Save(&o)
	c.JSON(http.StatusOK,gin.H{"msg":"添加成功"})


}
func GetUserWebsite(c *gin.Context)  {
	DB := data.GetDB()
	var web []model.Webcount
	name := c.PostForm("name")
	DB.Where("name = ?", name).Find(&web)
	if web[0].ID == 0 {
		c.JSON(402,gin.H{"msg":"没有收藏过一个网页"})
		return
	}

	c.JSON(http.StatusOK,web)


}
func ShowWebsite(c *gin.Context)  {
	DB := data.GetDB()
	var web []model.Allweb

	DB.Find(&web)
	//fmt.Printf("%s\n",web[0].Website)
	if len(web)==0 {
		c.JSON(402,gin.H{"msg":"没有收藏过一个网页"})
		return
	}
	c.JSON(http.StatusOK,web)


}