# Server

v0.0.1, Sep 29, 2025

Moto-Alert server

1. Install app dependencies

```
go mod download
```

2. Create .env file

```
JWT_SECRET="123123123abcabc"
DB_PASSWORD="123123123abcabc"
EXPO_ACCESS_TOKEN="123123123abcabc"
BREVO_ACCESS_TOKEN="123123123abcabc"

GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://user:password@localhost:5432/moto_alerts_local
GOOSE_MIGRATION_DIR=./migrations
```

3. Run migrations

```
make migration-run
```

4. Start app

```
make run
```
