package home

import (
	"github.com/astaxie/beego"
	. "beegosushuo/models"
	"strconv"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	var page int
	var prev_page_flag bool
	var next_page_flag bool
	perPageNums := 5

	id := this.Ctx.Input.Param(":id")
	pageParm, err := strconv.Atoi(id)
	if nil != err || pageParm < 0 {
		page = 1
	} else {
		page = pageParm
	}

	allnums ,err_ := AllArticleCout(-1)
	if err_ != nil {
	 	allnums = 0
	 } 

	maps, _, err__ := ListArticle(page,perPageNums,-1)
	if nil != err__ {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "get list failed", "refer": "/"}
		this.ServeJSON()
	} else {
		if page > 1 {
			prev_page_flag = true
			this.Data["prev_page_flag"] = prev_page_flag
			this.Data["prev_page"] = fmt.Sprintf("/%d", page-1)
		}
		if (page * perPageNums) < allnums {
			next_page_flag = true
			this.Data["next_page_flag"] = next_page_flag
			this.Data["next_page"] = fmt.Sprintf("/%d", page+1)
		}

		hottestviews,err:= HottestViewsList(15)
		if nil == err{
			this.Data["hottestviews"] = hottestviews
		}

		this.Data["list"] = maps

		account , ok := this.GetSession("account").(string)
	    if ok  {
	        this.Data["account"] = account
	    }else{
	    	this.Data["account"] = ""
	    }

		this.TplName = "home/index.tpl"
	}

}

func (this *MainController) OneArticle() {
	var articleId int
	id := this.Ctx.Input.Param(":id")
	id_, err := strconv.Atoi(id)
	if nil != err || id_ < 0 {
		articleId = 1
	} else {
		articleId = id_
	}

	art, err_ := GetArticle(articleId)
	if err_ != nil {
		this.Data["json"] = map[string]interface{}{"result": false, "msg": "get OneArticle failed", "refer": "/"}
		this.ServeJSON()
		return 
	}
	UpdateArticleViews(articleId)
	this.Data["id"] = art.Id
	this.Data["title"] = art.Title
	this.Data["time"] = art.Time
	this.Data["views"] = art.Views
	this.Data["comments"] = art.Comments
	this.Data["content"] = art.Content

	hottestviews,err:= HottestViewsList(15)
	if nil == err{
		this.Data["hottestviews"] = hottestviews
	}

	this.TplName = "home/onearticle.tpl"

}





