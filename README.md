# URL Shortener

A full-stack URL shortener built with Go, React, PostgreSQL, and Supabase. The application allows users to create short URLs, redirect using short codes, update existing mappings, delete mappings, and retrieve URL statistics.

## Features

- Create short URLs from long URLs
- Redirect users using short codes
- Update existing URL mappings
- Delete URL mappings
- Retrieve URL statistics

---

## Tech Stack

### Frontend

- React
- Vite
- Tailwind CSS
- JavaScript

### Backend

- Go
- Chi Router
- pgx PostgreSQL Driver

### Database

- PostgreSQL
- Supabase

### Deployment

- Render (Backend)
- Vercel (Frontend)

---



## Architecture

```text
┌─────────────┐      REST API     ┌─────────────┐      PostgreSQL      ┌─────────────┐
│ React (UI)  │ ────────────────► │ Go Backend  │ ───────────────────► │  Supabase   │
└─────────────┘                   └─────────────┘                      └─────────────┘
```

### URL Creation Flow

```text
User enters URL
        │
        ▼
React Frontend
        │
POST /shorten
        │
        ▼
Go Backend
        │
Generate Short Code
        │
Store in PostgreSQL
        │
        ▼
Return Short Code
        │
        ▼
Display Short URL
```

### Redirect Flow

```text
User visits /{shortcode}
        │
        ▼
Go Backend
        │
Lookup URL in Database
        │
        ▼
HTTP Redirect
        │
        ▼
Original Website
```

---

## Project Structure

```text
url_shortener/
│
├── backend/
│   ├── main.go
│   ├── routes.go
│   ├── handlers.go
│   ├── database.go
│   ├── config.go
│   ├── models.go
│   ├── utils.go
│   ├── go.mod
│   └── go.sum
│
├── client/
│   ├── src/
│   │   ├── App.jsx
│   │   ├── main.jsx
│   ├── package.json
│   ├── vite.config.js
│   └── index.html
│
└── README.md
```

---

## Backend Responsibilities

### main.go

Application entry point.

Responsibilities:

- Load configuration
- Connect to database
- Initialize routes
- Start HTTP server

### config.go

Handles application configuration.

Responsibilities:

- Read environment variables
- Provide configuration values

### database.go

Responsible for database connection and queries.

Responsibilities:

- Establish PostgreSQL connection
- Store global database instance
- Execute database operations

### handlers.go

Contains HTTP request handlers.

Responsibilities:

- Parse incoming requests
- Validate data
- Call database functions
- Return JSON responses

### routes.go

Responsible for API route registration.

Responsibilities:

- Configure router
- Register endpoints
- Configure middleware

### utils.go

Contains helper functions.

Responsibilities:

- Generate random short codes
- Reusable utility logic


---

## Environment Variables

### Backend

```env
DATABASE_URL=your_postgresql_connection_string
PORT=8000
```


---

## Running Locally

### Clone Repository

```bash
git clone https://github.com/Sri-Varshith/url_shortener.git
cd url_shortener
```

### Backend

```bash
cd backend

go mod download

go run .
```

Server starts on:

```text
http://localhost:8000
```

### Frontend

```bash
cd client

npm install

npm run dev
```

Frontend starts on:

```text
http://localhost:5173
```

---


## Future Improvements

- Custom short codes
- URL analytics dashboard
- Authentication
- Rate limiting
- QR code generation
- Custom domains
- Expiring URLs
- Improved statistics and tracking
