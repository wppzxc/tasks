package cmd

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wppzxc/tasks/src/database"
	"github.com/wppzxc/tasks/src/pkg/types"
	"github.com/wppzxc/tasks/src/server"
)

func Run(cfg *types.Config) error {
	db, err := gorm.Open("mysql", cfg.MysqlUrl)
	if err != nil {
		return err
	}
	defer db.Close()
	database.AutoMigrate(db)
	database.DB = db
	e := server.NewEcho()
	url := cfg.BindIP + ":" + cfg.BindPort
	if err := e.Start(url); err != nil {
		return err
	}
	return nil
}
