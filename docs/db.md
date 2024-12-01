# Database Module Documentation

## Overview

The `db` module is responsible for establishing a connection to a PostgreSQL database. It leverages the `database/sql` package and the `lib/pq` driver for interfacing with PostgreSQL. It uses environment variables to retrieve database configuration and ensures that the connection is active and functional.

---

## Package: `postgres`

### Imports

- **`database/sql`**: Standard SQL package in Go.
- **`fmt`**: For formatted I/O.
- **`os`**: To retrieve environment variables.
- **`github.com/joho/godotenv`**: To load environment variables from a `.env` file.
- **`github.com/lib/pq`**: PostgreSQL driver for `database/sql`.
- **`github.com/vladas9/backend-practice/internal/utils`**: Utility package, e.g., for logging.

---

## Variables

### `DB`

```go
var DB *sql.DB
```

- A global variable that holds the database connection instance.

---

## Functions

### `ConnectDB`

```go
func ConnectDB() error
```

- **Purpose**: Establishes a connection to the PostgreSQL database.
- **Returns**:
  - `nil`: If the connection is successful.
  - `error`: If any step in the connection process fails.

#### Steps

1. **Load Environment Variables**:
   - Uses `godotenv.Load()` to load `.env` file variables.
   - Reads the following variables:
     - `DB_HOST`: Hostname or IP of the database server.
     - `DB_USER`: Username for authentication.
     - `DB_PASSWORD`: Password for the user.
     - `DB_NAME`: Database name.
     - `DB_PORT`: Port number the database is listening on.
     - `DB_SSLMODE`: SSL mode (e.g., `disable`, `require`).
2. **Construct Connection String**:

   - Creates a connection string using the format:
     ```
     host=<host> user=<user> password=<password> dbname=<dbname> port=<port> sslmode=<sslmode>
     ```

3. **Open Connection**:

   - Calls `sql.Open` with the PostgreSQL driver and connection string.

4. **Ping Database**:

   - Ensures the connection is active by calling `DB.Ping()`.

5. **Logging**:
   - Logs a success message on successful connection using `u.Logger`.

#### Example

```go
err := ConnectDB()
if err != nil {
    log.Fatalf("Failed to connect to the database: %v", err)
}
```

---

## Environment Variables

| Variable      | Description                      | Example Value          |
| ------------- | -------------------------------- | ---------------------- |
| `DB_HOST`     | Hostname of the database server. | `localhost`            |
| `DB_USER`     | Username for the database.       | `postgres`             |
| `DB_PASSWORD` | Password for the database user.  | `mysecretpassword`     |
| `DB_NAME`     | Name of the database.            | `exampledb`            |
| `DB_PORT`     | Port for the database server.    | `5432`                 |
| `DB_SSLMODE`  | SSL mode for the connection.     | `disable` or `require` |

---

## Example `.env` File

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=mysecretpassword
DB_NAME=exampledb
DB_PORT=5432
DB_SSLMODE=disable
```

---

## Logging

- **Success**: Logs `Successfully connected to the database`.
- **Error Handling**: Errors during connection are returned for external handling.

---

## Dependencies

- **PostgreSQL**: Make sure PostgreSQL is running and accessible.
- **Environment Configuration**: Ensure a `.env` file exists with valid configuration values.

---

## Notes

- Always handle sensitive information (e.g., credentials) securely.
- Ensure to close the database connection (`DB.Close()`) when the application terminates.
