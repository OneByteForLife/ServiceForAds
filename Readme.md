# Ads Service

## Описание
Проект был реализован с целью практики навыков в проектировании REST-API.

В корневой директории проекта есть файл с техническими требованиями для этого проекта **"TermsOfReference.md"**

## Сборка
**Сборка через docker:**
- docker-compose up --build
    
**Сборка без docker:**
- go build cmd/app/main.go (В таком случае файл конфигурации config.yaml должен быть в одной директории с исполняемым файлом)

## Структура проекта
``` ├── cmd
    │   └── app
    │       └── main.go
    ├── config
    │   ├── config.go
    │   └── config.yaml
    ├── docker-compose.yml
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── app
    │   │   └── app.go
    │   ├── controller
    │   │   ├── apache
    │   │   └── webapi
    │   │       ├── ads_handlers
    │   │       │   ├── ads.go
    │   │       │   └── handler.go
    │   │       ├── start.go
    │   │       └── user_handlers
    │   ├── entity
    │   │   └── ads.go
    │   └── usecase
    │       ├── ads
    │       │   ├── ads.go
    │       │   ├── service.go
    │       │   └── storage.go
    │       └── user
    ├── pkg
    │   ├── database
    │   │   └── postgres.go
    │   └── log
    │       └── log.go
    ├── Readme.md
    ├── sql
    │   └── init.sql
    ├── TermsOfReference.md
    └── tools
        ├── conv
        │   └── convert.go
        └── validation
            └── validate.go
```
  
## API Спецификация

URL - http://localhost:8080

### **/api/v1**

`GET` : Стартовая страница (Health Check)

```
curl --request GET \
    --url http://localhost:8080/api/v1/
```

### **/api/v1/ads/create**

`POST` : Создание нового объявления

#### Перечень требований для запроса
- Должен быть заголовок с типом передаваемого контента (application/json)

- Должен быть заголовок с авторизацией (если он требуется)

#### Перечень требований для json объекта
- Ссылки на фотографии должны быть валидны (содержать протокол и расширение файла)

- Описание товара должно быть в диапозоне от 200 до 1000 символов (такой же диапозон имеет длина ссылок)

- Основная фотография обязательна 

#### JSON
```
    {
        "product_name": "Заголовок",
        "description": "Описание продукта",
        "main_picture": "Ссылка на основное изображение",
        "second_pictures": [
            "Ссылки на дополнительные фото не больше 5",
            "...",
            "...",
            "...",
            "..."
        ],
        "product_price": 15000
    }
```

```
curl --request POST \
  --url http://localhost:8080/api/v1/ads/create \
  --header 'Content-Type: application/json' \
  --data '{
    "product_name": "Iphone 5s 64GB",
    "description": "В отличном состоянии. Зарядное устройство. Использовался с защитным стеклом. Ростест. Отвязан от всех аккаунтов. Чехол в подарок. Все детали в рабочем состоянии. Любые проверки на месте. Без сколов и дефектов. Торг уместен",
    "main_picture": "https://iphone5s/image.png",
    "second_pictures": [
                "https://iphone5s/image0.png",
                "https://iphone5s/image1.png",
                "https://iphone5s/image2.png",
                "https://iphone5s/image3.png",
                "https://iphone5s/image4.png"
            ],
    "product_price": 3000
}'
```

### **/api/v1/ads/get?id={int}**

`GET` : Получение объявления по его ID

```
curl --request GET \
  --url 'http://localhost:8080/api/v1/ads/get?id=1'
```

### **/api/v1/ads/get/all?limit={int}&offset={int}**

`GET` : Получение всех объявлений с указанием пагинации.

- Параметры limit и offset обязательны

```
curl --request GET \
  --url 'http://localhost:8080/api/v1/ads/get/all?limit=10&offset=0'
```

### **/api/v1/ads/get/all?limit={int}&offset={int}&sortBy={string}&sortType={string}**

`GET` : Получение всех объявлений с указанием пагинации.

#### Список доступных параметров для сортировки:
- price
- date_create

#### Список допустимых типов сортировки:
- asc
- desc

```
curl --request GET \
  --url 'http://localhost:8080/api/v1/ads/get/all?limit=10&offset=0&sortBy=price&sortType=desc'
```