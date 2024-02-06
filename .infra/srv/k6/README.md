# запуск тестирования

curl --location 'http://social-net:8080/user/search?first_name=%D0%98&second_name=%D0%91%D1%80%D0%B0%D0%B3%D0%B8%D0%BD%D0%B0' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTI3OCwibmJmIjoxNzA0MTEwNDAwfQ.GP6RzCN4Z1bf_G3E_G2CsnA3QD4iq1YDLSvmVl7VySU'


docker-compose run k6 run /scripts/simple.js