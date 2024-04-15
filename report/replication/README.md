
# Настройка асинхронной репликации.
## Проводим нагрузочный тест master`a до переключения 2 запросов  на slave

Запускаем систему мониторинга (.infra/dev/Makefile): m srv-monitoring; n srv-grafana
Запускаем нагрузочное тестирование (.infra/srv/k6/docker-compose.yml): docker-compose run k6 run /scripts/hl_test_async_replica.js
![](hl_test_before_replice.png)

По данным графика видим нагрузку на master.


## Создаем пользователя, под которым будем подключаться со стороны вторичного сервера:
docker exec -it postgres_master su - postgres -c psql create role replicator with login replication password 'pass'; exit

## Настраиваем параметры репликации для мастера:
Открываем:
github.com/avp365/hl-sn/.infra/srv/psql/master/postgresql.conf

ssl = off
wal_level = replica
max_wal_senders = 3 #максимальное количество одновременно работающих процессов передачи

hot_standby = on
hot_standby_feedback = on

## Настраиваем параметры доступа для мастера:
Разрешаем подключаться к нашему серверу из любой подсети и без пароля (чтобы упростить настройки)
host    replication     all             0.0.0.0/0           trust

## Сделаем бэкап для реплик:
docker exec -it postgres_master bash
mkdir /pgslave
pg_basebackup -h postgres-16-1-master -D /pgslave -U replicator -v -P --wal-method=stream
exit

## Настраиваем параметры репликации для слейва s1:
Переходим в папку 1-го slave: cd .infra/srv/psql/s1

Копируем pgslave в папку postgres: docker cp postgres_master:/pgslave postgres/

Создадим файл, чтобы реплика узнала, что она реплика: touch ./postgres/standby.signal

Меняем postgresql.conf на реплике s1: primary_conninfo = 'host=postgres-16-1-master port=5432 user=replicator password=pass application_name=s1'

Запускаем: cd /.infra/dev/Makefile; make srv-psql-s1

## Настраиваем параметры репликации для слейва s2:
Переходим в папку 2-го slave: cd .infra/srv/psql/s2

Копируем pgslave в папку postgres: docker cp postgres_master:/pgslave postgres/

Создадим файл, чтобы реплика узнала, что она реплика: touch ./postgres/standby.signal

Меняем postgresql.conf на реплике s2: primary_conninfo = 'host=postgres-16-1-master port=5432 user=replicator password=pass application_name=s2'

Запускаем: cd /.infra/dev/Makefile; make srv-psql-s2

## Убеждаемся что обе реплики работают в асинхронном режиме на postgres_master
docker exec -it postgres_master su - postgres -c psql
select application_name, sync_state from pg_stat_replication;
exit;

![](async_replica.png)
## Проводим нагрузочный тест master`a после переключения 2 запросов  на slave s1. 
Добавляем в .env (.infra/dev/.env) POSTGRESS_SLAVE_1_URL=postgres://postgres:123456@postgres-16-1-s1:5432/socnet
<br>Меняем запросы  /user/get/{id} и /user/search на s1 /internal/repositories/user.go см. r.DBPostrS1.....
<br>Запускаем систему мониторинга (.infra/dev/Makefile): m srv-monitoring; n srv-grafana
<br>Запускаем нагрузочное тестирование (.infra/srv/k6/docker-compose.yml): docker-compose run k6 run /scripts/hl_test_async_replica.js
![](hl_test_after_replice.png) в 20.15

По данным графика не видим нагрузку на master.
![](result_test.png)

# Настройка синхронной репликации.

## Настраиваем параметры репликации для слейва:
docker exec -it postgres_container_s1 bash
rm -R /data/postgres/*
su - postgres -c "pg_basebackup --host=postgres-16-1 --username=repluser --pgdata=/data/postgres --wal-method=stream --write-recovery-conf"

docker exec -it postgres_container_s2 bash
rm -R /data/postgres/*
su - postgres -c "pg_basebackup --host=postgres-16-1 --username=repluser --pgdata=/data/postgres --wal-method=stream --write-recovery-conf"

## Проверяем:
Master
docker exec -it postgres_container su - postgres -c "psql -c 'select * from pg_stat_replication;'"
Slave 1
docker exec -it postgres_container_s1 su - postgres -c "psql -c 'select * from pg_stat_replication;'"
Slave 2
docker exec -it postgres_container_s2 su - postgres -c "psql -c 'select * from pg_stat_replication;'"