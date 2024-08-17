# Post Article

A web app for managing article posts.

API Documentation can be accessed through [Postman](https://documenter.getpostman.com/view/29785588/2sA3s9BnCK)

## Available Features

-   Post new article
-   List all articles
-   Get an article
-   Edit an article
-   Delete an article

## Prerequisites

This project is built using:

### Backend

-   Go v1.19.13
-   MySQL 9.0.1-1.el9
-   Gin v1.10.0
-   Go MySQL Driver v1.8.1

### Frontend

-   TypeScript
-   React
-   Vite

## Installation

### Backend

1. Clone the project and change directory to backend/api.

    ```sh
    cd backend/api
    ```

2. Run these commands:

    ```sh
    # tidy dependencies
    go mod tidy

    # create env file from example
    cp .env.example .env
    ```

3. Change the values in `.env` to your own values. Set the `APP_ENV` value in `.env` to either `release` (for production use) or `debug` for development use.

4. Run the SQL files in `/backend/db` with your favorite DB connection driver to migrate the database and required tables.

5. Run the file `/backend/db/seed/seed.sql` with your favorite DB connection driver for seeding required data.

6. (optional) Run the file `/backend/db/seed/example.sql` with your favorite DB connection driver for seeding example data.

7. Serve the app with `go run .`

### Frontend

1. Clone the project and change directory to frontend.

    ```sh
    cd frontend
    ```

2. Install dependencies with npm

    ```sh
    npm install
    ```

3. Generate a new .env file

    ```sh
    cp .env.example .env
    ```

4. Change the VITE_BASE_URL value according to the backend `BASE_URL`

5. Start the server or build to static files

    ```sh
    npm run dev # for running the server
    # or
    npm run build # for building the project
    ```
