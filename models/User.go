package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id      int
	Name    string
	AddTime int64
	Status  int
	Mobile  string
	Avatar  string
}

func init() {
	orm.RegisterModel(new(User))
}

func UserInfo(id int) (User, error) {
	var (
		err error
	)
	o := orm.NewOrm()
	user := User{Id: id}
	err = o.Read(&user)
	return user, err
}
