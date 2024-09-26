STACK_NAME=application-api

start:
	docker-compose -p ${STACK_NAME} up -d --remove-orphans

stop:
	docker-compose -p ${STACK_NAME} stop

restart: stop start

clean:
	docker-compose -p ${STACK_NAME} down -v

logs:
	docker-compose -p ${STACK_NAME} logs -f

ps:
	docker-compose -p ${STACK_NAME} ps
