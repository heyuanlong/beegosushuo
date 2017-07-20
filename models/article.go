package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"strconv"
	"time"
)

type Article struct {
	Id       		int
	Title    		string
	Content  		string
	Userid   		int
	Time     		string
	Views    		int
	Comments 		int
	Classifyid   	int
}

func init() {
	orm.RegisterModel(new(Article))
}

func GetArticle(id int) (Article, error) {
	var err error
	var art Article

	o := orm.NewOrm()
	o.Using("default")
	art = Article{Id: id}
	err = o.Read(&art, "id")

	return art, err
}
func UpdateArticleViews(id int) () {
	o := orm.NewOrm()
	o.Using("default")
	o.Raw("update article set views = views +1 where id=?",id).Exec()

}
func UpdateArticleComments(id int,status int) () {
	o := orm.NewOrm()
	o.Using("default")
	o.Raw("update article set Comments = Comments +1 where id=?",id).Exec()
}
// 添加文章
func AddArticle(title string, content string, userid int ,author string, status string) (int64, error) {
	o := orm.NewOrm()
	o.Using("default")
	ctime := time.Now()

	sql := "insert into article(title, content, userid ,time,views ) values(?, ?, ?, ?, ?)"
	res, err := o.Raw(sql, title, content, userid, ctime,1).Exec()
	if nil != err {
		return 0, err
	} else {
		return res.LastInsertId()
	}
}
func AllArticleCout(userid int)( int, error) {
	var sql string
	if userid < 0 {
		userid=1
		sql = "select count(*) as number  from article where 1=?"
	}else{
		sql = "select count(*) as number  from article where userid=?"
	}
	
	var maps []orm.Params
	o := orm.NewOrm()
	_, err := o.Raw(sql,userid).Values(&maps)
	if err != nil {
		fmt.Println("execute sql error:"+sql)
		fmt.Println(err)
		return 0, err
	}
	return strconv.Atoi(maps[0]["number"].(string))
}

func ListArticle(page int,numPerPage int,userid int)([]orm.Params, int64, error) {
	var sql string
	if page <=0 {
		page = 1
	}
	if userid < 0 {
		userid = 1
		sql = "select * from article where 1=? order by time desc limit ?," + fmt.Sprintf("%d", numPerPage)
	}else{
		sql = "select * from article where userid=? order by time desc limit ?," + fmt.Sprintf("%d", numPerPage)
	}

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(sql,userid, numPerPage*(page-1)).Values(&maps)
	if err != nil {
		fmt.Println("execute sql error:" + sql)
		fmt.Println(err)
		return nil, 0, err
	}

	return maps,num,nil
}
// 最热查看列表 
func HottestViewsList(num int) ([]orm.Params, error) {
	var maps []orm.Params

	sql := "select id,title,views from article  order by views desc limit ?"
	o := orm.NewOrm()
	_, err := o.Raw(sql,num).Values(&maps)

	return maps, err
}
// 最热评论列表 
func HottestCommentsList(num int) ([]orm.Params, error) {
	var maps []orm.Params

	sql := "select id,title,Comments from article where order by Comments desc limit ?"
	o := orm.NewOrm()
	_, err := o.Raw(sql,num).Values(&maps)

	return maps, err
}



/*func UpdateArticle(id int,title string,  keywords string,content string,classifyid string ) (error) {
	o := orm.NewOrm()
	o.Using("default")
	_ ,err := o.Raw("update article set title = ?,keywords=?,content=?,classifyID=? where id=?",title,keywords,content,classifyid,id).Exec()
	return err
}
*/


/*
func AllArticle()([]orm.Params, int64, error) {
	sql := "select * from article where status != -1 order by time desc"
	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(sql).Values(&maps)
	if err != nil {
		fmt.Println("execute sql1 error:"+sql)
		fmt.Println(err)
		return nil, 0, err
	}
	return maps,num,nil
}*/

/*func AllArticleAdmin()([]orm.Params, int64, error) {
	sql := "select id,title,time,status from article where status != -1 order by time desc"
	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw(sql).Values(&maps)
	if err != nil {
		fmt.Println("execute sql1 error:"+sql)
		fmt.Println(err)
		return nil, 0, err
	}
	return maps,num,nil
}*/



