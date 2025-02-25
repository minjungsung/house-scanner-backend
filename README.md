# House Scanner Backend

## Setup

1. Run the PostgreSQL container:
   ```bash
   docker run --name postgres-container -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=house_scanner -p 5432:5432 -d postgres
   ```

2. Run the MongoDB container:
   ```bash
   docker run --name mongo-container -p 27017:27017 -d mongo
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run the server:
   ```bash
   go run cmd/main.go
   ```

