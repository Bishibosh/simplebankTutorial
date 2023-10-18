# Simple Bank Tutorial

## Overview
This is a tutorial I've been following on Udemy to practice learning Golang and brushing up on some backend topics.

## Useful commands

To create a new db migration:
```bash
migrate create -ext sql -dir db/migration -seq {MIGRATION_NAME}
```

To invalidate test cache:
```bash
go clean -testcache
```

To generate sql from dbml:
```bash
make
```
or
```bash
dbml2sql --postgres -o docs/schema.sql docs/db.dbml
```