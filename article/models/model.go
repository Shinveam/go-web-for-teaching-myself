package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id       int
	UserName string
	Pwd      string
	Article  []*Article `orm:"rel(m2m)"` //用户表与文章表是多对多关系,m2m表示多对多
}

type Article struct {
	Id          int          `orm:"pk;auto"`
	ArtiName    string       `orm:"size(20)"`       //标题
	Content     string       `orm:"size(500)"`      // 内容
	PubTime     string //`orm:"auto_now_add;type(datetime)"` //创建时间
	Img         string       `orm:"size(50);null"` //图片（路径）
	Count       int          `orm:"defualt(0)"`     //浏览量
	ArticleType *ArticleType `orm:"rel(fk)"`        //设置外键(fk)，文章表和文章类型表是一对多的关系
	User        []*User      `orm:"reverse(many)"`  //与`orm:"rel()"`成对存在，与用户表结合表示多对多类型
}

type ArticleType struct {
	Id       int        `orm:"pk"`
	TypeName string     `orm:"size(20)"`
	Article  []*Article `orm:"reverse(many)"` //与`orm:"rel()"`成对存在，与文章表结合表示一对多类型
}

/*一对多关系的会在数据库中建立外键，而多对多关系则会建立一张表*/

func init()  {
	orm.Debug = true//开发环境下建议开启调试
	orm.RegisterDataBase(
		"default",
		"mysql",
		"root:123456@tcp(127.0.0.1:3306)/articles?charset=utf8")

	orm.RegisterModel(new(User), new(Article), new(ArticleType))
	orm.RunSyncdb("default", false, true)
}

//用户查找
func SearchUser(name string, pwd string) (maps []orm.Params, err error) {
	o := orm.NewOrm()
	num, err := o.Raw("select * from user where user_name = ? and pwd = ?", name, pwd).Values(&maps)
	if err != nil && num == 0 {
		return nil, err
	}

	return maps, err
}

//用户注册
func RegisterUser(name string, pwd string) (ret int, err error) {
	search, err := SearchUser(name, pwd)
	if err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	if search == nil {
		user := User{}
		user.UserName = name
		user.Pwd = pwd

		_, err := o.Insert(&user)
		if err != nil {
			return 0, err
		}
		return 1, err //此时的err应为nil
	}

	return 0, err
}

//创建文章
func CreatArticle(artiname string, content string, timer time.Time, img string, artiType ArticleType) (ret int, err error) {
	o := orm.NewOrm()
	arti :=  Article{}

	arti.ArtiName = artiname
	arti.Content = content
	arti.PubTime = timer.Format("2006-01-02 15:04:05")
	arti.Img = img
	arti.ArticleType = &artiType

	_, err = o.Insert(&arti)
	if err != nil {
		return 0, err
	}
	return 1, err
}

//删除文章
func DeleteArticle(id int) (err error) {
	o := orm.NewOrm()
	arti := Article{Id:id}
	err = o.Read(&arti)
	if err != nil {
		return err
	}
	_, err = o.Delete(&arti)
	if err != nil {
		return err
	}
	return err
}

//文章分类获取
func GetArtiType() (artiType []ArticleType, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("ArticleType").All(&artiType)
	if err != nil {
		return nil, err
	}

	return artiType, err
}