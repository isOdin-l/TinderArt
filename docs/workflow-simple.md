
## Create Profile

```mermaid
flowchart LR

Client["Client"]
Gateway["API Gateway"]
Auth["Auth-Service"]
Profile["Profile-Service"]
DB_Profile[("PostgreSQL+PostGIS")]

Client -->|1| Gateway -->|2| Profile-->|3| Auth-->|4| Profile 
Profile-->|5.1| Gateway -->|6.1| Client

Profile <-->|5.2| DB_Profile
Profile-->|6.2| Gateway -->|7.2| Client

```
1 - Запрос от Клиента попадает на API Gateway <br>
2 - Переадресация на Profile service <br>
3 - по gRPC преедаём jwt-токены на Auth service для валидации<br>
4 - Получаем ответ <br>
5.1-6.1 - Unauthorized<br>
5.2 - Если ответ от Auth-service "ok", то делаем запись в БД <br>
6.2-7.2 - Передаём ответ на API Gateway и на Client <br>

<br>

## Get Profile
```mermaid
flowchart LR

Client["Client"]
Gateway["API Gateway"]
Auth["Auth-Service"]
Profile["Profile-Service"]
Cache[("Redis")]:::redis
DB_Profile[("PostgreSQL+PostGIS")]

Client -->|1| Gateway -->|2| Profile-->|3| Auth-->|4| Profile 
Profile-->|5.1| Gateway -->|6.1| Client

Profile -->|5.2| Cache

Cache -->|6.2| Profile
Cache -->|5.2.1| DB_Profile  
DB_Profile -->|5.2.2| Cache

Profile -->|7| Gateway
Gateway -->|8| Client
```
1-2 - Запрос от Клиента на API Gateway, а потом на Profile service <br>
3-4 - Валидация токена <br>
5.1-6.1 - Unauthorized user - возврат ответа на Client<br>
5.2 - Authorized; Запрос в Cache для получения профиля <br>
6.2 - Получение пользователя <br>
5.2.1 - Cache miss, поэтому идём в БД за пользователем <br>
5.2.2 - Кладём профиль в Cache <br>
7-8 - Передача профиля на API Gateway и на Client <br>

<br>

## Update Profile
```mermaid
flowchart LR

Client["Client"]
Gateway["API Gateway"]
Auth["Auth-Service"]
Profile["Profile-Service"]
DB_Profile[("PostgreSQL+PostGIS")]
Cache[("Redis")]


Client -->|1| Gateway -->|2| Profile-->|3| Auth-->|4| Profile 
Profile-->|5.1| Gateway -->|6.1| Client

Profile<-->|5.2| DB_Profile
Profile<-->|6.2| Cache

Profile -->|7| Gateway-->|8| Client
```
1-2 - PATCH Запрос от Клиента на API Gateway, а потом на Profile service <br>
2-4 - Валидация по access_token <br>
5.1-6.1 - Unauthorized user - возврат ответа на Client<br>
5.2 - Authorized; Запрос в БД на изменение данных <br>
6.2 - Кладём новый токен "profile:{user_id}" <br>
7-8 - Передаём результат на API Gateway и на Клиент <br>

<br>

## Delete Profile

```mermaid
flowchart LR

Client["Client"]
Gateway["API Gateway"]
Auth["Auth-Service"]
Profile["Profile-Service"]
DB_Profile[("PostgreSQL+PostGIS")]
S3["MinIO (S3)"]
Cache[("Redis")]

Client -->|1| Gateway -->|2| Profile-->|3| Auth-->|4| Profile 
Profile-->|5.1| Gateway -->|6.1| Client

Profile <--> |5| DB_Profile
Profile <-->|6| S3
Profile -->|7| Cache

Profile -->|8| Gateway -->|9| Client
```
1-2 - Запрос от Клиента на API Gateway, а потом на Profile service <br>
3-4 - Валидация токена <br>
5.1-6.1 - Unauthorized user - возврат ответа на Client<br>
5 - Получаем список URL фотографий пользователя и удаляем пользователя из БД<br>
6 - Удаляем фотографии из S3 <br>
7 - Удаляем "profile:{user_id}" и "stack:{user_id}" кэш <br>
8-9 - Передаём результат на API Gateway и на Клиент <br>

