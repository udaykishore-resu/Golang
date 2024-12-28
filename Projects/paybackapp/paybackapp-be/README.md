# PayBack App Backend

This is the backend service for the PayBack App, written in Go.

## Prerequisites

- Go 1.21 or higher
- MySQL database
- Git

## Setup

1. Create a `.env` file in the root directory with the following content:
```env
DB_USER=your_db_user
DB_PASS=your_db_password
DB_HOST=localhost:3306
DB_NAME=paybackapp
```

2. Set up the database:
```sql
CREATE DATABASE paybackapp;

USE paybackapp;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    points INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE restaurants (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    image_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE points_transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    restaurant_id INT,
    points INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (restaurant_id) REFERENCES restaurants(id)
);
```

3. Install dependencies:
```bash
go mod download
```

4. Run the server:
```bash
go run cmd/main.go
```

The server will start on http://localhost:8080

## API Endpoints

### Restaurants
- GET /api/restaurants - Get all restaurants
- POST /api/restaurants - Create a new restaurant
- GET /api/restaurants/:id - Get a specific restaurant

### Users
- POST /api/register - Register a new user
- POST /api/login - Login user

### Points
- GET /api/points/:userId - Get user points
- POST /api/points - Add points to user

## Note
This is a basic implementation. In a production environment, you would want to add:
- Password hashing
- JWT authentication
- Input validation
- Error handling middleware
- Logging
- Rate limiting
- API documentation
- Tests