package transactions

import (
	model "github.com/RamadanRangkuti/multifinance-api/models"
	"github.com/RamadanRangkuti/multifinance-api/repository/transactions"
	"github.com/google/uuid"
)

type svc struct {
	transactionStore transactions.TransactionRepository
}

func NewTransactionSvc(transactionStore transactions.TransactionRepository) *svc {
	return &svc{transactionStore: transactionStore}
}

type TransactionSvc interface {
	CreateTransaction(req model.TransactionRequest) (*model.TransactionResponse, error)
	GetAllTransactions(userID uuid.UUID) ([]model.TransactionResponse, error)
}

func (s *svc) CreateTransaction(req model.TransactionRequest) (*model.TransactionResponse, error) {
	return s.transactionStore.CreateTransaction(req)
}

func (s *svc) GetAllTransactions(userID uuid.UUID) ([]model.TransactionResponse, error) {
	return s.transactionStore.GetAllTransactions(userID)
}
