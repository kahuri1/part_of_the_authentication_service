# Часть сервиса аутентификации

## Описание

В этом проекте реализована часть сервиса аутентификации

## Используемые технологии
**Go**, **JWT**, **Postgres**, **Docker**




##Установка

1.Клонируйте репозиторий
```bash
   git clone https://github.com/kahuri1/part_of_the_authentication_service.git
```
2. Установите зависимости
```bash
  go mod tidy
```
3. Поднимите Docker
```bash
   docker-compose up
```
4. В папке Configs создайте config.yml и скопируйте настройки:
```bash
port: "8000"

db:
  username: "myuser"
  host:     "localhost"
  port:     "5432"
  dbname:   "part_auth"
  sslmode:  "disable"
  password: "mypassword"

auth:
  accessTokenTTL: 15m
  refreshTokenTTL: 30m
  verificationCodeLength: 10

key:
  secretKey: "Ваш секретный ключ"


email :
  email: "Почта с который будут идти уведомления"
  password: "пароль приложения (я делал через него)"
  server: "Сервер"
  port: "порт"
```
5. И если ничего не забыл, то запустите проект)
6. Для удобства можете импортировать ручки в postman, файл в проекте
