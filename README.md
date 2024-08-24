# Quisqueya Project

Quisqueya is a Taíno word that means "mother of the earth". It's the indigenous name for the island of Hispaniola, 
which is now made up of Haiti and the [Dominican Republic](https://en.wikipedia.org/wiki/Dominican_Republic). I created 
this project as a way to practice and improve my `Go` skills. Currently it is a CRUD command line application that connects to sqlite3. 
I will slowly but steadily add new features to the repo, once I decide what those will be 😅.

## Develop
 - Go >= 1.2x
 - Sqlite3 >= 3.x

## Create DB and import data

There is a CSV that contains some initial data about provinces in the the Dominican Republic. To creare a DB and populate 
it with data, do the following:

```shell
cd cmd/importer
go build importer.go 
./importer -csv province_data.csv -db quisqueya.db 
```

*Notes*: 
 - The database named `quisqueya.db` will show in the top-level directory.
 - Running this command implies that you are starting a new database, this command should not be run if data already exists. 


