# Calculator 
Установка
Клонируйте репозиторий:

git clone https://github.com/Jl1KUR/LMS.git

Перейдите в директорию проекта:

cd LMS

Запустите приложение на локальном сервере (например, с помощью командной строки).

go run main.go

Примеры запросов
Пример 1: Корректное выражение

curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
"expression": "2+2*2"
}'

Пример 2: Некорректное выражение

curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
"expression": "2a+3"
}'

Ответы сервера

Пример ответа на корректный запрос
json
{
"result": 6
}
Пример ответа на некорректный запрос
json
{
"error":"Expression is not valid"
}