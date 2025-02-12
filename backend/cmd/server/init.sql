-- Create categories table
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