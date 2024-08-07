package models

import (
	"be/db"
	"time"
)

type Role struct {
	ID    uint64 `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"type:varchar(20) not null;column:role_name"`
	Menus []Menu `gorm:"many2many:role_menu;"`
	Users []User
}

type User struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement"`
	UserName   string `gorm:"type:varchar(30) not null;unique"`
	Email      string `gorm:"type:varchar(50) not null"`
	FullName   string `gorm:"type:varchar(50) not null"`
	Password   string
	RoleId     uint64
	Budgets    []Budget
	Categories []Category
}

type Menu struct {
	ID     uint64  `gorm:"primaryKey;autoIncrement"`
	Title  string  `gorm:"type:varchar(20) not null"`
	Icon   *string `gorm:"type:varchar(50) not null"`
	MenuId *uint64
	Menus  []Menu
	Roles  []Role `gorm:"many2many:role_menu;"`
}

type Budget struct {
	ID     uint64 `gorm:"primaryKey;autoIncrement"`
	Name   string `gorm:"column:budget_name;type:varchar(100) not null"`
	Amount float64
	UserId uint64
}

type Category struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"column:category_name;type:varchar(100) not null;unique"`
	UserId   uint64
	Expenses []Expense
}

type Expense struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement"`
	Time       time.Time `gorm:"column:expense_time"`
	Amount     float64
	CategoryId uint64
}

func InitialTables() {
	conn := db.DbConnect()
	err := conn.AutoMigrate(&Role{}, &Menu{}, &User{}, &Budget{}, &Category{}, &Expense{})
	if err != nil {
		panic(err)
	}
}
