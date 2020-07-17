# MyMint
CLI for tracking personal finances

## Overview
MyMint persists financial transactions in a sqlite database with optional user-specified categorization.  This project was built for personal use and one day it may even get unit tests!

## Getting Started
To use MyMint, set the `MYMINT_CONF_DIR` environment variable to specify the directory where MyMint should store its required files.
```
$ export MYMINT_CONF_DIR=~/.mymint
```

Then run the `init` command to initialize the configuration files and sqlite database.
```
$ mymint init
```

## Functionality
```
$ mymint
mymint persists personal finance data

Usage:
  mymint [command]

Available Commands:
  completion  Create shell completions
  delete      Delete rows from a table
  help        Help about any command
  init        Initialize MyMint
  ledger      Ledger operations
  source      Datasource file operations
  table       Table operations
```

## Ledger
The main functionality is via the ledger command.  The ledger is a table in a sqlite database that tracks transactions.  When adding a transaction to the ledger, an optional category can be specified to enable groupby aggregations and other analyses.
```
$ mymint ledger
Ledger operations

Usage:
  mymint ledger [flags]
  mymint ledger [command]

Available Commands:
  add         Add an entry to the ledger
  category    Interact with ledger categories
  dump        Print the ledger to the console
  ```

## Datasources
In addition to the ledger, MyMint can track any datasource (e.g., credit card transations) by adding a configuration file for the datasource specification and adding data from a CSV.

First, add a JSON file to the `$MYMINT_CONF_DIR/datasources/` directory.  The JSON file specification follows:
```
{
  "name": "datasource_name",
  "table": "name_of_table",
  "schema": {
    "id": "INTEGER PRIMARY KEY",
    "field1": "SQLITE3 TYPE",
    "field2": "SQLITE3 TYPE",
    ...
  },
  "csv": {
    "header": <boolean indicating presence of CSV header>,
    "fields": [
      "CSV_field1",
      "CSV_field2",
      ...
    ]
  }
}
```

## Categories
Description coming here
