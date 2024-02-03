# hl-sn
Заготовка для социальной сети

### Подготовка.
- установить docker (https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-20-04-ru)
- установить docker-compose (https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-compose-on-ubuntu-22-04)
- импортировать коллекцию "корень проекта/tools/collection/hl-sn.postman_collection.json"
- добавить переменную окружения для postman URL_HL_SRV. Пример: URL_HL_SRV = http://riva.local:8080
- установить make, если не установлен(sudo apt-get install make)

### Запуск проекта.
 - перейти в папку "корень проекта/.infra/dev/"
 - запустить make up (запустит контейнер для проекта)
 - запустить make srv-psql (запустит DB для проекта, следует посмотреть логи и убедиться, что все прошло успешно)
 - запустить make migr (сделает миграцию)
 - запустить go-run (запустить сам проект на GO)


 После запуска проекта, контейнер готов слушать :8080 сервис, пример запросов см. коллекцию "корень проекта/tools/collection/hl-sn.postman_collection.json"

