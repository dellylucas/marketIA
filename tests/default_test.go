package test

import (
	"marketIA/models"
	_ "marketIA/routers"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestGetAll obtiene todos los usuarios //go test tests/default_test.go -v
func TestGetAll(t *testing.T) {
	ob := models.GetAllUser()
	for _, user := range ob {
		beego.Trace("testing", "TestGetAll", "Code %s", user)
	}
}

// TestGetOneHttp obtiene un usuario por id
func TestGetOneHttp(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/user/888", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGetOneHttp", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
		Convey("Status Code Should Be 200", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("The Result Should Not Be Empty", func() {
			So(w.Body.Len(), ShouldBeGreaterThan, 0)
		})
	})
}

// TestGetOne obtiene un usuario por id
func TestGetOne(t *testing.T) {
	var idUser = 888
	ob := models.GetUser(idUser)
	beego.Trace("testing", "TestGetOne", "Code %s", ob)

}

// TestDeleteUser elimina un usuario por id
func TestDeleteUser(t *testing.T) {
	var idUser = "123"
	ob := models.DeleteUser(idUser)
	beego.Trace("testing", "TestDeleteUser", "Code %s", ob)

}

// TestLogin valida login
func TestLogin(t *testing.T) {
	var userLog models.Usuarios

	userLog.CLAVE = "123"
	userLog.CORREO = "fagu@hotm.com"
	ob := models.ValidateLogin(userLog)
	beego.Trace("testing", "TestLogin", "Code %s", ob)

}

// TestGetStores obtene las tiendas
func TestGetStores(t *testing.T) {

	ob := models.GetAllStore()
	beego.Trace("testing", "TestGetStores", "Code %s", ob)

}

// TestInsertCat inserta categoria
func TestInsertCat(t *testing.T) {
	var cate models.Categorias

	cate.ID = 3
	cate.NOMBRE = "Seguridad"

	models.InsertCategories(&cate)
	beego.Trace("testing", "TestInsertCat", "Code %s", "ok")

}

// TestGetAllCat obtene las categorias
func TestGetAllCat(t *testing.T) {

	ob := models.GetAllCategories()
	for _, cat := range ob {
		beego.Trace("testing", "TestGetAllCat", "Code %s", cat)
	}

}

// TestGetAllPayMethods obtene todos los metodo de pago
func TestGetAllPayMethods(t *testing.T) {

	ob := models.GetAllPayMethods()
	for _, met := range ob {
		beego.Trace("testing", "TestGetAllPayMethods", "Code %s", met)
	}

}

// TestDeletePayMethods elimina un metodo de pago
func TestDeletePayMethods(t *testing.T) {
	var idUser = 123
	models.DeletePayMethods(idUser)
	beego.Trace("testing", "TestDeletePayMethods", "Code %s", "ok")

}
