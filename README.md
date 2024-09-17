
# Translate Word API

Translate Word API — это API, которое принимает слова на английском языке и переводит их на русский.

## Установка

Для установки выполните следующие шаги:

```bash
git clone https://github.com/fentezi/translate-word-api
cd translate-word-api
```

## Использование

Проект использует docker-compose.yaml для запуска. Чтобы запустить проект, выполните команду:
```bash
docker compose up
```

## Конфигурация

Конфигурационные параметры находятся в файле config/config.yml. Пример структуры конфигурационного файла:
```yaml
server:
  host: "0.0.0.0"
  port: 8080
database:
  host: "redis_container"
  port: "6379"
  name: 0
  password: ""
env: "prod" // env, local
```

## Описание конфигурации

 - **server.host**: Хост для API (по умолчанию: `0.0.0.0`).
 - **server.port**: Порт для API (по умолчанию: `8080`).
 - **database.host**: Хост базы данных Redis (по умолчанию: `redis_container`).
 - **database.port**: Порт Redis (по умолчанию: `6379`).
 - **database.name**: Название базы данных Redis (по умолчанию: `0`).
 - **database.password**: Пароль для Redis (по умолчанию пустой).
 - **env**: Среда выполнения (по умолчанию: `prod`).

## Лицензия

Этот проект распространяется под лицензией MIT. Подробности можно найти в файле LICENSE.