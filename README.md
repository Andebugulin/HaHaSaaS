# HaHaSaaS (Humor as a Service)

A modern web application that serves as a joke management system with a terminal-style interface. Built with React, Go, and PostgreSQL.

![HaHaSaaS Terminal Interface]()

## Tech Stack

### Frontend
- React 18+ with TypeScript
- Vite build tool
- Tailwind CSS for styling
- Lucide React for icons

### Backend
- Go (Gin web framework)
- PostgreSQL database
- RESTful API architecture

## Prerequisites

- Node.js 16+
- Go 1.19+
- PostgreSQL 14+
- npm or yarn

## Installation

1. Clone the repository:
```bash
git clone https://github.com/andebugulin/hahasaas.git
cd hahasaas
```

2. Set up the database:
```sql
-- Create the database
CREATE DATABASE hahasaas;

-- Create caetgories table
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL
);

-- Create jokes table
CREATE TABLE jokes (
    id SERIAL PRIMARY KEY,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create jokes_categories table (many-to-many relationship)
CREATE TABLE jokes_categories (
    joke_id INT REFERENCES jokes(id),
    category_id INT REFERENCES categories(id),
    PRIMARY KEY (joke_id, category_id)
);

-- Create joke_reactions table
CREATE TABLE joke_reactions (
    joke_id INT REFERENCES jokes(id),
    likes INT DEFAULT 0,
    dislikes INT DEFAULT 0,
    PRIMARY KEY (joke_id)
);
```

3. Configure environment variables; put it where go files are located
```bash
# Create .env file in the root directory
cp .env.example .env

# Update the following variables
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=hahasaas
```

4. Install frontend dependencies:
```bash
npm install
```

5. Install backend dependencies:
```bash
cd backend
go mod download
```

## Running the Application

1. Start the backend server:
```bash
cd backend/cmd/server
go run .
```

2. Start the frontend development server:
```bash
npm run dev
```

3. Access the application at `http://localhost:5173`

## Available Commands

The terminal interface supports the following commands:

- `help`: Display available commands
- `clear`: Clear the terminal screen
- `rndj`: Get a random joke
- `rndj --category|-c CATEGORY`: Get a random joke from a specific category
- `lsc`: List all categories
- `lsj`: List all jokes
- `lsj --category|-c CATEGORY`: List jokes in a specific category
- `catj --id ID`: Get a specific joke by ID
- `addc CATEGORY`: Add a new category
- `addj --category|-c CATEGORY JOKE`: Add a new joke to a category
- `addj --category|-c CATEGORY --id ID`: Add an existing joke to another category
- `good --id ID`: Like a joke
- `bad --id ID`: Dislike a joke

## API Endpoints

### Jokes
- `GET /api/joke/random`: Get a random joke
- `GET /api/joke/random/:category`: Get a random joke from a category
- `GET /api/joke/categories`: Get all categories
- `GET /api/joke/all`: Get all jokes
- `GET /api/joke/category/:category`: Get jokes by category
- `GET /api/joke/:id`: Get a specific joke
- `POST /api/joke/category`: Add a new category
- `POST /api/joke`: Add a new joke
- `POST /api/joke/category/:category/:id`: Add joke to category
- `POST /api/joke/:id/like`: Like a joke
- `POST /api/joke/:id/dislike`: Dislike a joke

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Powered by jokes lol