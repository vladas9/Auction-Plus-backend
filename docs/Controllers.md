# Controllers Overview

Each controller file handles HTTP request and response logic specific to a model, bridging the client requests with corresponding services.

- **auctionController.go** - Manages auction-related endpoints for starting, ending, and retrieving auction details.
- **bidController.go** - Handles HTTP routes for placing bids, retrieving bid history, and validating bids.
- **controllers.go** - Centralizes controller initialization and routing, connecting each controller to the server.
- **userController.go** - Manages user-related endpoints, including registration, login, and profile management.
- **wsController.go** - Manages WebSocket connections for real-time updates in auctions.

Each file implements HTTP handler functions, handling incoming requests and calling the respective service methods to process data and return responses.

---

# Controller Breakdown

### auctionController.go

Handles HTTP requests related to auctions.

```go
type AuctionController struct {
	service *AuctionService
}
```

#### Functions

- **StartAuction** - HTTP POST endpoint to start a new auction.
  ```go
  func (c *AuctionController) StartAuction(w http.ResponseWriter, r *http.Request)
  ```

- **EndAuction** - HTTP POST endpoint to end an active auction and declare the winner.
  ```go
  func (c *AuctionController) EndAuction(w http.ResponseWriter, r *http.Request)
  ```

- **ExtendAuction** - HTTP POST endpoint to extend the auction’s duration.
  ```go
  func (c *AuctionController) ExtendAuction(w http.ResponseWriter, r *http.Request)
  ```

- **GetAuctionDetails** - HTTP GET endpoint to fetch detailed auction information.
  ```go
  func (c *AuctionController) GetAuctionDetails(w http.ResponseWriter, r *http.Request)
  ```

---

### bidController.go

Handles HTTP requests related to bidding on auctions.

```go
type BidController struct {
	service *BidService
}
```

#### Functions

- **PlaceBid** - HTTP POST endpoint for placing a bid in an auction.
  ```go
  func (c *BidController) PlaceBid(w http.ResponseWriter, r *http.Request)
  ```

- **GetBidHistory** - HTTP GET endpoint to retrieve the bid history for a specific auction.
  ```go
  func (c *BidController) GetBidHistory(w http.ResponseWriter, r *http.Request)
  ```

- **ValidateBid** - HTTP POST endpoint to validate a new bid.
  ```go
  func (c *BidController) ValidateBid(w http.ResponseWriter, r *http.Request)
  ```

---

### controllers.go

Centralizes the initialization of all controllers and sets up routing for each one.

```go
type Controllers struct {
	Auction *AuctionController
	Bid     *BidController
	User    *UserController
}
```

#### Functions

- **NewControllers** - Initializes all controllers and assigns them to the `Controllers` struct.
  ```go
  func NewControllers() *Controllers
  ```

- **RegisterRoutes** - Registers each controller’s routes with the server.
  ```go
  func (c *Controllers) RegisterRoutes(router *http.ServeMux)
  ```

---

### userController.go

Handles HTTP requests related to user operations, such as registration and profile management.

```go
type UserController struct {
	service *UserService
}
```

#### Functions

- **RegisterUser** - HTTP POST endpoint for registering a new user.
  ```go
  func (c *UserController) RegisterUser(w http.ResponseWriter, r *http.Request)
  ```

- **LoginUser** - HTTP POST endpoint for logging in a user.
  ```go
  func (c *UserController) LoginUser(w http.ResponseWriter, r *http.Request)
  ```

- **GetUserProfile** - HTTP GET endpoint to fetch a user’s profile data.
  ```go
  func (c *UserController) GetUserProfile(w http.ResponseWriter, r *http.Request)
  ```

- **UpdateUserProfile** - HTTP PUT endpoint to update the user’s profile details.
  ```go
  func (c *UserController) UpdateUserProfile(w http.ResponseWriter, r *http.Request)
  ```

---

### wsController.go

Handles WebSocket connections, facilitating real-time interactions within live auctions.

```go
type WSController struct {
	eventService *EventService
}
```

#### Functions

- **HandleAuctionEvents** - WebSocket handler for managing live auction events.
  ```go
  func (c *WSController) HandleAuctionEvents(ws *websocket.Conn)
  ```

- **BroadcastUpdate** - Sends an update message to all connected clients.
  ```go
  func (c *WSController) BroadcastUpdate(event AuctionEvent)
  ```

Each controller connects HTTP or WebSocket requests to the appropriate service, ensuring data is processed according to business logic and returned in response to client requests.