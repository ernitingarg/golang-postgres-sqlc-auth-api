start:
	docker-compose up -d

stop:
	docker-compose down

remove:
	docker-compose down -v --rmi all

logs_db:
	docker-compose logs postgres

logs_sqlc:
	docker-compose logs sqlc

migrate_create:
	migrate create -ext sql -dir db/migrations init

migrate_up:
	migrate -path db/migrations -database "postgresql://admin:password123@localhost:5432/postgresdb?sslmode=disable" up

migrate_down:
	migrate -path db/migrations -database "postgresql://admin:password123@localhost:5432/postgresdb?sslmode=disable" down

.PHNOY: start stop remove logs_db logs_sqlc migrate_create migrate_up migrate_down