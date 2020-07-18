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
  "name": <datasource name>,
  "table": <table name>,
  "schema": {
    "id": "INTEGER PRIMARY KEY",
    <field1>: <SQLITE3 TYPE>,
    <field3>: <SQLITE3 TYPE>,
    ...
  },
  "categoryMatchField": <field to match substrings for category assignment>,
  "categoryField": <field to use for assigned categories>,
  "csv": {
    "header": <boolean indicating presence of CSV header>,
    "fields": [
      <CSV_field1>,
      <CSV_field2>,
      ...
    ]
  }
}
```
Then create a table using this configuration:
```
mymint table create <datasource_filename>
```
Transactions can then be added from a CSV file:
```
mymint table csv -s <datasource_filename> path/to/file.csv
```

## Categories
### Ledger Categories
```
$ mymint ledger category
Interact with ledger categories

Usage:
  mymint ledger category [flags]
  mymint ledger category [command]

Available Commands:
  add         Add a category to the ledger categories
  delete      Delete a category from the ledger categories
  ls          List the ledger categories
```
Ledger categories must be added before a category can be applied to a transaction.  As all ledger transactions are manually added, an associated category is manually specified when using the `ledger add` subcommand.

### Datasource Categories
Categories are automatically added to non-ledger datasources when using the `table csv` subcommand.  Manage the mapping of keyword substrings from the `categoryMatchField` field of the CSV to desired categories using the `source category` subcommand.
```
$ mymint source category
Interact with the keyword category mapping for datasources

Usage:
  mymint source category [flags]
  mymint source category [command]

Available Commands:
  add         Add a category to the keyword category mappings
  delete      Delete a category from the keyword category mappings
  ls          List the keyword category mapping
```
For example, to automatically assign a `Netflix` transaction to the `subscriptions` category, add `Netflix` as a key and `subscriptions` as a value:
```
./mymint source category add -k Netflix -v subscriptions
```
The substring matching is case insensitive.
