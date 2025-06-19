# 🐦 birdie-talk (Backend)

Welcome to the **birdie-talk** backend! This service powers the backend infrastructure for the birdie-talk application. It provides API endpoints to manage bird-related data and uses a MySQL database running inside a Docker container.

---

## Getting Started

Follow these steps to set up and run the backend locally.

### Prerequisites

* [Docker](https://www.docker.com/)
* [Make](https://www.gnu.org/software/make/)
* [curl](https://curl.se/) or an API tool like Postman or Insomnia

---

## Setup Instructions

1. **Start the Docker environment**

   This will spin up the MySQL container:

   ```bash
   docker-compose up -d
   ```

2. **Run database migrations**

   This will apply the latest database schema:

   ```bash
   make migrate-up
   ```

3. **Start the backend server**

   ```bash
   make run
   ```

---

## Seed Initial Data

Once the server is running, seed the initial data:

1. Open your preferred API tool (e.g. Postman, Insomnia), or use `curl`.

2. Make a `POST` request to:

   ```
   http://localhost:8080/api/v1/birds/initial
   ```

3. Set the request body to the contents of:

   ```
   data/initial-data.json
   ```

   Make sure to set the `Content-Type` header to `application/json`.

   Example using `curl`:

   ```bash
   curl -X POST http://localhost:8080/api/v1/birds/initial \
     -H "Content-Type: application/json" \
     -d @data/initial-data.json
   ```

---

## Project Structure

```bash
birdie-talk/
├── data/                   # Contains initial seed data
│   └── initial-data.json
├── internal/               # Go backend code
├── migrations/             # SQL migration files
├── docker-compose.yml      # Docker setup for MySQL
├── Makefile                # Useful commands for development
└── README.md               # You're here!
```

---

## Contact

For issues or feature requests, please open an [issue](https://github.com/arianaw15/birdie-talk/issues) on the repository.

Happy coding!
