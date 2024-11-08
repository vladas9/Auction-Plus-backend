### `auctionService.go`

Manages the auction-related business logic, providing functions to interact with `AuctionModel`.

#### Functions

- **`StartAuction(auction AuctionModel) error`**: Initiates a new auction, setting its status as active.
- **`EndAuction(id string) error`**: Finalizes the auction, updates its status, and identifies the winning bid.
- **`ExtendAuctionTime(id string, duration time.Duration) error`**: Adds extra time to an auction if the bidding threshold is met.
- **`GetAuctionDetails(id string) (*AuctionDetails, error)`**: Retrieves comprehensive details about a specific auction.

---

### `bidServices.go`

Handles bid-related services, managing the interaction between auctions and user bids.

#### Functions

- **`PlaceBid(bid BidModel) error`**: Validates and records a new bid for a specified auction.
- **`GetBidHistory(auctionId string) ([]BidModel, error)`**: Retrieves all bids placed on a specific auction.
- **`UpdateHighestBid(auctionId string, bid BidModel) error`**: Updates the current highest bid for an auction.
- **`ValidateBid(bid BidModel) (bool, error)`**: Ensures bid compliance with auction rules and minimum bid increments.

---

### `eventService.go`

Manages real-time events within the platform, especially for handling WebSocket connections in live auctions.

#### Functions

- **`BroadcastAuctionUpdate(auctionId string, event AuctionEvent) error`**: Broadcasts updates for a live auction to connected clients.
- **`HandleNewBidEvent(bid BidModel) error`**: Handles a new bid event, updating the highest bid and notifying participants.
- **`ConnectUserToAuction(auctionId string, userId string) error`**: Establishes a WebSocket connection for a user to participate in a live auction.
- **`DisconnectUserFromAuction(auctionId string, userId string) error`**: Removes a user from a live auction's WebSocket connections.

---

### `services.go`

Serves as the central registry and coordinator for initializing and accessing various service modules within the platform.

#### Functions

- **`InitServices()`**: Initializes instances of all major service components (e.g., auction, bid, user).
- **`GetAuctionService() *AuctionService`**: Provides access to the auction service instance.
- **`GetBidService() *BidService`**: Provides access to the bid service instance.
- **`GetUserService() *UserService`**: Provides access to the user service instance.

---

### `serviceUtils.go`

Contains utility functions for shared logic across services, supporting consistent handling and validation.

#### Functions

- **`ValidateAuctionStatus(auction AuctionModel) (bool, error)`**: Checks if an auction is active and eligible for new bids.
- **`CalculateBidIncrement(currentBid decimal.Decimal) decimal.Decimal`**: Calculates the required minimum bid increment based on the current bid amount.
- **`GenerateEventID() string`**: Creates a unique identifier for real-time events.
- **`LogServiceError(service string, err error)`**: Logs errors specific to each service for debugging.

---

### `userService.go`

Handles user-related business logic and data interactions, including user authentication and profile management.

#### Functions

- **`RegisterUser(user UserModel) error`**: Registers a new user on the platform.
- **`LoginUser(email, password string) (*UserModel, error)`**: Authenticates a user’s credentials and returns user data upon success.
- **`GetUserProfile(id string) (*UserProfile, error)`**: Retrieves a user’s profile details based on their unique ID.
- **`UpdateUserProfile(id string, updates UserProfile) error`**: Updates specific fields in a user’s profile.

---