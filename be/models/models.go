package models

import (
	"be/db"
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID    uint64 `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"type:varchar(20) not null;column:role_name"`
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

type Budget struct {
	ID       uint64 `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"column:budget_name;type:varchar(100) not null"`
	Amount   float64
	UserId   uint64
	Expenses []Expense
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
	BudgetId   uint64
	CategoryId uint64
}

func InitialTables() {
	conn := db.DbConnect()
	err := conn.AutoMigrate(&Role{}, &User{}, &Budget{}, &Category{}, &Expense{})
	if err != nil {
		panic(err)
	}
	err = conn.Transaction(func(tx *gorm.DB) error {
		var roles []Role
		result := tx.Find(&roles)
		if result.Error != nil {
			return result.Error
		}
		if len(roles) == 0 {
			users := []User{
				{UserName: "admin", Email: "admin@admin.com", FullName: "admin", Password: "Admin@1234"},
			}
			adminRole := Role{Name: "Admin", Users: users}
			result := tx.Create(&adminRole)
			if result.Error != nil {
				return result.Error
			}
			userRole := Role{Name: "User"}
			result = tx.Create(&userRole)
			if result.Error != nil {
				return result.Error
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
