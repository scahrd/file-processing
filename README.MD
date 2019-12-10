# File Processing tool

## Running

### Installing
```
    go get go get ./apps/setup
    go get ./apps/importer
    go run setup/main.go
```

It will download all the dependecies and setup the detabase's structures (tables and triggers).

### Running the app

From within the apps directory, run
```
    go run importer/*.go
```

And now the app will read the files there are into the `importer/files` and process them.
