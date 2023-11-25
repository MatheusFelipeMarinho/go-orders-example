package routes

import (
	"github.com/gorilla/mux"
	"github/braip.com/internal/interfaces/controllers"
	"log"
	"net/http"
	"time"
)

// APIServer struct manages the HTTP server
type APIServer struct {
	router *mux.Router
}

// NewAPIServer creates a new instance of APIServer
func NewAPIServer() *APIServer {
	return &APIServer{
		router: mux.NewRouter(),
	}
}

// Start initiates the API server on the specified port
func (server *APIServer) Start(port string) {
	srv := &http.Server{
		Handler:      server.router,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting API server on port", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// ConfigureRoutes sets up the routes for the server
func (server *APIServer) ConfigureRoutes(oc *controllers.OrderController) {
	server.router.HandleFunc("/orders", oc.HandleCreateOrder).Methods("POST")
	server.router.HandleFunc("/orders", oc.HandleListOrders).Methods("GET")
	server.router.HandleFunc("/orders/{id}", oc.HandleGetOrder).Methods("GET")
	// Add other routes and their handlers here
}
