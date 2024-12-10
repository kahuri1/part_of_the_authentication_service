# Часть сервиса аутентификации

## Описание

В этом проекте реализована часть сервиса аутентификации. В базе данных храниться хэш ревреш токена, юзеру отправляется закодированный ревреш токен в base64. Реализовал только одно подключение сессии для юзера, если авторизируешься повторно, сессия обновляется. Если время сессии истекло, то нужно провести аутентификацию повторно. Ручки для добавления юзера есть, но они не доделаны(не хватило времени, но нужно их дошлепать). Альтеративу я добавил посредством миграций, в БД через миграции добавляются два юзера. Так же реализовал подключение к почте, если дернуть ручку createUser, на почту прийдет код подтверждения(да и логику кода подтверждения тоже можно добавить) так же если ip не сходится, то на почту приходит письмо о подозрительном входе. 

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
5. В миграции creating_users измените почту клиента
5. И если ничего не забыл, то запустите проект)
6. Для удобства можете импортировать ручки в postman. Файл в проекте

REST маршруты

1. **Получение токенов**
   - **Метод**: `POST`
   - **Маршрут**: `/auth`
   - **Параметры**: Идентификатор пользователя (GUID) в теле запроса.
   - **Ответ**: Пара токенов (Access(JWT) и Refresh(encoding base64)).
     ```bash
     http://localhost:8000/auth
     ```
     Пример запроса
```json
  {
    "uuid" : "dff723ba-4da7-4c55-9f07-27121ec53385"
  }
```
пример ответа
```json
{
    "tokens": {
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDI0LTEyLTEwVDE3OjUzOjMzLjA1OTQ5ODMrMDM6MDAiLCJpYXQiOjE3MzM4NDE1MTMsImlwIjoiOjoxIiwiaXNzIjoidG9kby1hcHAiLCJzdWIiOiJkZmY3MjNiYS00ZGE3LTRjNTUtOWYwNy0yNzEyMWVjNTMzODUifQ.QlPTA9eXZctJryDtOOvHh_k6dFiqXEP_DOVMFashSME",
        "refresh_token": "JDJhJDEwJG1VSTJuNmNhRi9OUU1TaS5CTk5VcC5PZVU3MWlSdWRXN25QY2g1S09ZN3Y4NzdtWGhZWWJ1"
    }
}
```
2. **Refresh токенов**
   - **Метод**: `POST`
   - **Маршрут**: `/auth/refresh`
   - **Параметры**: Refresh токен в теле запроса.
   - **Ответ**: Новая пара токенов (Access(JWT) и Refresh(encoding base64)).
пример запроса
```json
{
    "refresh_token": "JDJhJDEwJG1VSTJuNmNhRi9OUU1TaS5CTk5VcC5PZVU3MWlSdWRXN25QY2g1S09ZN3Y4NzdtWGhZWWJ1"
}
 ```
пример ответа
 ```json
{
    "tokens": {
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDI0LTEyLTEwVDE3OjU0OjI4LjU1NzEyMyswMzowMCIsImlhdCI6MTczMzg0MTU2OCwiaXAiOiIxMjcuMC4wLjEiLCJpc3MiOiJ0b2RvLWFwcCIsInN1YiI6ImRmZjcyM2JhLTRkYTctNGM1NS05ZjA3LTI3MTIxZWM1MzM4NSJ9.Mvct6Sa3UuO6wLBb5lgk0qODA1PmtmVwynq4lxiyeUw",
        "refresh_token": "JDJhJDEwJGJ3dWtnM1R3NWFNenlEYTlRMHhqM09WSEU5YkhsdEkyUWpZajJSWDdWQWZuOU5MaE5Wdk9l"
    }
}
 ```


## JWT token хранит следующию информацию
**HEADER:ALGORITHM & TOKEN TYPE**
```bash
{
  "alg": "HS256",
  "typ": "JWT"
}
```
**PAYLOAD**
```bash
{
  "exp": "2024-12-10T17:54:28.557123+03:00",
  "iat": 1733841568,
  "ip": "127.0.0.1",
  "iss": "todo-app",
  "sub": "dff723ba-4da7-4c55-9f07-27121ec53385"
}
```
