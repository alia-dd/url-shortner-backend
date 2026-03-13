# url-shortner-backend
A simple **URL shortening service** built with **Go** and **PostgreSQL** — generates short aliases for long URLs and redirects requests to the original links.

---

## 🚀 Features

* Create short URLs for given long URLs
* Redirect short URLs to the original long URLs
* Alias generation and lookup
* Built with Go for performance and simplicity
* PostgreSQL for persistent storage

---

## 📦 Tech Stack

| Component | Technology                                 |
| --------- | ------------------------------------------ |
| Backend   | Go                                         |
| Database  | PostgreSQL                                 |
| ORM       | GORM (optional, if used)                   |
| Router    | Gin / net/http (your choice based on code) |

---

## 🛠️ Prerequisites

Before you begin, make sure you have the following installed:

* Go (1.18+)
* PostgreSQL
* Git

---

## ⚙️ Setup & Installation

1. **Clone the repository:**

   ```
   git clone https://github.com/alia-dd/url-shortner-backend.git
   cd url-shortner-backend
   ```

2. **Create a `.env` file** in the root with your PostgreSQL settings:

   ```
   DATABASE_URL=your_db_connection_string
   PORT=8080
   ```

3. **Install dependencies and build:**

   ```
   go mod download
   go build -o server
   ```

4. **Run the server:**

   ```
   ./server
   ```

   The API will start on http://localhost:8080 (or the port you set).
    or use the already hosted backedn <br> https://url-shortner-backend-36xa.onrender.com
---

## 📌 API Endpoints

### Create a Short URL

**POST** `/api/shorten`

* **Request Body:**

  ```json
  {
    "url": "https://example.com/very/long/url"
  }
  ```

* **Response:**

  ```json
  {
    "id": 1,
    "original_url": "https://example.com/very/long/url",
    "short_code": "abc123",
    "created_at": "2026-03-12T15:04:05Z"
  }
  ```

---

### Redirect to Original URL

**GET** `/{short_code}`

* Follows the alias and redirects (HTTP 301) to the original URL.

---

## 📦 Deployment

If you wish to use this backend server in your project you can deploy on rendar.
Just make sure your **PostgreSQL instance** is accessible from your backend.

---

## 📜 License

This project is licensed under **Apache‑2.0**.

---

## ⭐ Contribution

Feel free to improve this service — add authentication, analytics (click tracking), or improve the frontend interface!

---