<br>

## Get Stack
```mermaid
flowchart LR

Client["Client"]
Gateway["API Gateway"]
Auth["Auth-Service"]
Profile["Profile-Service"]
Cache[("Redis")]:::redis
DB_Profile[("PostgreSQL+PostGIS")]

Client -->|1| Gateway -->|2| Profile-->|3| Auth-->|4| Profile 
Profile-->|5.1| Gateway -->|6.1| Client

Profile -->|5| Cache

Cache -->|6.1| Profile
Cache -->|6.2.1| DB_Profile  
DB_Profile -->|6.2.2| Cache

Profile -->|7| Gateway
Gateway -->|8| Client
```
1-2 - Запрос от Клиента на API Gateway, а потом на Profile service <br>
3-4 - Валидация токена <br>
5.1-6.1 - Unauthorized - возврат ответа на Client<br>
5 - Authorized; Запрос в Cache для получения части стопки <br>
6.1 - Cache hit <br>
6.2.1 - Cache miss, поэтому идём в БД за стопкой <br>
6.2.2 - Выдёляем часть стопки для возврата клиенту, а часть кладём большую часть кладём в Cache <br>
7-8 - Передача части стопки на API Gateway и на Client <br>


## Generate Daily Stack
```mermaid
flowchart LR

Stack["DailyStack-Service"]
Cache[("Redis")]:::redis
DB_Profile[("PostgreSQL+PostGIS")]

Time-->|1| Stack-->|2| DB_Profile-->|3| Stack-->|4| Cache

```

1 - Запуск задачи по времени (каждые 24 часа)<br>
2-3 - Идём в БД для получения пользователей, которые подходят по критериям <br>
4 - Отпраляем данные пользователей в Cache


## Sign-in
```mermaid
flowchart LR

Client["Client"]
Gateway["API Gateway"]
Auth["Auth-Service"]
DB_Profile[("PostgreSQL+PostGIS")]

Client-->|1| Gateway -->|2| Auth-->|3| DB_Profile

DB_Profile-->|4| Auth-->|5| Gateway-->|6| Client

```
1) Запрос от клиента на API Gateway
2) Передача на Auth
3) Запрос в БД:
    1) Refresh_token существует
    2) Пользователь существует
    3) Создание нового refresh_token
4) Получение от БД нового токена либо ошибки
5) Передача на API Gateway
6) Передача на клиент

<br>

## Refresh access token
```mermaid
flowchart LR

Client["Client"]
Gateway["API Gateway"]
Auth["Auth-service"]
DB_Profile[("PostgreSQL+PostGIS")]

Client-->|1| Gateway-->|2| Auth-->|3| DB_Profile
DB_Profile-->|4| Auth-->|5| Gateway-->|6| Client
```

1 - Запрос от Клиета на API Gateway <br>
2 - Передача на Auth <br>
3-4 - Запрос к БД на существование refresh_token <br>
5 - Передача нового access_token либо возврат ошибки <br>
6 - Передача результата на клиент <br>

<br>

## Create Swipe
```mermaid
flowchart LR

Client["Client"]
Gateway["API Gateway"]
Auth["Auth-Service"]
Swipe["Swipe-Service"]
DB_Swipe[("PostgreSQL")]
Kafka[("Kafka")]
Notify["Notify-Service"]
ExternalAPI["Notification API"]

Client -->|1| Gateway -->|2| Swipe-->|3| Auth-->|4| Swipe 
Swipe-->|5.1| Gateway -->|6.1| Client

Swipe -->|5| DB_Swipe-->|6| Swipe 
Swipe -->|7.1| Kafka -->|8.1| Notify -->|9.1| ExternalAPI

```
1-2 - Запрос от Клиента на API Gateway, а потом на Swipe service <br>
3-4 - Валидация токена <br>
5.1-6.1 - Unauthorized - возврат результата на Client
5 - Authorized; Запись свайпа в БД <br>
6 - Получение результата: есть ли match <br>
7.1 - Если match=true: Pub в Kafka "match":{user1:user2} <br>
8.1 - Notify service Sub из Kafka <br>
9.1 - Nofity service отправляет запрос во внешние API для создания уведомления <br>
