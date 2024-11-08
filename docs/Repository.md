---
tags:
  - project
  - programming
  - golang
  - backend
---
# Repository Overview

The project is structured around various files, each responsible for handling database queries and logic specific to different models. Below is a detailed overview:

- **repository.go** - Defines the `Store` type which encapsulates the database (`db`) and handles transactions (`tx`).
- **user.go** - Contains database queries related to the `User` [[model]].
- **auction.go** - Handles queries for the `Auction` [[model]].
- **bid.go** - Contains logic for interacting with the `Bid` [[model]].
- **item.go** - Manages data operations for the `Item` [[model]].
- **transaction.go** - Deals with queries involving the `Transaction` [[model]].
- **notification.go** - Defines queries for the `Notification` [[model]].
- **shipping.go** - Contains logic for database interactions related to shipping details.

Each file implements the database queries required for CRUD operations on the respective models.
# Logic Breakdown

## Initialization via Services

### File: `repository.go`

When any service is initialized, it calls the `NewStore` function, passing the database (`db`) as an argument. This is a key component as the `Store` type serves as the entry point for all database transactions.
```go
type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}
```
The `Store` type holds the database connection, allowing easy transaction management throughout the application.
## Handling Database Transactions

### File: `repository.go`

The application uses a `StoreTx` type, which wraps a database transaction (`sql.Tx`). This type is designed to simplify transaction handling, ensuring that all interactions are encapsulated within a single struct.
```go
type StoreTx struct {
	*sql.Tx
}
```
### Beginning a Transaction

The `BeginTx` function starts a new transaction by calling `db.Begin()` and returns a `StoreTx` that holds the transaction object.
```go
func (s *Store) BeginTx() (*StoreTx, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return nil, err
	}
	return &StoreTx{tx}, nil
}
```
### Executing Code Within a Transaction

The `WithTx` function encapsulates the core logic for managing transactions. It begins a new transaction via `BeginTx`, executes the provided function, and handles commit or rollback depending on whether the function returns an error. This function is widely used across various [[services]] when they make repository calls.
```go
func (s *Store) WithTx(fn func(stx *StoreTx) error) error {
	storeTx, err := s.BeginTx()
	defer storeTx.Rollback()
	if err != nil {
		return err
	}
	if err := fn(storeTx); err != nil {
		return err
	}
	return storeTx.Commit()
}
```
This method simplifies the transaction process by ensuring that if the operation fails, the transaction is rolled back, maintaining database consistency. If successful, the transaction is committed.

## Repository Methods

Each model in the application has its corresponding repository. The repository struct stores a reference to the active transaction (`tx`) and provides methods for interacting with the database.
```go
type <repo_name>Repo struct {
	tx *sql.Tx
}
```
### Initializing the Repository

For each model, a corresponding repository is initialized by the `Repo` method in `StoreTx`, which attaches the current transaction to the repository.
```go
func (s *StoreTx) <repo_name>Repo() *<repo_name>Repo {
	return &<repo_name>Repo{s.Tx}
}
```
# Queries

Each repository file implements standard CRUD operations, along with additional query logic specific to the model it handles.

### Standard CRUD Operations

#### Insert

Inserts a new record into the database for the corresponding model. This is a fundamental operation used for creating new entries.
```go
func (r *<repo_name>Repo) Insert(item *m.<repo_name>Model) (uuid.UUID, error)
```
#### GetById and GetAll

- **GetById** queries a record by its unique identifier (UUID) and retrieves the corresponding model.
	```go
func (r *<repo_name>Repo) GetById(id uuid.UUID) (*m.Model, error)
```
- **GetAll** retrieves a list of all records from the associated database table.
```go
func (r *<repo_name>Repo) GetAll() ([]*m.<repo_name>Model, error)
```
#### Update

Updates an existing record in the database based on its UUID. This operation is essential for modifying data.
```go
func (r *<repo_name>Repo) Update(item *m.<repo_name>Model) error
```
#### Remove

Deletes a record from the database by its UUID. This is used when a model needs to be removed from the data store
```go
func (r *<repo_name>Repo) Remove(id uuid.UUID) error
```
### Model-Specific Query 

In some models, there are additional specialized methods. 
#### User.go 
The `User` repository includes a `GetByEmail` method, which retrieves a user record based on their email address, functioning similarly to the `GetById` method.
```go
func (r *userRepo) GetByEmail(email string) (*m.UserModel, error)
```