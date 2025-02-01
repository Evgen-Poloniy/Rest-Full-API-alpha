# 1. Собрать docker-compose

```
make build
```

# 2. Поднять контейнеры

```
make up
```

> Про запросы смотрите пункт 5

# 3. Собрать, а затем поднять

```
make all
```

# 4. Посмотреть список образов и работающих контейнеров

```
make list
```

# 5. Совершение запросов к БД

> Запросы вводятся в другом терминале (сделано для того, чтобы видеть лог БД и Go API)

* 5.1. Скомпилируйте исполняемый файл с помощью:

```
make comp
```

* 5.2. Запустите исполняемый файл СУБД:

```
make run
```

> Запросы можно вводить без СУБД напрямую в терминал

____

* СУБД может отвечать на следующие запросы: кол-во пользователей в БД, найти запись о пользователе по его **id**, вывести всех пользователей