# Personal 1 week plan to MASTER GO

---

##  PHASE 1: GO FOUNDATIONS

### Go Introduction & Tooling
- [ ] Why Go exists (vs Java / Python / C++)
- [ ] History & design philosophy of Go
- [ ] Installing Go & environment setup
- [ ] go command basics
  - [ ] go run
  - [ ] go build
  - [ ] go install
- [ ] GOPATH vs Go Modules

### Go Modules & Packages
- [ ] go mod init
- [ ] go.mod vs go.sum
- [ ] go mod tidy
- [ ] Importing packages
- [ ] Exporting identifiers (capitalization)
- [ ] Splitting code across files (same package)
- [ ] Splitting code across multiple packages
- [ ] Why & when to use multiple packages
- [ ] Using third-party packages

---

##  PHASE 2: LANGUAGE BASICS & DATA TYPES

### Language Basics
- [ ] var vs :=
- [ ] Zero values
- [ ] Constants
- [ ] iota
- [ ] Scope rules
- [ ] Variable shadowing
- [ ] Idiomatic naming conventions

### Data Types
- [ ] Boolean
- [ ] Numeric types
  - [ ] Signed vs unsigned
  - [ ] Why int exists
  - [ ] When to use int8 / int16
- [ ] Floating point numbers
- [ ] Complex numbers
- [ ] Runes (Unicode)
- [ ] Strings
  - [ ] Raw string literals
  - [ ] Interpreted string literals
- [ ] Type conversion (explicit only)

---

##  PHASE 3: COMPOSITE TYPES & CONTROL FLOW

### Composite Types
- [ ] Arrays
- [ ] Slices
  - [ ] Length vs capacity
  - [ ] Growth behavior
  - [ ] Backing array
- [ ] make() vs new()
- [ ] Slice ↔ array conversion
- [ ] Maps
  - [ ] Comma-ok idiom
  - [ ] Reference semantics
- [ ] Struct vs map (decision criteria)

### Control Flow
- [ ] if / if-else
- [ ] switch
  - [ ] Expression switch
  - [ ] Type switch
- [ ] for loop
- [ ] for range
- [ ] break / continue
- [ ] goto (discouraged)

---

##  PHASE 4: FUNCTIONS, POINTERS & MEMORY

### Functions
- [ ] Function basics
- [ ] Multiple return values
- [ ] Named return values
- [ ] Variadic functions
- [ ] Anonymous functions
- [ ] Closures
- [ ] Recursion
- [ ] Functions as values
- [ ] Returning functions
- [ ] Call-by-value

### Pointers & Memory
- [ ] Pointer basics
- [ ] Nil pointers
- [ ] Pointer vs value semantics
- [ ] Pointers with structs
- [ ] Maps & slices as reference types
- [ ] Data mutation using pointers
- [ ] Why pointer arithmetic is not allowed
- [ ] Garbage collection overview
- [ ] defer keyword

---

##  PHASE 5: STRUCTS, METHODS & INTERFACES

### Structs
- [ ] Defining struct types
- [ ] Struct literals
- [ ] Zero values in structs
- [ ] Structs with pointers
- [ ] Struct tags
  - [ ] json
- [ ] Struct embedding
- [ ] Custom types
- [ ] Type aliases

### Methods
- [ ] Methods vs functions
- [ ] Pointer receivers
- [ ] Value receivers
- [ ] Mutation methods
- [ ] Constructor functions
- [ ] Validation inside constructors

### Interfaces
- [ ] Interface basics
- [ ] Implicit implementation
- [ ] Using interfaces
- [ ] Embedded interfaces
- [ ] Empty interface (any)
- [ ] Type assertions
- [ ] Type switches
- [ ] Interface limitations

### Generics
- [ ] Why generics exist
- [ ] Generic functions
- [ ] Generic types
- [ ] Type constraints
- [ ] Type inference
- [ ] When not to use generics

---

##  PHASE 6: ERROR HANDLING

- [ ] error interface
- [ ] errors.New
- [ ] fmt.Errorf
- [ ] Error wrapping
- [ ] Sentinel errors
- [ ] panic
- [ ] recover
- [ ] Debugging runtime errors

---

## ⚙️ PHASE 7: CONCURRENCY

### Goroutines
- [ ] Creating goroutines
- [ ] Goroutine lifecycle
- [ ] Common goroutine mistakes

### Channels
- [ ] Unbuffered channels
- [ ] Buffered channels
- [ ] Channel directions
- [ ] select statement
- [ ] Multiple goroutines with channels
- [ ] Error channels
- [ ] Fan-in / fan-out
- [ ] Pipelines
- [ ] Race detection

### sync & context
- [ ] sync.Mutex
- [ ] sync.WaitGroup
- [ ] context.Context
- [ ] Deadlines
- [ ] Cancellations
- [ ] Request-scoped context

---

##  PHASE 8: WEB DEVELOPMENT (GIN + SQL)

### Standard Library
- [ ] File I/O
- [ ] JSON encoding / decoding
- [ ] time
- [ ] flag
- [ ] regexp
- [ ] Logging (log / slog)

### Web APIs
- [ ] net/http fundamentals
- [ ] REST principles
- [ ] Gin framework setup
- [ ] Routing
- [ ] Middleware
- [ ] Middleware chaining
- [ ] Request validation

### Database
- [ ] Initializing database
- [ ] SQL drivers
- [ ] Creating tables
- [ ] INSERT queries
- [ ] SELECT queries
- [ ] Prepared statements
- [ ] Query vs Exec
- [ ] Refactoring DB logic

### Authentication & Authorization
- [ ] User signup
- [ ] Password hashing
- [ ] JWT generation
- [ ] JWT verification
- [ ] Authentication middleware
- [ ] Route protection
- [ ] Authorization (ownership checks)

---

##  FINAL
- [ ] Refactoring routes
- [ ] Code organization
- [ ] Project cleanup
- [ ] Course completion
