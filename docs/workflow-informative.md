## Create Profile
```mermaid
sequenceDiagram
    Client->>Gateway:       POST /register
    Gateway->>Profile:         register(profileData)

    Profile->>PostgreSQL:   insert user
    alt Already Exists
        PostgreSQL->>Profile: Error
        Profile->>Gateway: Error
        Gateway->>Client: Error
    else Insertion OK
        Profile-->>Auth:        gRPC - sign-tokens(profile -data)
        Auth->>PostgreSQL:      save refresh token
        Auth-->>Profile:        gRPC - ok
        Profile->>Client:       profile data
    end
```

<br>

## Get Profile
```mermaid
sequenceDiagram
    participant Client
    participant Gateway     as API Gateway
    participant Profile     as Profile-Service
    participant Auth        as Auth-Service
    participant Cache       as Redis
    participant DB          as PostgreSQL+PostGIS

    Note over Client, Gateway: 1 - Запрос от Клиента
    Client->>Gateway: GET /profile
    Gateway->>Profile: getProfile(user_id)

    Note over Gateway, Auth: 2 - Валидация по access_token
    Profile->>Auth: gRPC validate(access_token)
    Auth-->>Profile: validation result

    alt Unauthorized
        Gateway-->>Client: Unauthorized
    else Authorized
        Note over Gateway, DB: 3 - Получение профиля
        
        Profile->>Cache: get(user_id)

        alt Cache hit
            Cache-->>Profile: profile data
            Profile->>Gateway: profile data
            Gateway->>Client: profile data
        else Cache miss
            Profile->>DB: Get from DB
            alt Found
                DB-->>Profile: profile data
                Profile->>Cache: set(profile data)

                Profile-->>Gateway: profile data
                Gateway-->>Client: profile data
            else Not Found
                DB-->>Profile: 0 rows
                Profile-->>Gateway: Not Found
                Gateway-->>Client: Not Found
            end
        end

    end
```

<br>

## Update profile
```mermaid
sequenceDiagram
    participant Client
    participant Gateway         as API Gateway
    participant Profile         as Profile-Service
    participant Auth            as Auth-Service
    participant DB              as PostgreSQL
    participant Cache           as Redis

    Client->>Gateway: PATCH /profile
    Gateway->>Profile: updateProfile(data)
    Profile->>Auth: gRPC validate(access_token)
    Auth-->>Profile: validation result

    alt Unauthorized
        Profile-->>Client: Unauthorized
    else Authorized

        Profile->>DB: UPDATE user's preferences
        DB-->>Profile: ok

        Profile->>Cache: Update "profile:{user_id}"

        Profile-->>Gateway: ok
        Gateway-->>Client: ok
    end
```

<br>

## Delete profile
```mermaid
sequenceDiagram
    participant Client          as Client
    participant Gateway         as API Gateway
    participant Profile         as Profile-Service
    participant Auth            as Auth-Service
    participant DB              as PostgreSQL
    participant S3              as RustFS
    participant Cache           as Redis

    Client->>Gateway: DELETE /profile
    Gateway->>Profile: deleteProfile(user_id)
    Profile->>Auth: validate(access_token)
    Auth-->>Profile: validation result

    alt Unauthorized
        Profile-->>Client: Unauthorized
    else Authorized
        Profile->>DB: get photos URLs
        DB-->>Profile: photos

        loop delete photos
            Profile->>S3: delete object
        end

        Profile->>DB: DEL row by user_id
        DB-->>Profile: ok

        Profile->>Cache: DEL "profile:{user_id}"
        Profile->>Cache: DEL "stack:{user_id}"

        Profile-->>Gateway: ok
        Gateway-->>Client: ok 
    end
```

<br>

## Generate Daily Stack
```mermaid
sequenceDiagram
    participant Timer as InternalTimer
    participant Stack as DailyStack-Service
    participant DB as PostgreSQL+PostGIS
    participant Cache as Redis

    Note over Timer,Stack: 1 - Старт сервиса по Timer каждые 24 часа
    Timer->>Stack: trigger daily stack generation

    Note over Stack,DB: 2 - Получение подходящих пользователей
    Stack->>DB: Получение пользователей по критериям
    DB-->>Stack: Список профилей

    Note over Stack,Cache: 3 - Сохранение стопки в cache
    Stack->>Cache: set daily stack for users
```

<br>

## Get Stack
```mermaid
sequenceDiagram
    participant Client
    participant Gateway     as API Gateway
    participant Profile     as Profile-Service
    participant Auth        as Auth-Service
    participant Cache       as Redis
    participant DB          as PostgreSQL+PostGIS

    Note over Client, Gateway: 1 - Запрос от Клиента
    Client->>Gateway: GET /profile

    Gateway->>Profile: getStack(for user_id)
    
    Note over Gateway, Auth: 2 - Валидация по access_token
    Profile->>Auth: validate(access_token)
    Auth-->>Profile: validation result

    alt Unauthorized
        Profile-->>Client: Unauthorized
    else Authorized
        Note over Gateway, DB: 3 - Получение профиля
        
        Profile->>Cache: get("stack:{$user_id}")

        alt Cache hit
            Cache-->>Profile: part of stack
            Profile->>Gateway: part of stack
            Gateway->>Client: part of stack
        else Cache miss
            Profile->>DB: Get profiles from DB
            alt Found
                DB-->>Profile: profiles
                Profile->>Cache: set("stack:{$user_id} -> [userId, userId,...])

                Profile-->>Gateway: part of stack
                Gateway-->>Client: part of stack
            else Not Found
                DB-->>Profile: 0 rows
                Profile-->>Gateway: Not Found
                Gateway-->>Client: Not Found
            end
        end

    end
```

<br>

### Sign-in
```mermaid
sequenceDiagram

Client->>Gateway:       POST /sign-in
Gateway->>Auth:         sign-in(username, password)
Auth->>DB_profiles:     check userData
Auth-->>Client:         new access + refresh / error

```

<br>

### Refresh access token
```mermaid
sequenceDiagram

Client->>Gateway:       POST /refresh
Gateway->>Auth:         refresh(refreshToken)
Auth->>DB_profiles:     check refresh_token
Auth-->>Client:         new access / error
```

<br>

### Create Swipe
```mermaid
sequenceDiagram
    participant Client
    participant Gateway as API Gateway
    participant Swipe as Swipe-Service
    participant Auth as Auth-Service
    participant DB as PostgreSQL
    participant Kafka
    participant Notify as Notify-Service
    participant External as ExtarnalAPIs

    Client->>Gateway: POST /swipe

    Gateway->>Swipe: createSwipe(user1_id, user2_id)
    Swipe->>Auth: validate(access_token)
    Auth-->>Swipe: validation result

    alt Unauthorized
        Swipe-->>Client: Unauthorized
    else Authorized

        Swipe->>DB: INSERT swipe
        DB-->>Swipe: result

        alt Match = true
            Swipe->>Kafka: Pub into Match topic

            Kafka-->>Notify: Sub Match topic

            Notify->>External: notify users
        end

        Swipe-->>Gateway: results
        Gateway-->>Client: response
    end
```
