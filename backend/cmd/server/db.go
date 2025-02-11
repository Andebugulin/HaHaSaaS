package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// Data base schema
//-- Create categories table
//
// CREATE TABLE categories (
//     id SERIAL PRIMARY KEY,
//     name VARCHAR(100) UNIQUE NOT NULL
// );

// -- Create jokes table
// CREATE TABLE jokes (
//     id SERIAL PRIMARY KEY,
//     content TEXT NOT NULL,
//     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// );

// -- Create jokes_categories table (many-to-many relationship)
// CREATE TABLE jokes_categories (
//     joke_id INT REFERENCES jokes(id),
//     category_id INT REFERENCES categories(id),
//     PRIMARY KEY (joke_id, category_id)
// );

// -- Create joke_reactions table
// CREATE TABLE joke_reactions (
//     joke_id INT REFERENCES jokes(id),
//     likes INT DEFAULT 0,
//     dislikes INT DEFAULT 0,
//     PRIMARY KEY (joke_id)
// );

// Declare db as a package-level variable
var dbConn *sql.DB  // Changed name to dbConn to avoid confusion

type Joke struct {
    ID       int    `json:"id"`
    Content  string `json:"content"`
    Likes    int    `json:"likes"`
    Dislikes int    `json:"dislikes"`
}

func InitDB() error {
    connStr := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    var err error
    dbConn, err = sql.Open("postgres", connStr)
    if err != nil {
        return fmt.Errorf("error opening database: %v", err)
    }

    err = dbConn.Ping()
    if err != nil {
        return fmt.Errorf("error connecting to the database: %v", err)
    }

    return nil
}

// fetch a random joke
func fetchRandomJoke() (Joke, error) {
    var joke Joke
    err := dbConn.QueryRow(`
        SELECT j.id, j.content, COALESCE(jr.likes, 0), COALESCE(jr.dislikes, 0)
        FROM jokes j
        LEFT JOIN joke_reactions jr ON j.id = jr.joke_id
        ORDER BY RANDOM()
        LIMIT 1`).Scan(&joke.ID, &joke.Content, &joke.Likes, &joke.Dislikes)
    if err != nil {
        return Joke{}, fmt.Errorf("error fetching joke: %v", err)
    }
    return joke, nil
}

// fetching a random joke from a specific category
func fetchRandomJokeByCategory(category string) (Joke, error) {
    var joke Joke
    err := dbConn.QueryRow(`
        SELECT j.id, j.content, COALESCE(jr.likes, 0), COALESCE(jr.dislikes, 0)
        FROM jokes j
        LEFT JOIN joke_reactions jr ON j.id = jr.joke_id
        JOIN jokes_categories jc ON j.id = jc.joke_id
        JOIN categories c ON jc.category_id = c.id
        WHERE c.name = $1
        ORDER BY RANDOM()
        LIMIT 1`, category).Scan(&joke.ID, &joke.Content, &joke.Likes, &joke.Dislikes)
    if err != nil {
        return Joke{}, fmt.Errorf("error fetching joke: %v", err)
    }
    return joke, nil
}

