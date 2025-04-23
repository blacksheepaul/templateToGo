# templateToGo

Including a standard MVC structure, cfg, logger.

The database is SQLite. ( will be configurable soon )

# After clone

Replace module name in go.mod, then tidy.

```bash
sed -i -e 's|github.com/blacksheepaul/templateToGo|"your_module_name"|g' go.mod
go mod tidy
```
