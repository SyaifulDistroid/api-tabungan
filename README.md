# API-Tabungan
Simple API Tabungan 

# Config DB Postgres
```console
Setting config db with 'db_tabungan' as name of db
Exec Query 'Query Table Tabungan.sql'
```

# Build
```console
go build main.go
```

# API
```console
Import Tabungan.postman_collection_v2 For v2
Import Tabungan.postman_collection_v2.1 For v2.1
```

# Endpoint
```console
GET     - /v1/ping
GET     - /v1/health-check
POST    - /v1/tabungan/daftar
POST    - /v1/tabungan/tabung
POST    - /v1/tabungan/tarik
GET     - /v1/tabungan/saldo/:no_rekening
GET     - /v1/tabungan/mutasi/:no_rekenig
```

## Core library

Library | Usage
-- | --
fiber | Base framework
postgres | Database
logrus | Logger library

And others library are listed on `go.mod` file