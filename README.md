# Go Blog 

Go Blog - готовая основа для быстрой сборки backend-сервисов на основе `Go Fiber`, Документация на основе `Swagger`, в соответствии со стандартом OpenAPI.

  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;

![Python](https://img.shields.io/badge/go-v1.20.1+-blue.svg)

![Contributions welcome](https://img.shields.io/badge/contributions-welcome-orange.svg)

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

  

## 📋 Table of Contents

  

1. 🌀 [Описание проекта](#what-is-this)
2. 📈 [Краткая документация API](#api_docs)
3. 💾 [База данных](#database_scheme)
4. 🚀 [Инструкция по установке](#installation)
5. ©️ [License](#license)

  

## <a name="what-is-this"> 🌀 Описание проекта</a>

Go Blog - готовая основа для быстрой сборки backend-сервисов на основе `Go Fiber`, База данных - `PostgreSQL`. ORM - `GORM`. Интерфейс API `Swagger`.

Контейнеризация - `Docker`.

## <a name="api_docs"> 📈 Краткая документация API</a>

Работа с моделями осуществляется по следующим эндпоинтам:

| Method | HTTP request | Description |

| ------------- | ----------------------------- | ------------------------------------------------- |

| [**POST**] | /api/v1/register | Регистрация нового пользователя |
| [**POST**] | /api/v1/login | Авторизация пользователя про логину и паролю |
| [**GET**] | /api/v1/current_user | Извлечение авторизованного пользователя по токену |
| [**GET**] | /api/v1/logout | Разлогиниться |
| [**POST**] | /categories | Создать новую категорию |
| [**GET**] | /categories | Извлечь все категории |
| [**POST**] | /categories/add_article | Добавить статью в категорию |
| [**POST**] | /categories/remove_article | Удалить статью из категории |
| [**POST**] | /articles | Извлечь все статьи |
| [**POST**] | /articles | Создать новую статью |
| [**GET**] | /articles/:id | Извлечь статью по ID |
| [**PUT**] | /articles/:id | Обновить статью (только для авторов) |
| [**DELETE**] | /articles/:id | Удалить статью (только для авторов) |
| [**POST**] | /article/{id}/add_comment | Добавить комментарий к статье |

## <a name="database_scheme"> 💾 База данных </a>

  

База данных содержит 6 моделей:
**Автор публикации** (User),
**Категория статьи** (Category),
**Cтатья** (Article),
**Статья в категории** (ArticleCategory),
**Комментарий** (Comment)

  
  
  

<details>

<summary>ДЕТАЛЬНАЯ ИНФОРМАЦИЯ О МОДЕЛЯХ </summary>

  
  

</details>

  

<details>

<summary>ДЕТАЛЬНАЯ СХЕМА БАЗЫ ДАННЫХ</summary>

  

![Screen Shot](docs/extras/erd.jpg)

  

</details>

  

## <a name="installation"> 🚀 Установка и использование</a>

  

1. ### Подготовка проекта

  

1.1 Клонируете репозиторий

```sh

git clone https://github.com/XanderMoroz/goBlog.git

```

1.2 В корневой папки создаете файл .env

1.3 Заполняете файл .env по следующему шаблону:

```sh

# JWT SETTINGS

JWT_SECRET_KEY="SomeAppSecret"

# POSTGRES SETTINGS
DB_DRIVER=postgres
DB_USER=xander
DB_PASSWORD=password
DB_NAME=go_blog_api
DB_PORT=5432
DB_HOST=go_blog-postgres # С docker
# DB_HOST=127.0.0.1 # Без docker

# POSTGRES TEST SETTINGS
TEST_DB_HOST=go_blog-postgres_test # С docker
# TEST_DB_HOST=127.0.0.1 # Без docker
TEST_DB_DRIVER=postgres
TEST_API_SECRET=some_secret
TEST_DB_USER=xander
TEST_DB_PASSWORD=password
TEST_DB_NAME=go_blog_api_test
TEST_DB_PORT=5432

# PGADMIN SETTINGS
PGADMIN_DEFAULT_EMAIL=xander@admin.com
PGADMIN_DEFAULT_PASSWORD=pwd123

```

2. ### Запуск проекта с Docker compose

2.1 Создаете и запускаете контейнер через терминал:

```sh

sudo docker-compose up --build

```

2.3 Сервисы доступны для эксплуатации:

- Fiber APP: http://127.0.0.1:8080/
- Swagger: http://127.0.0.1:8080/swagger/index.html
- PGAdmin4: http://127.0.0.1:5050
- Prometheus: http://127.0.0.1:9090
- Grafana: http://127.0.0.1:3000


3. ### Дополнительные настройки 

<details>
<summary>Как подключить PGAdmin4 к БД? </summary>

1. Заходим в браузер по адресу http://127.0.0.1:5050 и вводим данные из .env

```bash
PGADMIN_DEFAULT_EMAIL=guest@admin.com
PGADMIN_DEFAULT_PASSWORD=pwd123
```
Картинка
  

</details>



5. Авто-генерация документации swagger

3.1 Как подключить PGAdmin4 к БД





Устанавливаете swag

```sh

go get github.com/swaggo/swag/cmd/swag

```

  

3.2 Устанавливаете GOPATH

```sh

export PATH=$PATH:$(go env GOPATH)/bin

```

  

3.3 Генерируете новый вариант документации

```bash

swag init -g main.go

```

  

4. ### Как подключить pgadmin к контейнеру с БД (postgres)

  

4.1 Поднимаем контейнеры

```bash

sudo docker-compose up --build

```

  


  

4.3 Уточняем порт, на котором работает БД (чтоб подключиться к ней)

```bash

sudo docker inspect go_blog_postgres | grep IPAddress

```

  

4.3 Создаем сервер и настраиваем подключение к БД

  
  

## <a name="license"> ©️ License
```
sudo ufw allow 9090/tcp
```
sudo docker inspect prometheus | grep IPAddress