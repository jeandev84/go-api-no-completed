# STRUCTURE PROJECT

```
app/
 - config (хранится все наши конфиги)
 - controllers (контроллер запроса)
 - entity (хранится все наши сущности)
 - migrations (хранится все миграции)
 - repository (репозитория для выполнения запроса к БД)
 - utils (вспомогательные функции)

public/
 - main.go (точка входа прилодения)

storage/
 - cache (хранилище кэша)
 - log   (хранилище лога)

```

# Packages
``go-martini/martini — мощный URL-маршрутизатор и диспетчер. Мы будем использовать этот пакет для сопоставления путей URL с их обработчиками.``

``jinzhu / gorm — отличная ORM-библиотека для Golang. Мы используем этот пакет, чтобы взаимодействовать с базой данных.``

``dgrijalva / jwt-go — используется для подписи и проверки токенов JWT.``

``joho / godotenv — используется для загрузки файлов .env в проект.``


# INIT MODULE
``go mod init {module_name}, Example: go mod init authuser``

# GET MODULE (github) 
``go.mod is like composer.json and go.sum like composer.lock``

``go get github.com/{package-name}, Example: "go get github.com/gomarkdown/markdown"`` 


# Init module and generated file go.mod, go.sum
``from github: go mod init github.com/gomarkdown/markdown``


# BUILD MODULE
``go build``


# LUNCH SERVER (Запускаем проекта)
``go run ./public/main.go``
