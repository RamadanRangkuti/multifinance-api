package transactions

import (
	"encoding/json"
	"net/http"

	"github.com/RamadanRangkuti/multifinance-api/helper"
	model "github.com/RamadanRangkuti/multifinance-api/models"
	"github.com/RamadanRangkuti/multifinance-api/usecase/transactions"
	"github.com/RamadanRangkuti/multifinance-api/util/middleware"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type Handler struct {
	transactionSvc transactions.TransactionSvc
	validator      *validator.Validate
}

func NewHandler(transactionSvc transactions.TransactionSvc, validator *validator.Validate) *Handler {
	return &Handler{
		transactionSvc: transactionSvc,
		validator:      validator,
	}
}

func (h *Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())

	// Konversi userID dari string ke uuid.UUID
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		helper.HandleResponse(w, http.StatusUnauthorized, "Invalid user ID", nil)
		return
	}

	var req model.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		helper.HandleResponse(w, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	req.UserID = parsedUserID

	if err := h.validator.Struct(req); err != nil {
		errorMessages := helper.FormatValidationError(err)
		helper.HandleResponse(w, http.StatusBadRequest, "Validation Error", errorMessages)
		return
	}

	transaction, err := h.transactionSvc.CreateTransaction(req)
	if err != nil {
		helper.HandleResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.HandleResponse(w, http.StatusCreated, "Transaction created successfully", transaction)
}

func (h *Handler) GetTransactions(w http.ResponseWriter, r *http.Request) {
	// Ambil user ID dari token yang tersimpan di context
	userID := middleware.GetUserID(r.Context())

	// Konversi user ID dari string ke uuid.UUID
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		helper.HandleResponse(w, http.StatusUnauthorized, "Invalid user ID", nil)
		return
	}

	// Ambil transaksi berdasarkan user ID
	transactions, err := h.transactionSvc.GetAllTransactions(parsedUserID)
	if err != nil {
		helper.HandleResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.HandleResponse(w, http.StatusOK, "Transactions retrieved successfully", transactions)
}
