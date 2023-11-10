mysqlinit:mysqlstop
	docker run --name task-scheduler-mysql-db -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -p 3306:3306 -d mysql
mymongoinit:mongostop
	docker run --name task-scheduler-mongo-db -d -p 27017:27017 mongo
	sleep 15
createmysqldb:
	 docker exec -it task-scheduler-mysql-db mysql -u root -e "CREATE DATABASE task_scheduler;"
dropmysqldb:
	 docker exec -it task-scheduler-mysql-db mysql -u root  -e "DROP DATABASE task_scheduler;"
mysqlstop:
	bash ./scripts/mysql_stop.sh
mongostop:
	bash ./scripts/mongo_stop.sh
check:
	bash ./scripts/check.sh
format:
    bash ./scripts/format.sh
db_mysql_prepare:mysqlinit
	docker cp task_scheduler.sql task-scheduler-mysql-db:task_scheduler.sql
	echo "Executing databases...wait for 15 seconds"
	sleep 15
	docker exec -it task-scheduler-mysql-db sh -c 'mysql -u root --socket=/var/run/mysqld/mysqld.sock < task_scheduler.sql'
db_mongo_prepare: mymongoinit
	docker exec -it task-scheduler-mongo-db mongosh --eval "use task_scheduler" --eval "db.createCollection('task')"
test: db_mongo_prepare db_mysql_prepare
	bash ./scripts/test.sh


.PHONY: test check