Примеры эхо-бота на фреймворке "Golang Telegram Bot" https://github.com/go-telegram/bot/ и стеке serverless-технологий Yandex Cloud.

#### src_go_only_function

Пример только с публичной функцией.

#### src_go_gateway

Пример без очереди сообщений. Апдейты в бота поступают сразу из шлюза.

#### src_go_mq

Пример с передачей апдейтов в бота через очередь сообщений.

#### Файлы

Версия функции состоит из черытёх файлов:

.env<br>
event.go<br>
go.mod<br>
index.go

В файле ```.env``` должен находиться токен бота. Пример в ```.env.example```

Примеры команд для подключения webhook-а в файле ```set-webhook.txt```. Команду можно запустить через Cloud Shell https://yandex.cloud/ru/docs/console/quickstart/cloud-shell

Примеры конфигов для шлюза в файле ```example_config_API_Gateway.yaml```
