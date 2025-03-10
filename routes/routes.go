package routes

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/RamadanRangkuti/multifinance-api/config"
	"github.com/RamadanRangkuti/multifinance-api/handlers/transactions"
	user "github.com/RamadanRangkuti/multifinance-api/handlers/users"
	"github.com/RamadanRangkuti/multifinance-api/util/middleware"
	"github.com/spf13/viper"
)

type Routes struct {
	Router       *http.ServeMux
	User         *user.Handler
	Transactions *transactions.Handler
}

func URLRewriter(baseURLPath string, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimPrefix(r.URL.Path, baseURLPath)

		next.ServeHTTP(w, r)
	}
}

func (r *Routes) SetupBaseURL() {
	baseURL := viper.GetString("BASE_URL_PATH")
	if baseURL != "" && baseURL != "/" {
		r.Router.HandleFunc(baseURL+"/", URLRewriter(baseURL, r.Router))
	}
}

func (r *Routes) SetupRouter() {
	r.Router = http.NewServeMux()
	r.SetupBaseURL()
	r.userRoutes()
	r.transactionRoutes()
}

func (r *Routes) userRoutes() {
	r.Router.Handle("POST /signup", middleware.ApplyMiddleware(
		http.HandlerFunc(r.User.SignUp),
		middleware.EnabledCors,
		middleware.LoggerMiddleware(),
	))

	r.Router.Handle("POST /signin", middleware.ApplyMiddleware(
		http.HandlerFunc(r.User.SignIn),
		middleware.EnabledCors,
		middleware.LoggerMiddleware(),
	))
}

// func (r *Routes) userRoutes() {
// 	r.Router.HandleFunc("POST /signup", middleware.ApplyMiddleware(r.User.SignUp, middleware.EnabledCors, middleware.LoggerMiddleware()))
// 	r.Router.Handle("POST /signin", middleware.ApplyMiddleware(r.User.SignIn, middleware.EnabledCors, middleware.LoggerMiddleware()))
// }

//	func (r *Routes) transactionRoutes() {
//		r.Router.HandleFunc("POST /transactions", middleware.ApplyMiddleware(r.Transactions.CreateTransaction, middleware.EnabledCors, middleware.LoggerMiddleware()))
//		r.Router.HandleFunc("GET /transactions", middleware.ApplyMiddleware(r.Transactions.GetTransactions, middleware.EnabledCors, middleware.LoggerMiddleware()))
//	}
func (r *Routes) transactionRoutes() {
	r.Router.Handle("POST /transactions", middleware.ApplyMiddleware(
		middleware.Authentication(http.HandlerFunc(r.Transactions.CreateTransaction)),
		middleware.EnabledCors,
		middleware.LoggerMiddleware(),
	))

	r.Router.Handle("GET /transactions", middleware.ApplyMiddleware(
		middleware.Authentication(http.HandlerFunc(r.Transactions.GetTransactions)),
		middleware.EnabledCors,
		middleware.LoggerMiddleware(),
	))
}

func (r *Routes) Run(port string) {
	r.SetupRouter()

	log.Printf("[Running-Success] clients on localhost on port :%s", port)
	srv := &http.Server{
		Handler:      r.Router,
		Addr:         "localhost:" + port,
		WriteTimeout: config.WriteTimeout() * time.Second,
		ReadTimeout:  config.ReadTimeout() * time.Second,
	}

	log.Panic(srv.ListenAndServe())
}
