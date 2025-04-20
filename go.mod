module github.com/VenzeneCorp/orderService

go 1.24.1

require (
	github.com/VenzeneCorp/loginSignup v0.0.0-20250417184840-1fa6a9c83eb9
	github.com/VenzeneCorp/mealService v0.0.0-20250420101711-d9c1814a5d65
	gorm.io/gorm v1.25.12
)

replace github.com/VenzeneCorp/loginSignup => ../loginSignup

replace github.com/VenzeneCorp/mealService => ../mealService

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/sony/sonyflake v1.2.0 // indirect
	golang.org/x/text v0.17.0 // indirect
)
