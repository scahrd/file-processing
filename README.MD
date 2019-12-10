# File Processing tool
A quite simple and small command line tool to process a given base file and store the data into a Postgres database.

## Getting started:

Download docker-compose if not already installed Then run the following commands:

```
$ cd ~/go/src
$ git clone https://github.com/scahrd/file-processing.git .
$ docker-compose up
```

This will setup the Postgres container.


### Setting things up
```
$ cd file-processing.git
$ go get go get ./apps/setup
$ go get ./apps/importer
$ go run setup/main.go
```

It will download all the dependecies and setup the detabase's structures (tables and triggers).

### Running the app

From within the apps directory, run:
```
$ go run importer/*.go
```

And now the app will read the files there are into the `importer/files` and process them.


_From within each app directory you can find other readmes to better understand each objective._