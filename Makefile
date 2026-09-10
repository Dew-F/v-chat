include .env
export

DATABASE_URL=postgres://$(DB_USER):$(DB_PASSWORD)@localhost:${DB_PORT}/$(DB_NAME)?sslmode=disable

migrate-up:
	migrate \
	-path backend/migrations \
	-database "$(DATABASE_URL)" \
	up

migrate-down:
	migrate \
	-path backend/migrations \
	-database "$(DATABASE_URL)" \
	down 1

migrate-create:
	migrate create \
	-ext sql \
	-dir backend/migrations \
	-seq $(name)