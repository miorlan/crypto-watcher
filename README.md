CryptoWatcher API

API для отслеживания стоимости криптовалют. Позволяет добавлять, удалять и получать актуальные данные о ценах криптовалют.
Оглавление

    Описание

    Установка и запуск

    Использование API

        Добавить валюту

        Получить цену валюты

        Удалить валюту

    Примеры запросов

    Лицензия

    Контакты

Описание

CryptoWatcher API предоставляет возможность отслеживать стоимость криптовалют. API позволяет:

    Добавлять криптовалюты для отслеживания.

    Получать цену криптовалюты на определенный момент времени.

    Удалять криптовалюты из списка отслеживаемых.

API использует Swagger для документации и поддерживает стандартные HTTP-методы.
Установка и запуск

    Клонируйте репозиторий:
    bash
    Copy

    git clone https://github.com/yourusername/cryptowatcher-api.git
    cd cryptowatcher-api

    Установите зависимости:
    bash
    Copy

    npm install  # или pip install -r requirements.txt, в зависимости от вашего стека

    Запустите сервер:
    bash
    Copy

    npm start  # или python app.py, или другой команды для запуска

    Документация API:
    После запуска сервера документация будет доступна по адресу:
    Copy

    http://localhost:8080/api/v1/docs

Использование API
Добавить валюту

Добавляет новую криптовалюту для отслеживания.

Endpoint: POST /currency/add

Пример запроса:
{
  "coin": "bitcoin"
}

Пример ответа:
{
  "message": "Currency added"
}

Получить цену валюты

Возвращает цену криптовалюты на указанный момент времени.

Endpoint: GET /currency/price

Пример запроса:
json
Copy

{
  "coin": "bitcoin",
  "timestamp": 1638316800
}

Пример ответа:
{
  "price": 97675
}

Удалить валюту

Удаляет криптовалюту из списка отслеживаемых.

Endpoint: POST /currency/remove

Пример запроса:
{
  "coin": "BTC"
}

Пример ответа:
{
  "message": "Currency removed"
}

Примеры запросов
Использование cURL

Добавить валюту:
curl -X POST "http://localhost:8080/api/v1/currency/add" \
-H "Content-Type: application/json" \
-d '{"coin": "bitcoin"}'

Получить цену валюты:
curl -X GET "http://localhost:8080/api/v1/currency/price" \
-H "Content-Type: application/json" \
-d '{"coin": "bitcoin", "timestamp": 1638316800}'

Удалить валюту:
curl -X POST "http://localhost:8080/api/v1/currency/remove" \
-H "Content-Type: application/json" \
-d '{"coin": "bitcoin"}'
