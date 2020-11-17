module github.com/Asuha-a/URLShortener/api/services/url

go 1.15

// replace github.com/Asuha-a/URLShortener/api => /Users/asuha/go/src/github.com/Asuha-a/URLShortener/api

require (
	github.com/Asuha-a/URLShortener/api v0.0.0-20201117104241-7e58ae45409a
	github.com/jackc/pgx/v4 v4.9.2 // indirect
	github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto v0.0.0-20201112155050-0c6587e931a9 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/grpc v1.33.2
	gorm.io/datatypes v1.0.0
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.6
)
