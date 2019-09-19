package db

import (
	"fmt"
	"marketIA/utils"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlCon = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s",
		utils.USER,
		utils.PASS,
		utils.HOST,
		utils.PORT,
		utils.NAMEDB,
		utils.CHARSET,
	)
)

func init() {
	orm.RegisterDriver(utils.ENGINE, orm.DRMySQL)
	orm.RegisterDataBase(utils.DBALIAS, utils.ENGINE, mysqlCon)
}

//GetSession - conexion de Base de Datos
func GetSession() orm.Ormer {

	/*force := true   // Drop table and re-create.
	verbose := true // Print log
	if err := orm.RunSyncdb(DBALIAS, force, verbose); err != nil {
		log.Println(err)
	}*/

	session := orm.NewOrm()
	session.Using(utils.NAMEDB)

	return session
}
