# File Processing tool

## Running with Docker
Download docker-compose if not already installed Then run the following commands:
```
	$ git clone https://github.com/scahrd/file-processing.git
	$ docker-compose up
```

## Running without Docker

### Installing
```
    cd apps/setup
    go get
```

Return to the apps directory and run
    ```go run setup/main.go```

It will now setup the detabase's structures (tables and triggers).
```
    cd ../importer
    go get
```

### Running the app

From within the apps directory, run
```
    go run importer/*.go
```
And now the app will read the files there are into the `importer/files` and process them.