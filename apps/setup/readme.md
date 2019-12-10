## How it works

This app is responsible for setting up the database

### Setup

When executed, the setup app will attempt to connect to the database (Postgres) and execute the SQL scripts that are stored in `setup/sql`. In this folder, you'll find `tiggers.sql`, which is responsible for the data check functions and the before insert or update tigger, and the `models.sql` which is responsible for creating the table following the initial file structure.

The whole process will be promped on the console to keep the user aware of whats is goin on.

Both of the scripts has check to update the functions and triggers if they're already created and the models script only tries to create the table once, so if you run accidentally, any data will be loss.