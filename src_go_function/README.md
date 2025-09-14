### Golang Telegram Bot

Пример эхо-бота на фреймворке Golang Telegram Bot

https://github.com/go-telegram/bot/


#### Как это работает

Используются сервисы:
- Yandex Identity and Access Management (сервисный аккаунт https://yandex.cloud/ru/docs/iam/concepts/users/service-accounts)
- Yandex Cloud Functions https://yandex.cloud/ru/docs/functions/lang/golang/

Общий принцип такой:
- Код бота представяет собой версию публичной функции в Yandex Cloud Functions.
- Телеграм отправляет боту апдейты через webhook, в качестве которого выступает функция.
- Функция запускается и бот обрабатывает апдейт.

Нюансы:
- Сделать функцию публичной https://yandex.cloud/ru/docs/functions/operations/function/function-public
- Апдейты от Телеграма передаются в функию как Json-строки.
- Непосредственно запуск бота происходит черезе ручную передачу ему апдейта, который парсится из Json-строки.

Код бота не отличается от кода для шлюза.
