package models

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex"`
	Password  string
	FirstName string
	LastName  string
	IsAdmin   bool
	IsActive  bool `gorm:"default:true"`
}

type Account struct {
	ID       int                 `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID   int                 `gorm:"not null" json:"userId"`
	BankID   int                 `gorm:"not null" json:"bankId"`
	Balance  float64             `gorm:"not null;default:1000" json:"balance"`
	IsActive bool                `gorm:"not null;default:true" json:"isActive"`
	Passbook []*TransactionEntry `gorm:"foreignKey:AccountID;references:ID" json:"passbook"`
}

type TransactionEntry struct {
	ID                     int       `gorm:"primaryKey;autoIncrement" json:"id"`
	AccountID              int       `gorm:"not null" json:"accountId"`
	CorrespondingAccountID int       `gorm:"not null" json:"correspondingAccountId"`
	CorrespondingBankID    int       `gorm:"not null" json:"correspondingBankId"`
	Type                   string    `gorm:"not null" json:"type"` // "credit" or "debit"
	Amount                 float64   `gorm:"not null" json:"amount"`
	BalanceAfter           float64   `gorm:"not null" json:"balanceAfter"`
	Timestamp              time.Time `gorm:"not null" json:"timestamp"`
}

type Bank struct {
	ID           int           `gorm:"primaryKey;autoIncrement" json:"id"`
	FullName     string        `gorm:"not null" json:"fullName"`
	Abbreviation string        `gorm:"size:5;not null" json:"abbreviation"`
	IsActive     bool          `gorm:"default:true" json:"isActive"`
	Accounts     []*Account    `gorm:"foreignKey:BankID;references:ID" json:"accounts"`
	LedgerData   []*LedgerData `gorm:"foreignKey:BankID;references:ID" json:"ledgerData"`
}

type LedgerData struct {
	ID                  int     `gorm:"primaryKey;autoIncrement" json:"id"`
	BankID              int     `gorm:"not null" json:"bankId"`
	CorrespondingBankID int     `gorm:"not null" json:"correspondingBankId"`
	Amount              float64 `gorm:"not null" json:"amount"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	if user.Username == "" {
		return tx.Error
	}
	return nil
}

var DB *gorm.DB

func InitDB() {
	dsn := "root:Forcepoint@2024@tcp(127.0.0.1:3306)/BankingApp?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	if err := DB.AutoMigrate(&Bank{}); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}
}

// Create the first Super Admin if no admins exist
func AddSuperAdmin() {
	var count int64
	DB.Model(&User{}).Where("is_admin = ?", true).Count(&count)

	if count == 0 {
		// Add a Super Admin
		superAdmin := User{
			Username:  "superadmin",
			Password:  "superpassword", // Ideally, hash this password
			FirstName: "Super",
			LastName:  "Admin",
			IsAdmin:   true,
			IsActive:  true,
		}
		DB.Create(&superAdmin)
		fmt.Println("Super Admin created successfully!")
	} else {
		fmt.Println("Super Admin already exists!")
	}
}

// ClearDatabase is useful for development/testing purposes to clear the data
func ClearDatabase() {
	DB.Exec("DROP DATABASE IF EXISTS GoBankingApp")
	DB.Exec("CREATE DATABASE GoBankingApp")
}

// ////// Claims struct that will be encoded into the JWT token.
type Claims struct {
	UserID   uint   `json:"userId"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.StandardClaims
}
