# Services Overview

The project uses service files to handle business logic and application-specific operations related to different models. Each service file interacts with repositories to perform CRUD operations and handle business rules.

- **auctionService.go** - Manages auction-related logic, such as starting, ending, and extending auctions.
- **bidServices.go** - Handles logic associated with placing and validating bids.
- **eventService.go** - Manages real-time event handling, particularly for WebSocket connections during live auctions.
- **services.go** - Acts as the initializer and provider of instances for each service.
- **serviceUtils.go** - Contains utility functions shared across various services.
- **userService.go** - Manages user-related operations, including authentication and profile management.

Each file implements the logic required to interact with the repositories and perform necessary operations on the corresponding models.

---

# Logic Breakdown

## Initialization via Services

### auctionService.go

Manages business logic for auctions, including operations to start, end, and extend auctions.

```go
type AuctionService struct {
	repo *auctionRepo
}
```

#### Functions

- **StartAuction** - Initializes a new auction by setting it to active.
  ```go
  func (s *AuctionService) StartAuction(auction AuctionModel) error
  ```

- **EndAuction** - Finalizes an auction, marking it as completed and determining the winner.
  ```go
  func (s *AuctionService) EndAuction(id string) error
  ```

- **ExtendAuctionTime** - Extends the auction duration if a bid is placed within the threshold time.
  ```go
  func (s *AuctionService) ExtendAuctionTime(id string, duration time.Duration) error
  ```

- **GetAuctionDetails** - Retrieves detailed information for a specific auction, including associated bids.
  ```go
  func (s *AuctionService) GetAuctionDetails(id string) (*AuctionDetails, error)
  ```

---

### bidServices.go

Handles bid-related business logic, including placing bids and maintaining bid history.

```go
type BidService struct {
	repo *bidRepo
}
```

#### Functions

- **PlaceBid** - Validates and places a new bid in an auction.
  ```go
  func (s *BidService) PlaceBid(bid BidModel) error
  ```

- **GetBidHistory** - Retrieves a history of bids for a specific auction.
  ```go
  func (s *BidService) GetBidHistory(auctionId string) ([]BidModel, error)
  ```

- **UpdateHighestBid** - Updates the auction’s highest bid.
  ```go
  func (s *BidService) UpdateHighestBid(auctionId string, bid BidModel) error
  ```

- **ValidateBid** - Ensures that a bid meets auction requirements, such as minimum increments.
  ```go
  func (s *BidService) ValidateBid(bid BidModel) (bool, error)
  ```

---

### eventService.go

Manages real-time events for auctions, such as broadcasting updates to connected users via WebSockets.

```go
type EventService struct {
	repo *eventRepo
}
```

#### Functions

- **BroadcastAuctionUpdate** - Sends auction updates to all connected clients.
  ```go
  func (s *EventService) BroadcastAuctionUpdate(auctionId string, event AuctionEvent) error
  ```

- **HandleNewBidEvent** - Processes a new bid event, updating the auction and notifying clients.
  ```go
  func (s *EventService) HandleNewBidEvent(bid BidModel) error
  ```

- **ConnectUserToAuction** - Connects a user to an auction session for live updates.
  ```go
  func (s *EventService) ConnectUserToAuction(auctionId string, userId string) error
  ```

- **DisconnectUserFromAuction** - Disconnects a user from a live auction session.
  ```go
  func (s *EventService) DisconnectUserFromAuction(auctionId string, userId string) error
  ```

---

### services.go

Initializes and provides instances of each service for use across the application.

```go
type Services struct {
	Auction *AuctionService
	Bid     *BidService
	User    *UserService
}
```

#### Functions

- **InitServices** - Initializes all services with their respective repositories.
  ```go
  func InitServices() *Services
  ```

- **GetAuctionService** - Provides access to the auction service.
  ```go
  func (s *Services) GetAuctionService() *AuctionService
  ```

- **GetBidService** - Provides access to the bid service.
  ```go
  func (s *Services) GetBidService() *BidService
  ```

- **GetUserService** - Provides access to the user service.
  ```go
  func (s *Services) GetUserService() *UserService
  ```

---

### serviceUtils.go

Provides shared utility functions used across various services for validation and error handling.

#### Functions

- **ValidateAuctionStatus** - Checks if an auction is active and eligible for bids.
  ```go
  func ValidateAuctionStatus(auction AuctionModel) (bool, error)
  ```

- **CalculateBidIncrement** - Calculates the required increment for a bid based on current bid.
  ```go
  func CalculateBidIncrement(currentBid decimal.Decimal) decimal.Decimal
  ```

- **GenerateEventID** - Generates a unique identifier for an event.
  ```go
  func GenerateEventID() string
  ```

- **LogServiceError** - Logs an error associated with a specific service.
  ```go
  func LogServiceError(service string, err error)
  ```

---

### userService.go

Handles user-related operations, including registration, login, and profile updates.

```go
type UserService struct {
	repo *userRepo
}
```

#### Functions

- **RegisterUser** - Registers a new user account on the platform.
  ```go
  func (s *UserService) RegisterUser(user UserModel) error
  ```

- **LoginUser** - Authenticates a user and returns user details upon successful login.
  ```go
  func (s *UserService) LoginUser(email, password string) (*UserModel, error)
  ```

- **GetUserProfile** - Fetches a user’s profile information by their ID.
  ```go
  func (s *UserService) GetUserProfile(id string) (*UserProfile, error)
  ```

- **UpdateUserProfile** - Updates the specified fields in a user's profile.
  ```go
  func (s *UserService) UpdateUserProfile(id string, updates UserProfile) error
  ```