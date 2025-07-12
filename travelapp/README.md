# Travel App

This is a simple example travel application written in Go. It exposes a small HTTP server
that lists available destinations.

## Running

```
go run ./travelapp
```

The server listens on port `8080` and provides the following endpoint:

- `/destinations` - returns a JSON array of destinations

