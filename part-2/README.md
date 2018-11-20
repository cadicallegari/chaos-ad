# Templatero


## Running  tests

### For unit tests:

```sh
make check parameters="<filepath>.py <Class>.<Function>"
```
`parameters filepath init after ./tests/unit`

#### Examples

- All tests
  ```sh
  make check
  ```

- File test
  ```sh
  make check parameters="path/bla_test.py"
  ```

- Class of file test
  ```sh
  make check parameters="path/bla_test.py ParsingTests"
  ```

- Function of class of file test:
  ```sh
  make check parameters="path/bla_test.py ParsingTests.test_name"
  ```

### For integration tests:

```sh
make check-integration parameters="<filepath>.py <Class>.<Function>"
```
`parameters path init after ./tests/integration`

#### Examples

- All tests:
  ```sh
  make check-integration
  ```

- File test:
  ```sh
  make check-integration parameters="path/bla_test.py"
  ```

- Class of file test:
  ```sh
  make check-integration parameters="path/bla_test.py ClassTests"
  ```

- Function of class of file test:
  ```sh
  make check-integration parameters="path/bla_test.py ClassTests.test_name"
  ```

### For coverage tests:

```sh
make coverage
