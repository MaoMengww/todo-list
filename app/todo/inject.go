package todo

import (
	"fmt"
	"log"
	"todo-list/app/todo/controllers/rpc"
	"todo-list/app/todo/infrastructure"
	"todo-list/app/todo/usecase"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TodoInit() *rpc.TodoServiceImpl {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.db2"),
		viper.GetString("mysql.charset"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
	if err := db.AutoMigrate(&infrastructure.TodoModel{}); err != nil {
		log.Fatalf("Error migrating database: %s", err)
	}
	todoRepo := infrastructure.NewMysqlTodoRepository(db)
	todoHandler := usecase.NewUsecase(todoRepo)
	return rpc.NewTodoServiceImpl(todoHandler)
}