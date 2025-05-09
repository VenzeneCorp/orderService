package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	// loginsignupclient "github.com/VenzeneCorp/orderService/clients/loginSignupClient"
	// mealserviceclient "github.com/VenzeneCorp/orderService/clients/mealServiceClient"
	"github.com/VenzeneCorp/orderService/middlewares"
	"github.com/VenzeneCorp/orderService/models"
	orderhistory "github.com/VenzeneCorp/orderService/repo/orderHistory"
	postgresdb "github.com/VenzeneCorp/orderService/repo/orderHistory/postgresDB"
	vendoranalytics "github.com/VenzeneCorp/orderService/repo/vendorAnalytics"
	vendorpostgresdb "github.com/VenzeneCorp/orderService/repo/vendorAnalytics/vendorPostgresDB"
	ordermgmt "github.com/VenzeneCorp/orderService/service/orderMgmt"
	ordertransport "github.com/VenzeneCorp/orderService/service/orderMgmt/transport"
	vendormgmt "github.com/VenzeneCorp/orderService/service/vendorMgmt"
	vendortransport "github.com/VenzeneCorp/orderService/service/vendorMgmt/transport"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Initialize database connection
	db, err := setupDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// // Create external service clients
	// mealSvcClient := mealserviceclient.NewClient(getEnv("MEAL_SERVICE_URL", "http://localhost:8081"))
	// loginSignupClient := loginsignupclient.NewClient(getEnv("LOGIN_SIGNUP_SERVICE_URL", "http://localhost:8082"))

	// Initialize repositories
	orderRepo := setupOrderRepository(db)
	vendorRepo := setupVendorRepository(db)

	// Initialize services
	orderService := ordermgmt.NewService(orderRepo)
	vendorService := vendormgmt.NewService(vendorRepo)

	// Initialize handlers
	orderHandler := ordertransport.NewHandler(orderService)
	vendorHandler := vendortransport.NewHandler(vendorService)

	// Create and configure router
	router := mux.NewRouter()

	// Register API routes
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	orderHandler.RegisterRoutes(apiRouter)
	vendorHandler.RegisterRoutes(apiRouter)

	// Middleware for CORS
	router.Use(middlewares.CorsMiddleware)

	// Create HTTP server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", getEnv("PORT", "8080")),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on port %s", getEnv("PORT", "8080"))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server gracefully stopped")
}

func setupDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "order_service"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_TIMEZONE", "UTC"),
	)

	// Configure GORM logger
	logLevel := logger.Info
	if getEnv("ENV", "development") == "production" {
		logLevel = logger.Error
	}

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}

	// Connect to database
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}

	// Migrate the schema (if needed)
	err = db.AutoMigrate(&models.Orders{}, &models.ItemOrdered{}, &models.Subscription{})

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	maxIdleConns, _ := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNS", "10"))
	maxOpenConns, _ := strconv.Atoi(getEnv("DB_MAX_OPEN_CONNS", "100"))
	connMaxLifetime, _ := strconv.Atoi(getEnv("DB_CONN_MAX_LIFETIME", "3600"))

	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)

	return db, nil
}

func setupOrderRepository(db *gorm.DB) orderhistory.Repository {
	postgresRepo := postgresdb.NewRepository(db)
	return orderhistory.NewUserHistoryRepo(&postgresRepo)
}

func setupVendorRepository(db *gorm.DB) vendoranalytics.Repository {
	vendorPostgresRepo := vendorpostgresdb.NewRepository(db)
	return vendoranalytics.NewVendorAnalyticsRepo(&vendorPostgresRepo)
}

// getEnv retrieves environment variables with fallback to default values
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
