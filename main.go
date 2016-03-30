package main

import (
	_ "beego-dashboard/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	models "beego-dashboard/models"
	//"fmt"
	"fmt"
	"encoding/json"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:admin111@tcp(127.0.0.1:3306)/test_dashboard?charset=utf8")
	orm.RegisterModel(new(models.Region), new(models.Vm))
	orm.RegisterModel(new(models.User), new(models.Profile))
	orm.RegisterModel(new(models.Province), new(models.City))
	orm.RunSyncdb("default", false, true)
	orm.Debug = true
}

func main() {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	//reg1 := models.Region{}
	//reg1.Name = "cd"
	//ret, err := o.Insert(&reg1)
	//fmt.Println(ret, err)
	//
	//reg2 := models.Region{}
	//reg2.Name = "cd"
	//ret, err = o.Insert(&reg2)
	//fmt.Println(ret, err)

	//var reg1 models.Region
	//err := o.QueryTable("region").Filter("Name", "cd").One(&reg1)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//var vm1 models.Vm
	//vm1.Cpu_size = 10
	//vm1.Mem_size = 4000
	//vm1.Name = "fisrt_vm"
	//ret, err := o.Insert(&vm1)
	//fmt.Println(ret, err)
	//
	//var vm2 models.Vm
	//vm2.Cpu_size = 2
	//vm2.Mem_size = 2000
	//vm2.Name = "second_vm"
	//vm2.Region_info = &reg1
	//ret, err = o.Insert(&vm2)
	//fmt.Println(ret, err)

	//o := orm.NewOrm()
	//province := models.Province{Name: "sichuan",
	//							ShortName: "sc"}
	//
	//o.Insert(&province)
	//
	////city1 := models.City{Name: "chengdu"}
	//city_list := []models.City{
	//	{Name: "chengdu", Province: &province},
	//	{Name: "leshan", Province: &province},
	//}
	//o.InsertMulti(len(city_list), city_list)
	//orm.NewOrm().Insert(province)

	//var city_obj
	var maps []orm.Params
	num, err := o.QueryTable("city").Filter("name", "leshan").Limit(1).Values(&maps, "id", "name",
		"Province__Name", "Province__ShortName")
	fmt.Println(num)
	fmt.Println(err)
	dump_json, err := json.MarshalIndent(maps, "", " ")
	fmt.Println(err)
	fmt.Println(string(dump_json))

	num, err = o.QueryTable("city").Values(&maps, "id", "name", "Province__name", "Province__ShortName")
	fmt.Println(num)
	fmt.Println(err)
	dump_json, err = json.MarshalIndent(maps, "", " ")
	fmt.Println(err)
	fmt.Println(string(dump_json))




	beego.Run()
}

