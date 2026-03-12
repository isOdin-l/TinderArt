## Create Profile
```mermaid
sequenceDiagram
    Client->>Gateway:       POST /register
    Gateway->>Auth:         register(userData)

    Auth->>Profile:         gRPC - createUser(userData)

    Profile->>PostgreSQL:   insert user
    alt Already Exists
        PostgreSQL->>Profile: Error
        Profile->>Gateway: Error
        Gateway->>Client: Error
    else Insertion OK
        Profile-->>Auth:        gRPC - ok
        Auth->>PostgreSQL:      save refresh token
        Auth-->>Client:         access + refresh
    end
```

<br>

## Get Profile
```mermaid
sequenceDiagram
    participant Client
    participant Gateway     as API Gateway
    participant Auth        as Auth-Service
    participant Profile     as Profile-Service
    participant Cache       as Redis
    participant DB          as PostgreSQL+PostGIS

    Note over Client, Gateway: 1 - Запрос от Клиента
    Client->>Gateway: GET /profile

    Note over Gateway, Auth: 2 - Валидация по access_token
    Gateway->>Auth: validate(access_token)
    Auth-->>Gateway: validation result

    alt Unauthorized
        Gateway-->>Client: Unauthorized
    else Authorized
        Note over Gateway, DB: 3 - Получение профиля
        
        Gateway->>Profile: getProfile(user_id)
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

<br>

## Delete profile
```mermaid
sequenceDiagram
    participant Client          as Client
    participant Gateway         as API Gateway
    participant Auth            as Auth-Service
    participant Profile         as Profile-Service
    participant DB              as PostgreSQL
    participant S3              as RustFS
    participant Cache           as Redis

    Client->>Gateway: DELETE /profile
    Gateway->>Auth: validate(access_token)
    Auth-->>Gateway: validation result

    alt Unauthorized
        Gateway-->>Client: Unauthorized
    else Authorized
        Gateway->>Profile: deleteProfile(user_id)

        Profile->>DB: get photos URLs
        DB-->>Profile: photos

        loop delete photos
            Profile->>S3: delete object
        end

        Profile->>DB: DEL row by user_id
        DB-->>Profile: ok

        Profile->>Cache: DEL "pop_":{user_id}
        Profile->>Cache: DEL "stack_":{user_id}

        Profile-->>Gateway: ok
        Gateway-->>Client: HTTP-204 
    end
```

<br>

## Update preferences
```mermaid
sequenceDiagram
    participant Client      as Client
    participant Gateway     as API Gateway
    participant Auth        as Auth-Service
    participant Profile     as Profile-Service
    participant DB          as PostgreSQL
    participant Cache       as Redis

    Client->>Gateway: PATCH /preferences
    Gateway->>Auth: validate(access_token)
    Auth-->>Gateway: validation result

    alt Unauthorized
        Gateway-->>Client: Unauthorized
    else Authorized
        Gateway->>Profile: updatePreferences(data)
        Profile->>DB: UPDATE preferences
        DB->>Profile: Results
        alt Error
            Profile->>Gateway: Error
            Gateway->>Client: Error
        else OK
            Profile->>Cache: DEL "stack_":{user_id}
            Profile->>Cache: DEL "pop_":{user_id}

            Profile-->>Gateway: ok
            Gateway-->>Client: ok
        end
    end
```

<br>

## Generate Daily Stack
```mermaid
sequenceDiagram
    participant Cron
    participant Stack as DailyStack-Service
    participant DB as PostgreSQL+PostGIS
    participant Cache as Redis

    Note over Cron,Stack: 1 - Старт сервиса по Cron каждый день
    Cron->>Stack: trigger daily stack generation

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
    participant Auth        as Auth-Service
    participant Profile     as Profile-Service
    participant Cache       as Redis
    participant DB          as PostgreSQL+PostGIS

    Note over Client, Gateway: 1 - Запрос от Клиента
    Client->>Gateway: GET /profile

    Note over Gateway, Auth: 2 - Валидация по access_token
    Gateway->>Auth: validate(access_token)
    Auth-->>Gateway: validation result

    alt Unauthorized
        Gateway-->>Client: Unauthorized
    else Authorized
        Note over Gateway, DB: 3 - Получение профиля
        
        Gateway->>Profile: getStack(for user_id)
        Profile->>Cache: get("stack_"$user_id)

        alt Cache hit
            Cache-->>Profile: part of stack
            Profile->>Gateway: part of stack
            Gateway->>Client: part of stack
        else Cache miss
            Profile->>DB: Get profiles from DB
            alt Found
                DB-->>Profile: profiles
                Profile->>Cache: set("stack_"$user_id, profiles)

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
    participant Auth as Auth-Service
    participant Swipe as Swipe-Service
    participant DB as PostgreSQL
    participant Kafka
    participant Notify as Notify-Service
    participant External as ExtarnalAPIs

    Client->>Gateway: POST /swipe

    Gateway->>Auth: validate(access_token)
    Auth-->>Gateway: validation result

    alt Unauthorized
        Gateway-->>Client: 401 Unauthorized
    else Authorized
        Gateway->>Swipe: createSwipe(user1_id, user2_id)

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