### Golang Telegram Bot

Пример эхо-бота на фреймворке Golang Telegram Bot

https://github.com/go-telegram/bot/


#### Как это работает

Используются сервисы:
- Yandex Identity and Access Management (сервисный аккаунт https://yandex.cloud/ru/docs/iam/concepts/users/service-accounts)
- Yandex API Gateway https://yandex.cloud/ru/docs/api-gateway/
- Yandex Message Queue https://yandex.cloud/ru/docs/message-queue/
- Yandex Cloud Functions https://yandex.cloud/ru/docs/functions/lang/golang/ и триггер для очереди https://yandex.cloud/ru/docs/functions/concepts/trigger/ymq-trigger

Общий принцип такой:
- Код бота представяет собой версию функции в Yandex Cloud Functions.
- Телеграм отправляет боту апдейты через webhook, в качестве которого выступает шлюз Yandex API Gateway.
- Шлюз отправляет апдейты в очередь сообщений Yandex Message Queue.
- Срабатывает триггер и запускает функцию (бота).
- Функция при запуске в качестве параметра получает сообщение с апдейтом из Yandex Message Queue.
- Бот обрабатывает сообщение.

Нюансы:
- Апдейты от Телеграма передаются через очередь сообщений в качестве Json-строки.
- Непосредственно запуск бота происходит черезе ручную передачу ему апдейта, который парсится из Json-строки.
