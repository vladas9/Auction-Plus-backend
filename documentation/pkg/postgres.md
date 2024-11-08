
# Documentation for `postgres.go`

### Package Declaration
`package db`

Defines the `db` package, responsible for establishing and managing connections to the PostgreSQL database.

---

### Imports
- **`database/sql`**: Provides interfaces to SQL or SQL-like databases.
- **`fmt`**: Used for formatting strings, specifically the connection string.
- **`os`**: Used to retrieve environment variables.
- **`github.com/joho/godotenv`**: Loads environment variables from a `.env` file.
- **`github.com/lib/pq`**: The PostgreSQL driver for `database/sql`.
- **`internal/utils`** (aliased as `u`): Provides utilities, including logging.

---

### Functions

#### `ConnectDB`
```go
func ConnectDB() (*sql.DB, error)
```
Establishes a connection to the PostgreSQL database.

- **Functionality**:
  - Loads environment variables from the `.env` file using `godotenv.Load()`.
  - Retrieves database configuration details (host, user, password, database name, port, and SSL mode) from environment variables.
  - Constructs the PostgreSQL connection string using the retrieved parameters.
  - Opens a connection to the PostgreSQL database using `sql.Open` with the constructed connection string.
  - Pings the database to verify the connection is active.

- **Returns**:
  - `*sql.DB`: A database object for performing queries.
  - `error`: An error if any step fails, including loading environment variables, opening the connection, or pinging the database.

- **Logging**:
  - Logs a success message if the connection is successfully established.

