
## Create Profile

```mermaid
flowchart LR

Client["Client"]
Gateway["API Gateway"]
Auth["Auth-Service"]
Profile["Profile-Service"]
DB_Profile[("PostgreSQL+PostGIS")]

Client -->|1| Gateway -->|2| Auth-->|3| Profile-->|4| DB_Profile-->|5| Profile-->|6| Auth

Auth<-->|7.1| DB_Profile 
Auth-->|8| Gateway-->|9| Client
```
1 - Запрос от Клиента попадает на API Gateway <br>
2 - Переадресация на Auth service <br>
3 - по gRPC преедаём данные пользователя на Profile service<br>
4 - Создаём запись в БД <br>
5-6 - Получаем результат и передаём<br>
7.1 - Если ответ от Profile-service "ok", то делаем запись о refresh_token в БД и получаем ответ <br>
8-9 - Передаём ответ на API Gateway и на Client <br>

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

Client -->|1| Gateway -->|2| Auth -->|3| Gateway

Gateway -->|4.1| Client
Gateway -->|4.2| Profile -->|5| Cache

Cache -->|6.1| Profile
Cache -->|6.2.1| DB_Profile  
DB_Profile -->|6.2.2| Cache

Profile -->|7| Gateway
Gateway -->|8| Client
```
1 - Запрос от Клиента попадает на API Gateway <br>
2 - Валидация по access_token <br>
3 - Возврат результата <br>
4.1 - Unauthorized user <br>
4.2 - authorized <br>
5 - Запрос в Cache для получения профиля <br>
6.1 - Cache hit <br>
6.2.1 - Cache miss, поэтому идём в БД за пользователем <br>
6.2.2 - Кладём профиль в Cache <br>
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


Client -->|1| Gateway -->|2| Auth -->|3| Gateway -->|4.1| Client
Gateway-->|4.2| Profile

Profile<-->|5| DB_Profile
Profile<-->|6| Cache

Profile -->|7| Gateway-->|8| Client
```
1 - PATCH Запрос от клиента попадает на API Gateway <br>
2 - Валидация по access_token <br>
3 - Возврат результата <br>
4.1 - Unauthorized user <br>
4.2 - Authorized <br>
5 - Обновляем данные профиля в БД <br>
6 - Очищаем кэш "pop_":{user_id} <br>
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

Client -->|1| Gateway -->|2| Auth -->|3| Gateway -->|4.1| Client
Gateway -->|4.2| Profile <--> |5| DB_Profile

Profile <-->|6| S3
Profile -->|7| DB_Profile
Profile -->|8| Cache

Profile -->|9| Gateway -->|10| Client
```
1 - Запрос от клиента попадает на API Gateway <br>
2 - Валидация по access_token <br>
3 - Возврат результата <br>
4.1 - Unauthorized user <br>
4.2 - Authorized <br>
5 - Получаем из БД список URL фотографий пользователя <br>
6 - Удаляем фотографии из S3 <br>
7 - Удаляем пользователя из БД <br>
8 - Удаляем "pop_":{user_id} и "stack_":{user_id} кэш <br>
9-10 - Передаём результат на API Gateway и на Клиент <br>

<br>

## Update preferences
```mermaid
flowchart LR

Client["Client"]
Gateway["API Gateway"]
Auth["Auth-Service"]
Profile["Profile-Service"]
DB_Profile[("PostgreSQL+PostGIS")]
Cache[("Redis")]


Client -->|1| Gateway -->|2| Auth -->|3| Gateway -->|4.1| Client
Gateway-->|4.2| Profile

Profile<-->|5| DB_Profile
Profile<-->|6| Cache

Profile -->|7| Gateway-->|8| Client
```
1 - PATCH запрос от клиента идёт на API Gateway <br>
2 - Валидация по access_token <br>
3 - Возврат результата <br>
4.1 - Unauthorized user <br>
4.2 - Authorized <br>
5 - Обновляем предпочтения профиля в БД <br>
6 - Очищаем кэш "stack_":{user_id} <br>
7-8 - Передаём результат на API Gateway и на Клиент <br>

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

Client -->|1| Gateway -->|2| Auth -->|3| Gateway

Gateway -->|4.1| Client
Gateway -->|4.2| Profile -->|5| Cache

Cache -->|6.1| Profile
Cache -->|6.2.1| DB_Profile  
DB_Profile -->|6.2.2| Cache

Profile -->|7| Gateway
Gateway -->|8| Client
```
1 - Запрос от Клиента попадает на API Gateway <br>
2 - Валидация по access_token <br>
3 - Возврат результата <br>
4.1 - Unauthorized user <br>
4.2 - authorized <br>
5 - Запрос в Cache для получения части стопки <br>
6.1 - Cache hit <br>
6.2.1 - Cache miss, поэтому идём в БД за стопкой <br>
6.2.2 - Выдёляем часть стопки для возврата клиенту, а часть кладём большую часть кладём в Cache <br>
7-8 - Передача части стопки на API Gateway и на Client <br>


## Generate Daily Stack
```mermaid
flowchart LR
classDef redis fill:#d82c20,stroke:#333,stroke-width:2px,color:white;

Stack["DailyStack-Service"]

Cache[("Redis")]:::redis
DB_Profile[("PostgreSQL+PostGIS")]
Cron["CRON"]

Cron-->|1| Stack-->|2| DB_Profile-->|3| Stack-->|4| Cache

```

1 - Старт сервиса по Cron каждый день <br>
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
    1) Refresh_token не существует
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

Client -->|1| Gateway -->|2| Auth -->|3| Gateway

Gateway -->|4.1| Client
Gateway -->|4.2| Swipe -->|5| DB_Swipe

DB_Swipe -->|6| Swipe 
Swipe -->|7.1| Kafka -->|8.1| Notify -->|9.1| ExternalAPI

Swipe -->|7.2| Gateway -->|8.2| Client
```
1 - POST запрос от клиента на API Gateway <br>
2 - Валидация по access_token <br>
3 - Возврат результата <br>
4.1 - Unauthorized user <br>
4.2 - Authorized <br>
5 - Запись свайпа в БД <br>
6 - Получение результата: есть ли match <br>
7.1 - Если match=true: Pub в Kafka "match":{user1: data, user2: data} <br>
8.1 - Notify service Sub из Kafka <br>
9.1 - Nofity service отправляет запрос во внешние API для создания уведомления <br>
7.2-8.2 - Передача части стопки на API Gateway и на Client <br>