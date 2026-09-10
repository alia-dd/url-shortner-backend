# url-shortner-backend

![Go](https://img.shields.io/badge/Go-1.18%2B-00ADD8?logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Router-Gin-00ADD8)
![PostgreSQL](https://img.shields.io/badge/Database-PostgreSQL-4169E1?logo=postgresql&logoColor=white)
![GORM](https://img.shields.io/badge/ORM-GORM-informational)
![Deployed on Render](https://img.shields.io/badge/Deployed%20on-Render-46E3B7)
![License](https://img.shields.io/badge/License-Apache--2.0-blue)

A simple URL shortening service built with Go and PostgreSQL. Generates short aliases for long URLs and redirects requests to the original links.

## Features

* Create short URLs for given long URLs
* Redirect short URLs to the original long URLs
* Deduplication — shortening the same long URL twice returns the existing alias
* Built with Go and Gin for performance and simplicity
* PostgreSQL for persistent storage, via GORM

## Tech Stack

| Component | Technology  |
| --------- | ----------- |
| Backend   | Go          |
| Router    | Gin         |
| ORM       | GORM        |
| Database  | PostgreSQL  |

## Prerequisites

Before you begin, make sure you have the following installed:

* Go (1.18+)
* PostgreSQL
* Git

## Setup and Installation

1. Clone the repository:

   ```
   git clone https://github.com/alia-dd/url-shortner-backend.git
   cd url-shortner-backend
   ```

2. Create a `.env` file in the root with your PostgreSQL connection string:

   ```
   DATABASE_URL=your_db_connection_string
   PORT=8000
   ```

3. Install dependencies and build:

   ```
   go mod download
   go build -o server
   ```

4. Run the server:

   ```
   ./server
   ```

   The API starts on http://localhost:8000 (or the port set in `PORT`).

   Or use the already hosted backend: https://url-shortner-backend-36xa.onrender.com

## API Endpoints

### Create a Short URL

**POST** `/shorten`

Request Body:

```json
{
  "url": "https://example.com/very/long/url"
}
```

Response:

```json
{
  "alias": "abc123"
}
```

If the long URL has already been shortened, the existing alias is returned instead of creating a duplicate.

### Redirect to Original URL

**GET** `/{alias}`

Follows the alias and redirects (HTTP 301) to the original URL. Returns 404 if the alias doesn't exist.

## Deployment

If you wish to use this backend server in your project, you can deploy it on Render. Just make sure `DATABASE_URL` and `PORT` are set in the service's environment variables and that your PostgreSQL instance is accessible from your backend.

## License

This project is licensed under Apache-2.0.

## Contribution

Feel free to improve this service — add authentication, analytics (click tracking), or improve the frontend interface.