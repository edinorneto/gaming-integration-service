# Gaming Integration Service

A backend service built with Go to simulate a small gaming aggregation and session management service.

The project models a domain involving game providers, games, and player
game sessions, exposing these capabilities through a REST API.

> This is an independent portfolio project inspired by the gaming industry
> domain. It is not an official integration with Jungle Gaming or any
> third-party provider.

## Overview

The service currently provides a REST API for:

- listing games
- retrieving a game by ID
- creating games
- listing game providers
- creating game sessions
- listing game sessions

The current version uses in-memory storage. Data is lost when the
application is restarted.

PostgreSQL persistence and Docker-based development are planned for the
next development stages.

## Tech Stack

### Current

- Go
- `net/http`
- JSON
- In-memory storage

### Planned

- PostgreSQL
- Docker
- Docker Compose
- Automated tests

## Domain

The current domain contains three main entities.

### Provider

Represents a game provider available in the system.

Fields:

    ID
    Nome
    Active

### Game

Represents a game associated with a provider.

Fields:

    ID
    Nome
    ProviderID

### Game Session

Represents a game session started by a player.

Fields:

    ID
    PlayerID
    GameID
    Status

A game session can only be created when the referenced game exists.

New sessions are created with the initial status:

    active

## Sample Data

The application starts with a small in-memory dataset containing providers
and games.

The current sample data is intended for local development and testing.
It does not represent official integrations with the companies or games
used as sample data.

## API

### Games

#### List all games

    GET /games

Returns all games currently stored in memory.

#### Get a game by ID

    GET /game/{id}

Example:

    GET /game/2

Possible responses:

- `200 OK` when the game exists
- `400 Bad Request` when the ID is invalid
- `404 Not Found` when the game does not exist

#### Create a game

    POST /games

Content-Type:

    application/json

Example request:

    {
      "nome": "Dark Souls III",
      "provider_id": 1
    }

The provider referenced by `provider_id` must exist.

Successful creation returns:

    201 Created

Example response:

    {
      "id": 4,
      "nome": "Dark Souls III",
      "provider_id": 1
    }

Possible responses:

- `201 Created` when the game is successfully created
- `400 Bad Request` when the JSON is invalid
- `404 Not Found` when the provider does not exist

### Providers

#### List all providers

    GET /providers

Returns all providers currently stored in memory.

### Game Sessions

#### List all sessions

    GET /sessions

Returns all game sessions created during the current application execution.

#### Create a game session

    POST /sessions

Content-Type:

    application/json

Example request:

    {
      "player_id": 234,
      "game_id": 2
    }

The referenced game must exist in the current game catalog.

The server generates the session ID and sets its initial status to
`active`.

Example response:

    {
      "id": 1,
      "player_id": 234,
      "game_id": 2,
      "status": "active"
    }

Possible responses:

- `201 Created` when the session is successfully created
- `400 Bad Request` when the JSON is invalid
- `404 Not Found` when the referenced game does not exist

## HTTP Status Codes

The API currently uses the following status codes:

| Status | Description |
|---|---|
| `200 OK` | Request completed successfully |
| `201 Created` | Resource was successfully created |
| `400 Bad Request` | Invalid JSON or request parameter |
| `404 Not Found` | Requested resource or related entity does not exist |

Errors are returned using a standardized JSON structure:

    {
      "error": "Jogo não encontrado."
    }

## Project Structure

    gaming-integration-service/
    ├── domain/
    │   └── models.go
    ├── errors.go
    ├── handlers.go
    ├── main.go
    ├── go.mod
    └── README.md

### Responsibilities

#### `domain/`

Contains the core domain entities and their basic behavior:

- `Game`
- `Provider`
- `GameSession`

#### `handlers.go`

Contains the HTTP handlers and the current in-memory application data.

#### `errors.go`

Contains the helper used to return standardized JSON error responses.

#### `main.go`

Registers the API routes and starts the HTTP server.

## Current API Routes

    GET  /games
    GET  /games/{id}
    POST /games

    GET  /providers

    GET  /sessions
    POST /sessions

## Running Locally

Make sure Go is installed, then run:

    go run .

The API will be available at:

    http://localhost:8080

## Example Requests

### Get all games

    GET http://localhost:8080/games

### Get a specific game

    GET http://localhost:8080/games/2

### Get all providers

    GET http://localhost:8080/providers

### Get all sessions

    GET http://localhost:8080/sessions

### Create a game

    POST http://localhost:8080/games

Request body:

    {
      "nome": "Dark Souls III",
      "provider_id": 1
    }

### Create a game session

    POST http://localhost:8080/sessions

Request body:

    {
      "player_id": 234,
      "game_id": 2
    }

## Current State

The project currently provides:

- REST API built with Go
- HTTP routing with `net/http`
- JSON request and response handling
- In-memory game catalog
- In-memory provider catalog
- In-memory game session storage
- Game lookup by ID
- Game creation
- Provider lookup
- Game session creation
- Basic request validation
- Standardized JSON error responses
- HTTP status code handling
- Route parameters

## Planned Development

The next stages of the project are:

- PostgreSQL persistence
- Database schema and migrations
- Data access layer
- Docker and Docker Compose
- Automated tests
- Further API and domain refinements

## Goals

This project is being developed as a practical backend portfolio project
focused on learning and demonstrating:

- Go backend development
- REST API design
- HTTP and JSON handling
- Domain modeling
- Error handling
- HTTP status codes
- PostgreSQL persistence
- Docker-based development
- Backend architecture

## Roadmap

    Phase 1
    Go + Domain Modeling
            ↓
    Phase 2
    REST API + In-Memory Storage
            ↓
    Phase 3
    PostgreSQL Persistence
            ↓
    Phase 4
    Docker / Docker Compose
            ↓
    Phase 5
    Automated Tests + Refinements

## Disclaimer

This is an independent educational and portfolio project.

It is not affiliated with, sponsored by, or officially integrated with
Jungle Gaming, its partners, or any third-party gaming provider.