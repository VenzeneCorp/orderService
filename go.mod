module github.com/VenzeneCorp/orderService

go 1.24.1

require (
	github.com/VenzeneCorp/loginSignup v0.0.0-20250417184840-1fa6a9c83eb9
	github.com/VenzeneCorp/mealService v0.0.0-20250420101711-d9c1814a5d65
	github.com/golang-jwt/jwt/v5 v5.2.2
	github.com/gorilla/mux v1.8.1
	github.com/joho/godotenv v1.5.1
	github.com/sony/sonyflake v1.2.0
	gorm.io/driver/postgres v1.5.11
	gorm.io/gorm v1.25.12
)

replace github.com/VenzeneCorp/loginSignup => ../loginSignup

replace github.com/VenzeneCorp/mealService => ../mealService

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.5.5 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/razorpay/razorpay-go v1.3.3 // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/text v0.17.0 // indirect
)
