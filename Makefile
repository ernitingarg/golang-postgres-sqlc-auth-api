start_db:
	docker-compose up -d

stop_db:
	docker-compose down -v --rmi all

logs_db:
	docker-compose logs postgres

.PHNOY: start_db stop_db logs_db