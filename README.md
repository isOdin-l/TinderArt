
Warning:
Архитектура и workflow на данный момент устаревшие. При разработке было решено преейти на другие решения:
  - Auth service больше не отвечает за Create User - Auth service только работает с jwt: валидация, создание
  - Auth service валидирует токены по gRPC в middleware каждого сервиса, а не через nginx как раньше

## Architecture
```mermaid
flowchart TB
classDef redis fill:#d82c20,stroke:#333,stroke-width:2px,color:white;

Client["Client"]
Gateway["API Gateway"]

Auth["Auth-Service"]
Profile["Profile-Service"]
Swipe["Swipe-Service"]
Stack["DailyStack-Service"]
Notify["Notify-Service"]
ExternalApis["External notify APIs"]

Cache[("Redis")]:::redis
DB_Profile[("PostgreSQL+PostGIS")]
DB_Swipes[("PostgreSQL")]
S3@{ shape: lin-cyl, label: "RustFS" }
Msg_Broker["Kafka"]

Client <--> Gateway

Gateway <--> Swipe
Gateway <--> Profile
Gateway <--> Auth


Profile <--> | CRUD photos | S3
Profile <--> | Get popular profiles, stack / Update stack for 1 user| Cache
Profile <--> | CRUD user data| DB_Profile
Profile <--> | Create user | Auth


Auth --> | Refresh token | DB_Profile

Swipe --> | Pub Match | Msg_Broker
Swipe <--> | Create swipe/Get is match| DB_Swipes

Msg_Broker -->| Sub Match | Notify

DB_Profile --> | Get profiles | Stack
Stack -->| Every day creates new stack | Cache

Notify --> | Notify users on diff platforms | ExternalApis

```
## Data structures

### Database tables
#### Profiles
| Field | Type |
|---|---|
| Id | uuid |
| Username | Text |
| Name | Text |
| Surname | Text |
| Email | Text |
| Password | Text |
| Description | Text |
| Latitude | Float |
| Longitude | Float |
| Created_at | TIMESTAMP|

#### JWT_Tokens
| Field | Type |
|---|---|
| Id | PK(profiles.id) |
| refresh_token | Text |

<br>

#### Arts
| Field | Type |
|---|---|
| Id | uuid |
| Author_id | uuid |
| Filename | Text |

<br>

#### Swipes
| Field | Type |
|---|---|
| userId_1 | uuid |
| userId_2 | uuid |
| desicion_1 | bool |
| desicion_2 | bool |


### Kafka topics
matches:
```
{
    "userId_1":{
        platform-data
    }
    "userId_2":{
        platform-data
    }
}
```

### Redis fields

```
Popular profile

Key: "pop_" + $profile_id
Value: user-data - JSON
```


```
User's stack
Key: "stack_" + $profile_id
Value: [ {name, surname, description, geo} ] - ZSET
```

### S3 storage
```
Art-bucket:
filename    UUID
```

## Workflow

### <a href="./docs/workflow-informative.md">Informative workflow graphics</a> <br>
### <a href="./docs/workflow-simple.md">Simple workflow graphics</a>

### For example:
#### Simple graph:
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

#### Informative graph:
```mermaid
sequenceDiagram
    participant Client
    participant Gateway         as API Gateway
    participant Auth            as Auth-Service
    participant Profile         as Profile-Service
    participant DB              as PostgreSQL
    participant Cache           as Redis

    Client->>Gateway: PATCH /profile
    Gateway->>Auth: validate(access_token)
    Auth-->>Gateway: validation result

    alt Unauthorized
        Gateway-->>Client: Unauthorized
    else Authorized
        Gateway->>Profile: updateProfile(data)

        Profile->>DB: UPDATE user's preferences
        DB-->>Profile: ok

        Profile->>Cache: DEL "pop_":{user_id}

        Profile-->>Gateway: ok
        Gateway-->>Client: ok
    end
```

Нагрузочное тестирование:
![Нагрузочное тестированик](assets/load_test.jpg)

<a href="./tests/results.md">Other tests results</a>

For future:
- Add photo 
- Delete photo
