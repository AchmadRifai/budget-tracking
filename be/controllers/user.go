package controllers

import (
	"be/db"
	"be/dtos"
	errorhandlers "be/errorHandlers"
	"be/models"
	myutils "be/myUtils"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"

	arrayutils "github.com/AchmadRifai/array-utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetChart(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	user := myutils.GetUser(r)
	var expenses []models.Expense
	conn := db.DbConnect()
	catQuery := conn.Model(&models.Category{}).Where("user_id=?", user.ID).Select("id")
	budgetQuery := conn.Model(&models.Budget{}).Where("user_id=?", user.ID).Select("id")
	result := conn.Where("category_id IN(?) AND budget_id IN(?)", catQuery, budgetQuery).Order("expense_time desc").Find(&expenses)
	if result.Error != nil {
		panic(result.Error)
	}
	categories, budgets := user.Categories, user.Budgets
	charts := make(map[string]map[string]map[string]float64)
	for _, expense := range expenses {
		budget := arrayutils.Filter(budgets, func(v1 models.Budget, _ int) bool { return v1.ID == expense.BudgetId })[0]
		if !myutils.MapKeyExists(charts, budget.Name) {
			charts[budget.Name] = make(map[string]map[string]float64)
		}
		category := arrayutils.Filter(categories, func(v1 models.Category, _ int) bool { return v1.ID == expense.CategoryId })[0]
		if !myutils.MapKeyExists(charts[category.Name], category.Name) {
			charts[budget.Name][category.Name] = make(map[string]float64)
		}
		month := expense.Time.Format("Jan 2006")
		if !myutils.MapKeyExists(charts[budget.Name][category.Name], month) {
			charts[budget.Name][category.Name][month] = 0
		}
		charts[budget.Name][category.Name][month] = charts[budget.Name][category.Name][month] + expense.Amount
	}
	myutils.SendJson(w, map[string]interface{}{"message": "Success", "charts": charts}, 200)
}

