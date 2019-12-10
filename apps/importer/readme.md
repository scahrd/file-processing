## How it works

First, the importer will check the `importer/files` directory for any not-yet processed files.
When there's one, it will get the file, split it into smaller chunks and process each one individually.

### Processing the file

Each line will be split by words and some string-processing will be checked (put into strings and values formatting), after that, the line will be inserted into the database.

Before insert the data, a trigger will run directly in the database to check and format the data, removing special chars and validating CPFs and CNPJ numbers. 

If something goes wrong, the line will be skipped and logged in the `importer/files/filed/YYYY_MM_DD.log`.

After processing all the chunks, the temporary files (chunks) will be removed from the `importer/temp_file` and the original file will be moved to `importer/files/processed`.

Every part of the process will be printed in the console so the user can be aware of what is going on along with the whole process.