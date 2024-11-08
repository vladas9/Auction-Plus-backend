
### `auction.go`

Handles the `Auction` model, including fields and CRUD repository functions for auction operations.

#### Types

- **`Auction`**: Represents an auction, including details like seller, item, start price, current bid, and status.
- **Fields**:
  - `ID`, `SellerId`, `ItemId`, `StartPrice`, `CurrentBid`, `MaxBidderId`, `BidCount`, `StartTime`, `EndTime`, `ExtraTimeDuration`, `ExtraTimeThreshold`, `ExtraTimeEnabled`, `Status`.

#### Functions

- **`CreateAuction()`**: Adds a new auction record to the database.
- **`GetAuctionById(id string)`**: Retrieves an auction by its unique ID.
- **`UpdateAuction(id string, updates Auction)`**: Updates specified fields in an auction.
- **`DeleteAuction(id string)`**: Deletes an auction record from the database.

---

### `bid.go`

Defines the `Bid` model and repository functions for managing bids associated with auctions.

#### Types

- **`Bid`**: Represents a bid within an auction.
- **Fields**:
  - `ID`, `AuctionId`, `BidderId`, `Amount`, `Timestamp`.

#### Functions

- **`CreateBid()`**: Adds a new bid to the database.
- **`GetBidsByAuctionId(auctionId string)`**: Retrieves all bids associated with a specified auction ID.
- **`UpdateBid(id string, updates Bid)`**: Updates specific fields in a bid record.
- **`DeleteBid(id string)`**: Deletes a bid from the database.

---

### `item.go`

Handles the `Item` model, including fields and CRUD functions for items associated with auctions.

#### Types

- **`Item`**: Represents an item in an auction.
- **Fields**:
  - `ID`, `Name`, `Description`, `Category`, `Condition`, `StartingPrice`.

#### Functions

- **`CreateItem()`**: Inserts a new item record into the database.
- **`GetItemById(id string)`**: Retrieves an item by its unique ID.
- **`UpdateItem(id string, updates Item)`**: Updates specified fields in an item.
- **`DeleteItem(id string)`**: Deletes an item record from the database.

---

### `models.go`

Includes shared types, such as `AuctionDetails`, which encapsulate auction-related data.

#### Types

- **`AuctionDetails`**: Extends `Auction` with additional information like items and bidders.
- **Fields**:
  - `Auction`, `Item`, `BidList`, `MaxBidder`.

#### Functions

- **`NewAuctionDetails(auction *Auction)`**: Initializes a new `AuctionDetails` instance.
- **`ItemHas(condition, category string)`**: Checks if an item matches specified condition and category criteria.

---

### `notification.go`

Defines the `Notification` model and repository functions for managing notifications.

#### Types

- **`Notification`**: Represents notifications sent to users.
- **Fields**:
  - `ID`, `UserId`, `Content`, `Type`, `Timestamp`, `IsRead`.

#### Functions

- **`CreateNotification()`**: Adds a new notification record.
- **`GetNotificationsByUserId(userId string)`**: Retrieves notifications for a specific user.
- **`MarkAsRead(id string)`**: Marks a notification as read.
- **`DeleteNotification(id string)`**: Deletes a notification record.

---

### `shipping.go`

Manages the `Shipping` model, which handles shipping-related information for auctions.

#### Types

- **`Shipping`**: Represents shipping details.
- **Fields**:
  - `ID`, `AuctionId`, `Address`, `Status`, `EstimatedDelivery`.

#### Functions

- **`CreateShipping()`**: Adds shipping details for an auction.
- **`GetShippingByAuctionId(auctionId string)`**: Retrieves shipping information by auction ID.
- **`UpdateShipping(id string, updates Shipping)`**: Updates specific shipping details.
- **`DeleteShipping(id string)`**: Deletes shipping details from the database.

---

### `transaction.go`

Handles `Transaction` records associated with payments and financial transactions.

#### Types

- **`Transaction`**: Represents a financial transaction.
- **Fields**:
  - `ID`, `AuctionId`, `UserId`, `Amount`, `Timestamp`, `Status`.

#### Functions

- **`CreateTransaction()`**: Records a new transaction.
- **`GetTransactionsByAuctionId(auctionId string)`**: Retrieves transactions related to a specific auction.
- **`UpdateTransaction(id string, updates Transaction)`**: Updates specific transaction details.
- **`DeleteTransaction(id string)`**: Deletes a transaction record.

---

### `types.go`

Defines shared types and constants used across the project, such as enums for categories or conditions.

#### Types

- **Enums**: Defines enumerations like item categories or conditions.
- **Constants**: Shared constants used across different models and services.

---

### `user.go`

Manages the `User` model and repository functions for handling user-related data.

#### Types

- **`User`**: Represents a platform user.
- **Fields**:
  - `ID`, `Username`, `Email`, `PasswordHash`, `CreatedAt`, `Role`.

#### Functions

- **`CreateUser()`**: Adds a new user record.
- **`GetUserById(id string)`**: Retrieves a user by ID.
- **`UpdateUser(id string, updates User)`**: Updates specific user fields.
- **`DeleteUser(id string)`**: Deletes a user record.

---