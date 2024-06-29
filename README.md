## Getting Started

### Prerequisites

- Go (version 1.16 or later)
- Docker and Docker Compose
- MongoDB

### Installation

1. **Clone the repository**

   ```sh
   git clone https://github.com/yourgithubusername/searchstreet.git
   cd searchstreet
   ```

2. **Environment Setup**

   Copy the `.env` file provided in the root directory to configure the application settings.

   ```properties
   PORT=:8080
   APP_ENV=local
   DB_HOST=localhost
   MANAGER_PORT=8081
   DB_PORT=27017
   DB_USERNAME=root
   DB_PASS=example
   DB_COLLECTION=streets
   DB_DATABASENAME=searchstreet
   ```

3. **Running with Docker Compose**

   Use Docker Compose to start the MongoDB instance and the application.

   ```sh
   docker-compose up -d
   ```

4. **Building and Running Locally**

   If you prefer to run the application locally, use the Makefile commands.

   ```sh
   make run
   ```

   This command builds the application and starts the server with the environment variables specified in the `.env` file.

### Usage

Once the application is running, it will start listening for HTTP requests on the port specified in the `.env` file (`8080` by default).

## Development

### Structure

- `bin/`: Contains the compiled server binary.
- `data/`: MongoDB data files.
- `internal/`: Internal application code including services and business logic.
- `types/`: Go type definitions used across the application.

### Key Files

- `main.go`: The entry point of the application.
- `server.go`: HTTP server setup and route definitions.
- `docker-compose.yml`: Docker Compose configuration for running the application and MongoDB.
- `Makefile`: Contains commands for building and running the application.

### Dependencies

The project uses several external Go modules for its functionality, including:

- `github.com/joho/godotenv`: For loading environment variables from the `.env` file.
- `go.mongodb.org/mongo-driver`: MongoDB driver for Go.

Dependencies are managed via `go.mod` and `go.sum`.

