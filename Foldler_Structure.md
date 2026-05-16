my-renting-service/
├── cmd/
│   └── api/
│       └── main.go                 # App entrypoint (initializes DB, connects layers, starts server)
│
├── internal/
│   ├── domain/
│   │   ├── owner.go                # Owner entity struct & Repository interface definition
│   │   ├── room.go                 # Room entity struct & Repository interface definition
│   │   └── user.go                 # User (Tenant) entity struct & Repository interface definition
│   │
│   ├── repository/
│   │   ├── postgres/
│   │   │   ├── connection.go       # Establishes connection to PostgreSQL in Docker
│   │   │   ├── owner_repo.go       # Raw SQL execution for the Owner table
│   │   │   ├── room_repo.go        # Raw SQL execution for the Rooms table
│   │   │   └── user_repo.go        # Raw SQL execution for the Users table
│   │   └── redis/
│   │       └── lock_repo.go        # (Optional) For distributed booking locks later
│   │
│   ├── service/
│   │   ├── owner_service.go        # Business logic for onboarding owners
│   │   ├── room_service.go         # Business logic for rooms (allocating spaces, calculations)
│   │   └── user_service.go         # Business logic for managing tenants
│   │
│   └── handler/
│       ├── http/
│       │   ├── middleware/
│       │   │   └── auth.go         # Authentication/Role checking middleware
│       │   ├── owner_handler.go    # HTTP request/response parsing for Owners
│       │   ├── room_handler.go     # HTTP request/response parsing for Rooms
│       │   ├── user_handler.go     # HTTP request/response parsing for Users
│       │   └── router.go           # HTTP routing setup (using Gin, Chi, or Fiber)
│
├── config/
│   └── config.go                   # Reads environment variables (DB credentials, Ports)
│
├── database/
│   └── migrations/                 # Holds your .sql migration files (Step 4.2 blueprint)
│       ├── 000001_init_schema.up.sql
│       └── 000001_init_schema.down.sql
│
├── .env                            # Local configuration variables (passwords, ports)
├── .gitignore
├── Dockerfile                      # For building your Go binary container
├── docker-compose.yml              # Manages your app container, Postgres, and Redis together
├── go.mod
└── go.sum