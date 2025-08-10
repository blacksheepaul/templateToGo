# templateToGo

Including a standard MVC structure, cfg, logger.

The database is SQLite. ( will be configurable soon )

# After clone

Replace module name in go.mod, then tidy.

## linux bash

```bash
sed -i 's|github.com/blacksheepaul/templateToGo|your_module_name|g' go.mod
find . -name '*.go' -exec sed -i 's|github.com/blacksheepaul/templateToGo|your_module_name|g' {} \;
```

## macOS

```bash
sed -i '' 's|github.com/blacksheepaul/templateToGo|your_module_name|g' go.mod
find . -name '*.go' -exec sed -i '' 's|github.com/blacksheepaul/templateToGo|your_module_name|g' {} \;
```

## fish shell

```bash
sed -i -e 's|github.com/blacksheepaul/templateToGo|your_module_name|g' go.mod
find . -name '*.go' -exec sed -i -e 's|github.com/blacksheepaul/templateToGo|your_module_name|g' {} \;
```

## then common

```bash
go mod tidy
git add **/*.go
git add go.mod go.sum
```

# Migrate

for example:

create new migration

```bash
migrate -database "sqlite3://dev.db" create -seq -ext sql --dir model/migrations/ init_xxx_table
```

forward

```bash
migrate -database "sqlite3://dev.db" --path model/migrations/ up
```
