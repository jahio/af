module github.com/jahio/af

go 1.15

require (
	github.com/jahio/af/api v0.0.0
	github.com/jahio/af/models v0.0.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.8
)

replace github.com/jahio/af/models v0.0.0 => ./models

replace github.com/jahio/af/api v0.0.0 => ./api
