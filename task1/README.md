ДЗ No1. Знакомство с Linux системами. Написание сервиса погоды.
Написать сервис для отдачи прогноза погоды.
Конфигурация (environment variable): - listen port (LISTEN_PORT)
Данные брать из внешнего API прогноза погоды.
URL внешнего API вынести в конфигурационный файл или env переменную.

API (HTTP/1.1)
request (получение прогноза):
GET /v1/forecast/?city=<city>&dt=<timestamp>
response: {
"city": "Moscow", "unit": "celsius", "temperature": 32
}

API (HTTP/1.1)
request (получение текущей погоды):
GET /v1/current/?city=<city>
response: {
"city": "Moscow", "unit": "celsius", "temperature": 25
}