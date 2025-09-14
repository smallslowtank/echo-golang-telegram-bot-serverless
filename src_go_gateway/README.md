### Golang Telegram Bot

Пример эхо-бота на фреймворке Golang Telegram Bot

https://github.com/go-telegram/bot/


#### Как это работает

Используются сервисы:
- Yandex Identity and Access Management (сервисный аккаунт https://yandex.cloud/ru/docs/iam/concepts/users/service-accounts)
- Yandex API Gateway https://yandex.cloud/ru/docs/api-gateway/
- Yandex Cloud Functions https://yandex.cloud/ru/docs/functions/lang/golang/

Общий принцип такой:
- Код бота представяет собой версию функции в Yandex Cloud Functions.
- Телеграм отправляет боту апдейты через webhook, в качестве которого выступает шлюз Yandex API Gateway.
- При получении апдейта, шлюз вызывает функцию (бота).
- Функция при запуске в качестве параметра получает апдейт, пришедший на шлюз.
- Бот обрабатывает апдейт.

Нюансы:
- Апдейты от Телеграма передаются через шлюз как Json-строки.
- Непосредственно запуск бота происходит черезе ручную передачу ему апдейта, который парсится из Json-строки.
