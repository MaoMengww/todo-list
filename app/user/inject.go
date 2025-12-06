package user

import (
	"fmt"
	"log"
	"todo-list/app/user/controllers/rpc"
	"todo-list/app/user/infrastructure"
	"todo-list/app/user/usercase"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func UserInit() *rpc.UserServiceImpl {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.db"),
		viper.GetString("mysql.charset"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
	if err := db.AutoMigrate(&infrastructure.UserModel{}); err != nil {
		log.Fatalf("Error migrating database: %s", err)
	}
	userRepo := infrastructure.NewMysqlUserRepository(db)
	userHaddler := usercase.NewUserUseCase(userRepo)
	return rpc.NewUserServiceImpl(userHaddler)
}