// retrieve all categories
func fetchCategories() ([]string, error) {
	rows, err := dbConn.Query("SELECT name FROM categories")
	if err != nil {
		return nil, fmt.Errorf("error fetching categories: %v", err)
	}
	defer rows.Close()

	var categories []string
	for rows.Next() {
		var category string
		err := rows.Scan(&category)
		if err != nil {
			return nil, fmt.Errorf("error scanning categories: %v", err)
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// listing all jokes
func fetchAllJokes() ([]string, error) {
	rows, err := dbConn.Query("SELECT content FROM jokes")
	if err != nil {
		return nil, fmt.Errorf("error fetching jokes: %v", err)
	}
	defer rows.Close()

	var jokes []string
	for rows.Next() {
		var joke string
		err := rows.Scan(&joke)
		if err != nil {
			return nil, fmt.Errorf("error scanning jokes: %v", err)
		}
		jokes = append(jokes, joke)
	}

	return jokes, nil
}

// list all jokes from category
func fetchJokesByCategory(category string) ([]string, error) {
	rows, err := dbConn.Query(
		`SELECT content
		FROM jokes
		JOIN jokes_categories ON jokes.id = jokes_categories.joke_id
		JOIN categories ON jokes_categories.category_id = categories.id
		WHERE categories.name = $1`,
		category,
	)
	if err != nil {
		return nil, fmt.Errorf("error fetching jokes: %v", err)
	}
	defer rows.Close()

	var jokes []string
	for rows.Next(){
		var joke string
		err := rows.Scan(&joke)
		if err != nil {
			return nil, fmt.Errorf("error scanning jokes: %v", err)
		}
		jokes = append(jokes, joke)
	}

	return jokes, nil
}


// fetch a joke by id
func fetchJokeByID(id int) (Joke, error) {
    var joke Joke
    err := dbConn.QueryRow(`
        SELECT j.id, j.content, COALESCE(jr.likes, 0), COALESCE(jr.dislikes, 0)
        FROM jokes j
        LEFT JOIN joke_reactions jr ON j.id = jr.joke_id
        WHERE j.id = $1`, id).Scan(&joke.ID, &joke.Content, &joke.Likes, &joke.Dislikes)
    if err != nil {
        return Joke{}, fmt.Errorf("error fetching joke: %v", err)
    }
    return joke, nil
}

// adding a new category
func addCategory(name string) error {
	_, err := dbConn.Exec("INSERT INTO categories (name) VALUES ($1)", name)
	if err != nil {
		return fmt.Errorf("error adding category: %v", err)
	}

	return nil
}

// add a new joke (category is obligatory)
func addJoke(content, category string) error {
	tx, err := dbConn.Begin()
	if err != nil {
		return fmt.Errorf("error starting transaction: %v", err)
	}

	var jokeID int
	err = tx.QueryRow("INSERT INTO jokes (content) VALUES ($1) RETURNING id", content).Scan(&jokeID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error adding joke: %v", err)
	}

	var categoryID int
	err = tx.QueryRow("SELECT id FROM categories WHERE name = $1", category).Scan(&categoryID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error fetching category: %v", err)
	}

	_, err = tx.Exec("INSERT INTO jokes_categories (joke_id, category_id) VALUES ($1, $2)", jokeID, categoryID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error adding joke to category: %v", err)
	}

	_, err = tx.Exec("INSERT INTO joke_reactions (joke_id) VALUES ($1)", jokeID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error adding joke reactions: %v", err)
	}

	tx.Commit()
	return nil
}

// liking a joke with id
func likeJoke(id int) error {
	_, err := dbConn.Exec("UPDATE joke_reactions SET likes = likes + 1 WHERE joke_id = $1", id)
	if err != nil {
		return fmt.Errorf("error liking joke: %v", err)
	}

	return nil
}

// dislike a joke with id
func dislikeJoke(id int) error {
	_, err := dbConn.Exec("UPDATE joke_reactions SET dislikes = dislikes + 1 WHERE joke_id = $1", id)
	if err != nil {
		return fmt.Errorf("error disliking joke: %v", err)
	}

	return nil
}
// test all functions in db.go besides InitDB
func testDB() {
    fmt.Println("\n=== Starting Database Tests ===\n")

    // fetch a random joke
    fmt.Println("1. Testing fetchRandomJoke()...")
    joke, err := fetchRandomJoke()
    if err != nil {
        fmt.Printf("❌ Error: %v\n", err)
    } else {
        fmt.Printf("✓ Got random joke: %+v\n", joke)
    }

    // fetch a random joke by category
    fmt.Println("\n2. Testing fetchRandomJokeByCategory('test')...")
    joke, err = fetchRandomJokeByCategory("test")
    if err != nil {
        fmt.Printf("❌ Error: %v\n", err)
    } else {
        fmt.Printf("✓ Got random test joke: %+v\n", joke)
    }

    // fetch all categories
    fmt.Println("\n3. Testing fetchCategories()...")
    categories, err := fetchCategories()
    if err != nil {
        fmt.Printf("❌ Error: %v\n", err)
    } else {
        fmt.Printf("✓ Got categories: %v\n", categories)
    }

    // fetch all jokes
    fmt.Println("\n4. Testing fetchAllJokes()...")
    jokes, err := fetchAllJokes()
    if err != nil {
        fmt.Printf("❌ Error: %v\n", err)
    } else {
        fmt.Printf("✓ Got %d jokes\n", len(jokes))
    }

    // fetch jokes by category
    fmt.Println("\n5. Testing fetchJokesByCategory('dad')...")
    jokes, err = fetchJokesByCategory("dad")
    if err != nil {
        fmt.Printf("❌ Error: %v\n", err)
    } else {
        fmt.Printf("✓ Got %d dad jokes\n", len(jokes))
    }

    // fetch joke by id
    fmt.Println("\n6. Testing fetchJokeByID(1)...")
    joke, err = fetchJokeByID(1)
    if err != nil {
        fmt.Printf("❌ Error: %v\n", err)
    } else {
        fmt.Printf("✓ Got joke: %+v\n", joke)
    }

    // add a new category
    fmt.Println("\n7. Testing addCategory('test')...")
    err = addCategory("test")
    if err != nil {
        fmt.Printf("❌ Error: %v\n", err)
    } else {
        fmt.Println("✓ Category added successfully")
    }

    // add a new joke
    fmt.Println("\n8. Testing addJoke('test joke', 'test')...")
    err = addJoke("test joke", "test")
    if err != nil {
        fmt.Printf("❌ Error: %v\n", err)
    } else {
        fmt.Println("✓ Joke added successfully")
    }

    // like a joke
    fmt.Println("\n9. Testing likeJoke(1)...")
    err = likeJoke(1)
    if err != nil {
        fmt.Printf("❌ Error: %v\n", err)
    } else {
        fmt.Println("✓ Joke liked successfully")
    }

    // dislike a joke
    fmt.Println("\n10. Testing dislikeJoke(1)...")
    err = dislikeJoke(1)
    if err != nil {
        fmt.Printf("❌ Error: %v\n", err)
    } else {
        fmt.Println("✓ Joke disliked successfully")
    }

    fmt.Println("\n=== Database Tests Completed ===")
}