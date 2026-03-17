# 🧠 TodoList Microservices Platform

Production-style **microservices TodoList system** built with Go.
The system demonstrates **service separation, asynchronous communication, caching, and event-driven architecture**.

This project is designed to practice **backend system design concepts** such as:

* Microservices architecture
* gRPC communication
* Event-driven systems
* Redis caching
* RabbitMQ job queues
* Kafka event streaming
* Worker processes
* JWT authentication

---

# 🏗 Architecture

System consists of several independent services.

```
                ┌──────────────┐
                │  API Gateway │
                └──────┬───────┘
                       │
        ┌──────────────┼──────────────┐
        │              │              │
   ┌────▼─────┐  ┌─────▼─────┐  ┌─────▼─────┐
   │  Auth    │  │   User    │  │   Notes   │
   │ Service  │  │ Service   │  │ Service   │
   └────┬─────┘  └─────┬─────┘  └─────┬─────┘
        │               │              │
        │               │              │
   RabbitMQ          gRPC           Redis
        │
        ▼
   ┌─────────────┐
   │ Email       │
   │ Service     │
   └──────┬──────┘
          │
        Kafka
          │
          ▼
     ┌───────────┐
     │ Task      │
     │ Service   │
     └───────────┘
```

---

# ⚙️ Services

## 🔐 Auth Service

Responsible for authentication.

Features:

* Register user
* Login user
* Generate JWT access tokens
* Publish welcome email jobs

Endpoints:

```
POST /register
POST /login
```

Example request:

```json
{
  "name": "Azizbek",
  "email": "aziz@mail.com",
  "password": "12345678"
}
```

On successful registration:

```
RabbitMQ → email_auth_queue
```

Payload:

```json
{
  "name": "Azizbek",
  "email": "aziz@mail.com"
}
```

---

# 👤 User Service

Handles **user profile management**.

Features:

* Create user
* Get user by id
* Update profile
* Delete user

Endpoints:

```
POST /users
GET /users/{id}
PATCH /users/{id}
DELETE /users/{id}
```

Database:

```
users
```

| column        | type   |
| ------------- | ------ |
| id            | PK     |
| name          | string |
| email         | string |
| password_hash | string |

Indexes:

```
PRIMARY KEY (id)
UNIQUE (email)
```

Redis Cache:

```
user:<id>
```

Pattern:

```
Cache Aside
```

---

# 📝 Notes Service

Handles **user notes**.

Features:

* Create note
* Get user notes
* Update note
* Delete note

Endpoints:

```
POST /notes
GET /notes/{user_id}
GET /notes/item/{note_id}
PATCH /notes/{note_id}
DELETE /notes/{note_id}
```

Example:

```json
{
  "user_id": 12,
  "title": "Study",
  "text": "Learn Go concurrency"
}
```

Database:

```
notes
```

Indexes:

```
PRIMARY KEY (id)
INDEX (user_id)
```

Redis cache:

```
notes:list:<user_id>
notes:item:<user_id>:<note_id>
```

Strategy:

```
Cache invalidation on CRUD
```

---

# ⏰ Task Service

Responsible for **task reminders**.

Features:

* Task CRUD
* Reminder scheduling
* Kafka event publishing

Endpoints:

```
POST /tasks
GET /tasks/{user_id}
PATCH /tasks/{task_id}
DELETE /tasks/{task_id}
```

Example:

```json
{
  "user_id": 12,
  "task": "Submit documents",
  "remind_at": "2026-02-12T18:30:00Z"
}
```

Database:

```
tasks
```

| column    | type      |
| --------- | --------- |
| id        | PK        |
| user_id   | FK        |
| task      | string    |
| remind_at | timestamp |
| notified  | boolean   |

Indexes:

```
INDEX(user_id)
INDEX(notified, remind_at)
```

Worker logic:

```
SELECT *
FROM tasks
WHERE remind_at <= NOW()
AND notified = false
```

Worker interval:

```
every 5 seconds
```

After event sent:

```
UPDATE tasks SET notified = true
```

Kafka event:

Topic:

```
task_event
```

Payload:

```json
{
  "user_id": 12,
  "task": "Submit documents"
}
```

---

# 📬 Email Service

Handles email sending.

Consumers:

### RabbitMQ Consumer

Receives:

```
welcome_email
```

Purpose:

```
Send welcome email
```

---

### Kafka Consumer

Receives:

```
task reminder event
```

Flow:

```
Kafka → Email Service
          │
          ▼
gRPC → User Service
          │
          ▼
Fetch user email
          │
          ▼
Send reminder email
```

Worker pool used for sending emails.

---

# 📡 Event System

### RabbitMQ

Used for **background jobs**.

Queue:

```
email_auth_queue
```

Used for:

```
Welcome emails
```

---

### Kafka

Used for **system events**.

Topic:

```
task_event
```

Used for:

```
Task reminder events
```

---

# ⚡ Redis Caching Strategy

Global key format:

```
<service>:<entity>:<type>:<identifier>
```

Examples:

```
user:12
notes:list:12
notes:item:12:45
tasks:list:12
tasks:item:12:55
notifs:list:12
```

---

# 🔐 Authentication

Authentication uses **JWT access tokens**.

Algorithm:

```
HS256
```

TTL:

```
15 minutes
```

JWT payload:

```json
{
  "sub": "user_id"
}
```

Claims:

```
iat
exp
```

---

# ⏱ Time Handling

All times stored in **UTC**.

User timezone converted to UTC before storing.

Example:

```
User time → UTC → Database
```

Worker checks reminders every:

```
5 seconds
```

---

# 🧪 Technologies Used

* Go
* PostgreSQL / MySQL
* Redis
* RabbitMQ
* Kafka
* gRPC
* Docker
* JWT

---

# 🎯 Purpose of the Project

This project was built to practice **backend engineering concepts**:

* Distributed systems
* Microservice communication
* Event driven architecture
* Background workers
* System scalability
* Production-style backend design

---

