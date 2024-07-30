
# Настройка кеширования.
## Создаем набор данных с помощью генераторов.
### Генерим список друзей с помощью скрипта. Файл сохранится в формате .csv по пути /home/dev/go/src/github.com/avp365/hl-sn/.infra/srv/psql/pgadmin/data , где будет залит в БД
[/home/dev/go/src/github.com/avp365/hl-sn/tools/generator/friend/main.go](https://github.com/avp365/hl-sn/blob/cache_queue/tools/generator/friend/main.go)
### Генерим посты друзей с помощью скрипта, на основе файла friends. Файл сохранится в формате .csv по пути /home/dev/go/src/github.com/avp365/hl-sn/.infra/srv/psql/pgadmin/data , где будет залит в БД
[/home/dev/go/src/github.com/avp365/hl-sn/tools/generator/posts/main.go](https://github.com/avp365/hl-sn/blob/cache_queue/tools/generator/posts/main.go)

Миграция для таблицы:
https://github.com/avp365/hl-sn/blob/cache_queue/internal/migrations/20240629232858_table_friends.sql
https://github.com/avp365/hl-sn/blob/cache_queue/internal/migrations/20240630302858_table_posts.sql

Слудет учесть что механизм добавлений и хранение друзей, а так же подбор является очень простейшим. Например. 
Если рассматривать таблицу friends и столбцы id_user_1 и id_user_2, то мы все время генерим нашего пользователя в id_user_1, а его друзей в id_user_2. А посты достаем join`amи. Хотя по настоящему, пользователь может генерится как в id_user_1, так id_user_2. следовательно запрос будет более сложным. Так же мы не обращаем внимание на оптимизацию на уровне DB. Нас интересует механизм кешей.


## Выбор технологии кеширования и очередей.
Для кеширования и очередей будем использовать Redis. Redis был выбран в силу простоты использования.  Есть и другие варианты.

## Описание кеширования.
Кешировать будем на уровне роутов.
https://github.com/avp365/hl-sn/blob/3c4bf8b7aaa9d0bb51a9c80cf917bdf1296bbac9/internal/routers/post.go#L128
Время кешированием выберем 48 часов.

Так же при изменении постов будем добавлять в очередь пользователя, у которого были изменения.
Например так: https://github.com/avp365/hl-sn/blob/348875a8b2f1e1765cb2ecc7436cac3c6cfb4945/internal/routers/post.go#L40

Напишмем сервис https://github.com/avp365/hl-sn/blob/cache_queue/cmd/cl/main.go, который будет запускать функцию https://github.com/avp365/hl-sn/blob/348875a8b2f1e1765cb2ecc7436cac3c6cfb4945/internal/routers/cache.go#L26 для сбора друзей по ID взятой из очереди и обновления их кешей.

В данном случае у нас есть два механизма сброса кеша. Время и очереди. В зависимости от нагрузки и наблюдениям за проектом можно будет их тюнить по тем или иным правилам. В данной работе представляен базовый механизм, один из многих.