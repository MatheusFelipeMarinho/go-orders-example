package entity

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Reference       string
	OrderBump       bool `gorm:"column:has_order_bump"`
	IsUpsell        bool
	UpsellReference *string
	CheckoutData    CheckoutData `json:"checkout_data"`
	Customer        Customer
	CustomerAddress Address `json:"customer_address"`
	Payment         Payment
	Offers          []Offer
}

type CheckoutData struct {
	gorm.Model
	AffiliateHash *string
	CheckoutLink  string
	Fingerprint   string
	IP            string
	Device        string
	OrderID       uint // Chave estrangeira
}

type Customer struct {
	gorm.Model
	Name      string
	Document  string
	Phone     string
	Email     string
	BirthDate time.Time `gorm:"column:birth_at"`
	Genre     string
	OrderID   uint // Chave estrangeira
}

type Address struct {
	gorm.Model
	ZipCode      string
	Street       string
	Number       string
	Complement   *string
	Neighborhood string
	City         string
	State        string
	Country      string
	OrderID      uint // Chave estrangeira
	//CustomerID   uint // Chave estrangeira
}

type Payment struct {
	gorm.Model
	Amount          int
	Currency        string
	Parcels         int
	Discount        int
	Interest        int
	Freight         int
	PlatformRate    int
	Method          string
	CreditCardToken string
	OrderID         uint // Chave estrangeira
}

type Offer struct {
	gorm.Model
	Hash                string
	ProductHash         string
	ProductName         string
	ProductQty          int
	OfferPrice          int
	FeeBankSlipDiscount int
	FeeCreditDiscount   *int
	FeeDebitDiscount    int
	FeePixDiscount      int
	OrderID             uint // Chave estrangeira
}
