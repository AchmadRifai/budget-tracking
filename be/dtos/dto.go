package dtos

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegRequest struct {
	LoginRequest
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type NewBudget struct {
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
}

type NewCategory struct {
	Name string `json:"name"`
}

type NewExpenses struct {
	CategoryId uint64  `json:"category_id"`
	BudgetId   uint64  `json:"budget_id"`
	Amount     float64 `json:"amount"`
	Time       int64   `json:"time"`
}

type ExpensesChart struct {
	Category string  `json:"category"`
	Budget   string  `json:"budget"`
	Month    string  `json:"month"`
	Amount   float64 `json:"amount"`
}
