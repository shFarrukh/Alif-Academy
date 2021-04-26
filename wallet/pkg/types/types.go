package types

type Money int64


type PaymentCategory string


type PaymentStatus string


const (
	PaymentStatusOk         PaymentStatus = "OK"
	PaymentStatusFail       PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

// Payment представляет информацию о платеже.
type Payment struct {
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}

type Phone string

// Account представляет информацию о счёте пользователя.
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}

// Favorite представляет информацию об элементе "Избранное".
type Favorite struct {
	ID        string
	AccountID int64
	Amount    Money
	Name      string
	Category  PaymentCategory
}
