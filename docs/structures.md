
## Database tables
### Profiles
| Field | Type |
|---|---|
| id | uuid |
| username | Text |
| name | Text |
| surname | Text |
| email | Text |
| password | Text |
| description | Text |
| location | GEOGRAPHY(Point, 4326) |
| created_at | TIMESTAMP|

<br>

### JWT_Tokens
| Field | Type |
|---|---|
| id - PK(profiles.id) | uuid |
| refresh_token | Text |

<br>

### Preferences
| Field | Type |
|---|---|
| profile_id - PK(profiles.id) | uuid |
| max_distance_meters | int |

<br>

### Fav_art_styles
| Field | Type |
|---|---|
| id | UUID |
| profile_id PK(profiles.id) | UUID |
| style | art_style_enum |
| created_at | time |

```
art_style_enum = 'realism', 'minimalism', 'futurism', 'anarchism', 'cubism', 'surrealism', 'impressionism', 'expressionism', 'constructivism', 'dadaism', 'photorealism', 'romanticism', 'cyberpunk'
````

### Arts
| Field | Type |
|---|---|
| id | uuid |
| profile_id | uuid |
| url | Text |
| created_at | time |

<br>

### Swipes
| Field | Type |
|---|---|
| userId_1 | uuid |
| userId_2 | uuid |
| desicion_1 | bool |
| desicion_2 | bool |
| created_at | time |

### Kafka topics
matches:
```
${userId_1}:${userId_2}
```

<br>

### Redis

```
Popular profile

Key: "profile:{profile_id}"
Value: user-data - JSON
```

<br>

```
User's stack

Key: "stack:{profile_id}"
Value: [ user-data, user-data, ... ] - List
```

<br>
    
### S3 storage
```
profile-photos: 

filename    UUID
```
