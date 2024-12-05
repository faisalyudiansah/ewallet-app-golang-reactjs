# Sea Wallet API

## How to setup

- Install the dependencies and create the mock files.

  ```
  go mod tidy
  make mock
  ```

- Create a new postgres database in your local environment.

- Modify the **Makefile**.

  ```Makefile
  ...
  migrateforce:
  	migrate -path ./internal/database/migration/ -database "postgresql://<db_user>:<db_pass>@<db_host>:<db_port>/<db_name>?sslmode=disable" -verbose force 1
  
  migratedown:
  	migrate -path ./internal/database/migration/ -database "postgresql://<db_user>:<db_pass>@<db_host>:<db_port>/<db_name>?sslmode=disable" -verbose down
  
  migrateup:
  	migrate -path ./internal/database/migration/ -database "postgresql://<db_user>:<db_pass>@<db_host>:<db_port>/<db_name>?sslmode=disable" -verbose up
  ...
  ```

- Run the migration.

  ```
  make migrateup
  ```

## How to run

```
make run-test
```

## How to use

[API Documentation](https://documenter.getpostman.com/view/12104547/2sA3Bt2pXq).

## References