func AddBudget(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var req dtos.NewBudget
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}
	if req.Name == "" {
		panic(errors.New("name is required"))
	}
	if req.Amount <= 0 {
		panic(errors.New("amount is required"))
	}
	user := myutils.GetUser(r)
	budgets := user.Budgets
	if arrayutils.AnyOf(budgets, func(v models.Budget, _ int) bool {
		return v.Name == req.Name
	}) {
		panic(errors.New("duplicate name budget"))
	}
	err = db.DbConnect().Transaction(func(tx *gorm.DB) error {
		budget := models.Budget{Name: req.Name, Amount: req.Amount, UserId: user.ID}
		result := tx.Create(&budget)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	myutils.SendJson(w, map[string]interface{}{"message": "Success"}, 200)
}

func AllBudget(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	user := myutils.GetUser(r)
	data := arrayutils.Map(user.Budgets, func(v models.Budget, _ int) map[string]interface{} {
		return map[string]interface{}{"id": v.ID, "amount": v.Amount, "name": v.Name}
	})
	myutils.SendJson(w, map[string]interface{}{"message": "Success", "data": data}, 200)
}

func EditBudget(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var req dtos.NewBudget
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}
	if req.Name == "" {
		panic(errors.New("name is required"))
	}
	if req.Amount <= 0 {
		panic(errors.New("amount is required"))
	}
	vars := mux.Vars(r)
	user := myutils.GetUser(r)
	budgets := user.Budgets
	if !arrayutils.AnyOf(budgets, func(v models.Budget, _ int) bool {
		return vars["id"] == strconv.FormatUint(v.ID, 10)
	}) {
		panic(errors.New("id not found"))
	}
	if arrayutils.AnyOf(budgets, func(v models.Budget, _ int) bool {
		return req.Name == v.Name && vars["id"] != strconv.FormatUint(v.ID, 10)
	}) {
		panic(errors.New("duplicate name budget"))
	}
	err = db.DbConnect().Transaction(func(tx *gorm.DB) error {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			return err
		}
		var budget models.Budget
		result := tx.Where("id=?", id).First(&budget)
		if result.Error != nil {
			return result.Error
		}
		budget.Amount = req.Amount
		budget.Name = req.Name
		result = tx.Save(&budget)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	myutils.SendJson(w, map[string]interface{}{"message": "Success"}, 200)
}

func DelBudget(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	vars := mux.Vars(r)
	user := myutils.GetUser(r)
	err := db.DbConnect().Transaction(func(tx *gorm.DB) error {
		budgets := user.Budgets
		if !arrayutils.AnyOf(budgets, func(v models.Budget, _ int) bool {
			return vars["id"] == strconv.FormatUint(v.ID, 10)
		}) {
			return errors.New("id not found")
		}
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			return err
		}
		var budget models.Budget
		result := tx.Where("id=?", id).First(&budget)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Delete(&budget)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	myutils.SendJson(w, map[string]interface{}{"message": "Success"}, 200)
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var req dtos.NewCategory
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}
	if req.Name == "" {
		panic(errors.New("name is required"))
	}
	user := myutils.GetUser(r)
	categories := user.Categories
	if arrayutils.AnyOf(categories, func(v models.Category, _ int) bool {
		return v.Name == req.Name
	}) {
		panic(errors.New("duplicate name budget"))
	}
	err = db.DbConnect().Transaction(func(tx *gorm.DB) error {
		category := models.Category{Name: req.Name, UserId: user.ID}
		result := tx.Create(&category)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	myutils.SendJson(w, map[string]interface{}{"message": "Success"}, 200)
}

func AllCategory(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	user := myutils.GetUser(r)
	data := arrayutils.Map(user.Categories, func(v models.Category, _ int) map[string]interface{} {
		return map[string]interface{}{"id": v.ID, "name": v.Name}
	})
	myutils.SendJson(w, map[string]interface{}{"message": "Success", "data": data}, 200)
}

func EditCategory(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var req dtos.NewCategory
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}
	if req.Name == "" {
		panic(errors.New("name is required"))
	}
	vars := mux.Vars(r)
	user := myutils.GetUser(r)
	categories := user.Categories
	if !arrayutils.AnyOf(categories, func(v models.Category, _ int) bool {
		return vars["id"] == strconv.FormatUint(v.ID, 10)
	}) {
		panic(errors.New("id not found"))
	}
	if arrayutils.AnyOf(categories, func(v models.Category, _ int) bool {
		return req.Name == v.Name && vars["id"] != strconv.FormatUint(v.ID, 10)
	}) {
		panic(errors.New("duplicate name category"))
	}
	err = db.DbConnect().Transaction(func(tx *gorm.DB) error {
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			return err
		}
		var category models.Category
		result := tx.Where("id=?", id).First(&category)
		if result.Error != nil {
			return result.Error
		}
		category.Name = req.Name
		result = tx.Save(&category)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	myutils.SendJson(w, map[string]interface{}{"message": "Success"}, 200)
}

func DelCategory(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	vars := mux.Vars(r)
	user := myutils.GetUser(r)
	err := db.DbConnect().Transaction(func(tx *gorm.DB) error {
		catgeories := user.Categories
		if !arrayutils.AnyOf(catgeories, func(v models.Category, _ int) bool {
			return vars["id"] == strconv.FormatUint(v.ID, 10)
		}) {
			return errors.New("id not found")
		}
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			return err
		}
		var category models.Category
		result := tx.Where("id=?", id).First(&category)
		if result.Error != nil {
			return result.Error
		}
		result = tx.Delete(&category)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	myutils.SendJson(w, map[string]interface{}{"message": "Success"}, 200)
}

func AddExpenses(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	user := myutils.GetUser(r)
	budgets, categories := user.Budgets, user.Categories
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var req dtos.NewExpenses
	if err = json.Unmarshal(body, &req); err != nil {
		panic(err)
	}
	err = db.DbConnect().Transaction(func(tx *gorm.DB) error {
		if !arrayutils.AnyOf(budgets, func(v models.Budget, _ int) bool { return req.BudgetId == v.ID }) {
			return errors.New("budget not found")
		}
		if !arrayutils.AnyOf(categories, func(v models.Category, _ int) bool { return req.CategoryId == v.ID }) {
			return errors.New("category not found")
		}
		if req.Amount <= 0 {
			return errors.New("amount is required")
		}
		times := time.Unix(req.Time, 0)
		expense := models.Expense{Time: times, Amount: req.Amount, BudgetId: req.BudgetId, CategoryId: req.CategoryId}
		result := tx.Create(&expense)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	myutils.SendJson(w, map[string]interface{}{"message": "Success"}, 200)
}

func AllExpenses(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	user := myutils.GetUser(r)
	var expenses []models.Expense
	conn := db.DbConnect()
	catQuery := conn.Model(&models.Category{}).Where("user_id=?", user.ID).Select("id")
	budgetQuery := conn.Model(&models.Budget{}).Where("user_id=?", user.ID).Select("id")
	result := conn.Where("category_id IN(?) AND budget_id IN(?)", catQuery, budgetQuery).Order("expense_time desc").Find(&expenses)
	if result.Error != nil {
		panic(result.Error)
	}
	categories := user.Categories
	budgets := user.Budgets
	data := arrayutils.Map(expenses, func(v models.Expense, _ int) map[string]interface{} {
		category := arrayutils.Filter(categories, func(v2 models.Category, _ int) bool { return v.CategoryId == v2.ID })[0]
		budget := arrayutils.Filter(budgets, func(v2 models.Budget, _ int) bool { return v.BudgetId == v2.ID })[0]
		return map[string]interface{}{"id": v.ID, "time": v.Time.Unix(), "category": category.Name, "budget": budget.Name, "amount": v.Amount}
	})
	myutils.SendJson(w, map[string]interface{}{"message": "Success", "data": data}, 200)
}

func EditExpenses(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	vars := mux.Vars(r)
	user := myutils.GetUser(r)
	budgets, categories := user.Budgets, user.Categories
	body, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var req dtos.NewExpenses
	if err = json.Unmarshal(body, &req); err != nil {
		panic(err)
	}
	err = db.DbConnect().Transaction(func(tx *gorm.DB) error {
		var expense models.Expense
		result := tx.Where("id=?", vars["id"]).First(&expense)
		if result.Error != nil {
			return result.Error
		}
		if !arrayutils.AllOf(budgets, func(v models.Budget, _ int) bool { return v.ID == expense.BudgetId }) {
			return errors.New("budget not found")
		}
		if !arrayutils.AllOf(categories, func(v models.Category, _ int) bool { return v.ID == expense.CategoryId }) {
			return errors.New("category not found")
		}
		times := time.Unix(req.Time, 0)
		if !arrayutils.AnyOf(budgets, func(v models.Budget, _ int) bool { return req.BudgetId == v.ID }) {
			return errors.New("budget not found")
		}
		if !arrayutils.AnyOf(categories, func(v models.Category, _ int) bool { return req.CategoryId == v.ID }) {
			return errors.New("category not found")
		}
		if req.Amount <= 0 {
			return errors.New("amount is required")
		}
		expense.Amount = req.Amount
		expense.BudgetId = req.BudgetId
		expense.CategoryId = req.CategoryId
		expense.Time = times
		result = tx.Save(&expense)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	myutils.SendJson(w, map[string]interface{}{"message": "Success"}, 200)
}

func DelExpenses(w http.ResponseWriter, r *http.Request) {
	defer errorhandlers.NormalErrorRest(w, r)
	vars, user := mux.Vars(r), myutils.GetUser(r)
	budgets, categories := user.Budgets, user.Categories
	err := db.DbConnect().Transaction(func(tx *gorm.DB) error {
		var expense models.Expense
		result := tx.Where("id=?", vars["id"]).First(&expense)
		if result.Error != nil {
			return result.Error
		}
		if !arrayutils.AllOf(budgets, func(v models.Budget, _ int) bool { return v.ID == expense.BudgetId }) {
			return errors.New("budget not found")
		}
		if !arrayutils.AllOf(categories, func(v models.Category, _ int) bool { return v.ID == expense.CategoryId }) {
			return errors.New("category not found")
		}
		result = tx.Delete(&expense)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	myutils.SendJson(w, map[string]interface{}{"message": "Success"}, 200)
}
