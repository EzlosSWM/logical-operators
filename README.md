# Logical Operators 

Logical operators is a simple application made with Go, HTMX & TailwindCSS to demonstrate how they work. 

**Operators**
- AND (&&)
- OR (||)
- NOT (!)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## Installation

### Requirements
1. Install TailwindCSS
2. make (optional)
3. air (optional)

```Shell 
go mod download

# if you don't have make installed 
go run cmd/api/main.go

```

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

live reload (air & tailwind)
```bash
make watch
```

tailwind
```bash
make css
```

live reload
```bash
make css
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```

## To Do 
- Deploy (*if someone can help with deploying to Vercel please message me*)
- GIF showcasing work
