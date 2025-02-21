# HaHaSaaS API Reference

## Introduction

Welcome to the HaHaSaaS API! This API provides a collection of endpoints to manage and interact with jokes, including retrieving random jokes, managing categories, and handling user interactions like likes and dislikes.

## Base URL

All endpoints are relative to:

```
http://localhost:5173
```

## Health Check

### Ping

```endpoint
GET /ping
```

Simple health check endpoint to verify the API is running.

#### Example Response

```json
{
  "message": "pong"
}
```

## HaHaSaaS

### Get Random Joke

```endpoint
GET /api/joke/random
```

Retrieve a random joke from all available jokes in the database.

#### Example Response

```json
{
  "joke": {
    "id": 8,
    "content": "test joke",
    "likes": 0,
    "dislikes": 0
  }
}
```

### Get Random Joke from Category

```endpoint
GET /api/joke/random/{category}
```

Retrieve a random joke from a specific category.

#### Parameters

| Name | Type | Description |
|------|------|-------------|
| category | string | The category name to get a random joke from |

#### Example Request

```
GET /api/joke/random/test
```

#### Example Response

```json
{
  "joke": {
    "id": 6,
    "content": "test joke",
    "likes": 0,
    "dislikes": 0
  }
}
```

### Get All Categories

```endpoint
GET /api/joke/categories
```

Retrieve a list of all available joke categories.

#### Example Response

```json
{
  "categories": [
    "test"
  ]
}
```

### Get All HaHaSaaS

```endpoint
GET /api/joke/all
```

Retrieve all jokes in the database.

#### Example Response

```json
{
  "jokes": [
    "test joke",
    "test joke",
    // ... more jokes
  ]
}
```

### Get HaHaSaaS by Category

```endpoint
GET /api/joke/category/{category}
```

Retrieve all jokes for a specific category.

#### Parameters

| Name | Type | Description |
|------|------|-------------|
| category | string | The category name to get jokes from |

#### Example Response

```json
{
  "jokes": [
    "test joke",
    "test joke",
    // ... more jokes
  ]
}
```

### Get Joke by ID

```endpoint
GET /api/joke/{id}
```

Retrieve a specific joke by its ID.

#### Parameters

| Name | Type | Description |
|------|------|-------------|
| id | integer | The ID of the joke to retrieve |

#### Example Response

```json
{
  "joke": {
    "id": 14,
    "content": "test joke",
    "likes": 0,
    "dislikes": 0
  }
}
```

## Categories Management

### Add New Category

```endpoint
POST /api/joke/category
```

Add a new category for jokes.

#### Request Body

```json
{
  "name": "test-2"
}
```

#### Example Response

```json
{
  "message": "Category added successfully"
}
```

### Add New Joke

```endpoint
POST /api/joke
```

Add a new joke to a specific category.

#### Request Body

```json
{
  "Category": "test-2",
  "Joke": "dad joke :D)"
}
```

#### Example Response

```json
{
  "message": "Joke added successfully"
}
```

## User Interactions

### Like a Joke

```endpoint
POST /api/joke/like/{id}
```

Give a like to a specific joke.

#### Parameters

| Name | Type | Description |
|------|------|-------------|
| id | integer | The ID of the joke to like |

#### Example Response

```json
{
  "message": "Joke liked successfully"
}
```

### Dislike a Joke

```endpoint
POST /api/joke/dislike/{id}
```

Give a dislike to a specific joke.

#### Parameters

| Name | Type | Description |
|------|------|-------------|
| id | integer | The ID of the joke to dislike |

#### Example Response

```json
{
  "message": "Joke disliked successfully"
}
```

## Error Handling

The API uses conventional HTTP response codes to indicate the success or failure of requests:

- 200: Success
- 400: Bad Request
- 404: Not Found
- 500: Internal Server Error

## Rate Limiting

Currently, there are no rate limits implemented on the API endpoints.

## Changelog

### Version 1.0.0
- Initial release of the HaHaSaaS API
- Added support for joke categories
- Implemented like/dislike functionality
- Added random joke retrieval