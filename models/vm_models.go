package models

import (
	//"github.com/astaxie/beego/orm"
	"time"
)

type Region struct {
	Id      int		`orm:"auto"`
	Name    string	`orm:"size(100)"`
	Created	time.Time `orm:"auto_now_add;type(datetime)"`
}

type Vm struct {
	Id			int			`orm:"auto"`
	Name		string		`orm:"size(100)"`
	Region_info	*Region 	`orm:"rel(fk);null"`
	Mem_size	int			`orm:"default(1000)"`
	Cpu_size	int			`orm:"default(1)"`
	Created		time.Time 	`orm:"auto_now_add;type(datetime)"`
}


type User struct {
	Id          int
	Name    string
	Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

type Profile struct {
	Id     int
	Age  int16
	User *User `orm:"reverse(one)"` // 设置反向关系（可选）
}



type Province struct {
	Id          	int64
	Name    		string		`orm:"size(100);unique"`
	ShortName    	string		`orm:"size(100)"`
	Created		time.Time 		`orm:"auto_now_add;type(datetime)"`
	Citys 		[]*City 		`orm:"reverse(many)"` // many relation
}


type City struct {
	Id          int64
	Name    	string		`orm:"size(100);unique"`
	Created		time.Time 	`orm:"auto_now_add;type(datetime)"`
	Province 	*Province 	`orm:"rel(fk)"` // fk relation
}


