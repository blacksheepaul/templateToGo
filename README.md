# templateToGo

Including a standard MVC structure, cfg, logger.

The database is SQLite. ( will be configurable soon )

# After clone

Replace module name in go.mod, then tidy.

linux bash

```bash
sed -i 's|github.com/blacksheepaul/templateToGo|your_module_name|g' go.mod
find . -name '*.go' -exec sed -i 's|github.com/blacksheepaul/templateToGo|your_module_name|g' {} \;
```

macOS
```bash
sed -i '' 's|github.com/blacksheepaul/templateToGo|your_module_name|g' go.mod
find . -name '*.go' -exec sed -i '' 's|github.com/blacksheepaul/templateToGo|your_module_name|g' {} \;
```

fish shell
```bash
sed -i -e 's|github.com/blacksheepaul/templateToGo|your_module_name|g' go.mod
find . -name '*.go' -exec sed -i -e 's|github.com/blacksheepaul/templateToGo|your_module_name|g' {} \;
```

then
```bash


go mod tidy
```
