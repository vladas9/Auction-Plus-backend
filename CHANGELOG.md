# Auction+ Changelog

## v0.5.0 - Final Release and Documentation (Date: 2024-11-08)

### Added
- **Detailed Documentation**:
  - Comprehensive README files with installation, setup, and usage instructions for both backend and frontend repositories.
  - Added API documentation, detailing each endpoint and required parameters for easier integration and testing.
- **User Testing and Feedback Loop**:
  - Introduced user feedback forms and implemented minor adjustments based on beta user suggestions.

### Improved
- **Frontend Testing**:
  - Extended frontend testing coverage for critical features, including auction display, bid functionality, and navigation, ensuring a more stable release.
- **Final UI Adjustments**:
  - Updated UI elements based on final feedback, focusing on accessibility and visual consistency.

---

## v0.4.0 - Code Refactoring and Performance Optimization (Date: 2024-10-25)

### Added
- **Repository Layer in Backend**:
  - Created a repository layer to better manage database interactions and reduce code redundancy, primarily developed by Andrei and Alexander.

### Improved
- **Code Refactoring**:
  - Conducted an extensive refactoring process to improve code readability, modularity, and maintainability across both backend and frontend codebases.
  - Enhanced performance by optimizing database queries and reducing latency in API responses.
- **Performance Tuning**:
  - Vladislav Amza optimized WebSocket connections and reduced response times, enabling the backend to support concurrent user connections more efficiently.

### Fixed
- **Bug Fixes**:
  - Resolved minor bugs in auction display, real-time bid updates, and frontend navigation.

---

## v0.3.0 - Real-Time Communication & UX Updates (Date: 2024-10-11)

### Added
- **WebSocket Integration**:
  - Researched and integrated WebSocket functionality by Andrei and Vladislav Amza to support real-time bid updates.
  - Enabled live auction tracking, where users can see bids in real time as they occur.
- **Trigger Method for Bidding**:
  - Implemented a backend trigger method, developed by Alexander, to automatically process and validate incoming bids in real-time.

### Improved
- **UI/UX Enhancements**:
  - Refined the user interface based on UX research, conducted by Vladislav Titerez, to improve navigation and visual hierarchy.
  - Updated layout for auction listings and bid pages to make the platform more intuitive and engaging.
- **Testing**:
  - Introduced initial frontend testing for core components and user interactions, guided by Dumitru's research on testing frameworks.

### Fixed
- **Data Synchronization Issues**:
  - Resolved discrepancies between frontend and backend data displays, ensuring real-time updates reflect accurately across all user interfaces.

---

## v0.2.0 - Database and API Enhancements (Date: 2024-09-27)

### Added
- **Authentication and Authorization**:
  - Developed and integrated user authentication using JWT for secure login.
  - Added role-based authorization to restrict access to certain features for regular users versus admins.
- **Auction Management API**:
  - Implemented CRUD operations for auctions, allowing admins to create, update, and delete auctions.
  - Established API endpoints for bid submission, with basic validation and bid tracking.
- **Statistics Module**:
  - Added a new statistics module in the backend, developed by Alexander, to compute real-time data for auction analytics.

### Improved
- **Database Populating**:
  - Enhanced dummy data generator, allowing for diverse item categories and auction types to support realistic testing.
  - Automated database population as part of the development setup process.

---

## Initial Release - v0.1.0 (Date: 2024-09-13)

### Added
- **Backend Foundation**:
  - Set up core project structure using Go with a modular architecture.
  - Implemented essential server configurations and environment management.
  - Established connection to the database with foundational API endpoints.
- **Frontend Foundation**:
  - Initialized the frontend project using React and Vite.
  - Created basic layout and interface structure for Auction+ platform.
  - Developed a static homepage and basic navigation for initial user testing.
- **Database Schema**:
  - Defined database schema to handle users, auctions, bids, and items.
  - Populated the database with initial dummy data to support testing and development.
- **Project Management**:
  - Created GitHub project board to manage tasks and track project progress.

---