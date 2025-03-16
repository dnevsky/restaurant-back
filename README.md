# restaurant-back
The backend for the restaurant's website.

Поднять миграции (локально):

```makefile
make migration-up PG_DSN="host=localhost user=user password=pgpwd4 dbname=restaurant port=54324 sslmode=disable TimeZone=Europe/Moscow"
```