PORT_FORWARD=5432
DB_HOST=db
DB_NAME=reg_db
DB_PASS=password
DB_USER=postgres

DB_CONN="postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(PORT_FORWARD)/$(DB_NAME)?sslmode=disable"

# Поднять миграции
migrations-up:
	goose -dir migrations postgres $(DB_CONN) up

# Опустить миграции
migrations-down:
	goose -dir migrations postgres $(DB_CONN) down

# Пример как создавать новую миграцию
migrations-create-example:
	goose -dir migrations create add_some_column sql
