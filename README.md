<!-- # 🚀 Go Mastery Roadmap (2 Weeks)

A structured, in-depth roadmap to master **Go (Golang)** from foundations to production-ready backend development.

---

## 📌 Table of Contents

- [Phase 1: Go Foundations](#phase-1-go-foundations)
- [Phase 2: Language Basics & Data Types](#phase-2-language-basics--data-types)
- [Phase 3: Composite Types & Control Flow](#phase-3-composite-types--control-flow)
- [Phase 4: Functions, Pointers & Memory](#phase-4-functions-pointers--memory)
- [Phase 5: Structs, Methods & Interfaces](#phase-5-structs-methods--interfaces)
- [Phase 6: Error Handling](#phase-6-error-handling)
- [Phase 7: Concurrency](#phase-7-concurrency)
- [Phase 8: Web Development (Gin + SQL)](#phase-8-web-development-gin--sql)
- [Final](#final)

---

## PHASE 1: GO FOUNDATIONS

### Go Introduction & Tooling

- [ ] Why Go exists (vs Java / Python / C++)
- [ ] History & design philosophy of Go
- [ ] Installing Go & environment setup
- [ ] go command basics
  - [ ] go run
  - [ ] go build
  - [ ] go install
- [ ] GOPATH vs Go Modules

---

## PHASE 2: LANGUAGE BASICS & DATA TYPES

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

## PHASE 3: COMPOSITE TYPES & CONTROL FLOW

### Composite Types

- [x] Arrays
- [x] Slices
  - [x] Length vs capacity
  - [x] Growth behavior
  - [x] Backing array
- [x] make() vs new()
- [x] Slice ↔ array conversion
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

## PHASE 4: FUNCTIONS, POINTERS & MEMORY

### Functions

- [x] Function basics
- [x] Multiple return values
- [x] Named return values
- [x] Variadic functions
- [x] Anonymous functions
- [x] Closures
- [x] Recursion
- [x] Functions as values
- [x] Returning functions
- [x] Call-by-value

### Pointers & Memory

- [x] Pointer basics
- [x] Nil pointers
- [x] Pointer vs value semantics
- [x] Pointers with structs
- [x] Maps & slices as reference types
- [x] Data mutation using pointers
- [x] Why pointer arithmetic is not allowed
- [x] Garbage collection overview
- [x] defer keyword

---

## PHASE 5: STRUCTS, METHODS & INTERFACES

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

- [x] Interface basics
- [x] Implicit implementation
- [x] Using interfaces
- [x] Embedded interfaces
- [x] Empty interface (any)
- [x] Interface basics
- [x] Implicit implementation
- [x] Using interfaces
- [x] Embedded interfaces
- [x] Empty interface (any)
- [x] Type assertions
- [x] Type switches
- [x] Interface limitations

### Generics

- [] Why generics exist
- [] Generic functions
- [] Generic types
- [] Type constraints
- [] Type inference
- [] When not to use generics

---

## PHASE 6: ERROR HANDLING

- [ ] error interface
- [ ] errors.New
- [ ] fmt.Errorf
- [ ] Error wrapping
- [ ] Sentinel errors
- [ ] panic
- [ ] recover
- [ ] Debugging runtime errors

---

## PHASE 7: CONCURRENCY

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

## PHASE 8: WEB DEVELOPMENT (GIN + SQL)

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

## FINAL

- [ ] Refactoring routes
- [ ] Code organization
- [ ] Project cleanup
- [ ] Course completion ✅ -->
