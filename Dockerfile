# Используйте официальный образ PostgreSQL
FROM postgres

# Установите необходимые пользователи и пароли (замените на свои значения)
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=mypassword
ENV POSTGRES_DB=myDB

# Копируйте SQL-файл с инициализацией в контейнер
COPY createTables.sql /docker-entrypoint-initdb.d/

# Укажите порт, на котором PostgreSQL будет слушать (по умолчанию 5432)
EXPOSE 5432
