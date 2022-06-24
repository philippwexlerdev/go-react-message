package main

import (
	"backend-test-template/app"
	"backend-test-template/app/common"
	"backend-test-template/app/config"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	config.LoadConfig()
	_ = os.Setenv("PORT", config.GetConfig().Port)
	_ = os.Setenv("DB_HOST", config.GetConfig().DBHost)
	_ = os.Setenv("DB_PORT", config.GetConfig().DBPort)
	_ = os.Setenv("DB_USER", config.GetConfig().DBUser)
	_ = os.Setenv("DB_PASS", config.GetConfig().DBPass)
	_ = os.Setenv("DB_NAME", config.GetConfig().DBName)

	r := mux.NewRouter()
	dsn := "host=" + config.GetConfig().DBHost
	dsn += " user=" + config.GetConfig().DBUser
	dsn += " password=" + config.GetConfig().DBPass
	dsn += " dbname=" + config.GetConfig().DBName
	dsn += " port=" + config.GetConfig().DBPort
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		common.ExitWithMsg("==== Fail to connect database ====")
	}
	application := app.NewApplication(r, db)
	log.Fatal(http.ListenAndServe(":"+config.GetConfig().Port, application.Router))
	fmt.Fprintln(os.Stderr, "==== API server is running ====")
}

//func TestQuickly(writer http.ResponseWriter, request *http.Request) {
//	responseWithJson(writer, http.StatusOK, map[string]string{"Message": "COmpletel"})
//}
