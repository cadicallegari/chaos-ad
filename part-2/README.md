# URL aggregator


## Running  tests

### For unit tests:

```sh
make check parameters="<filepath>.py <Class>.<Function>"
```

#### Examples

- All unit tests
```sh
make check
```

- Test module
```sh
make check parameters="tests/unit/bla_test.py"
```

- Entire test class in a module
```sh
make check parameters="tests/unit/bla_test.py ParsingTests"
```

- Especific test
```sh
make check parameters="tests/unit/bla_test.py ParsingTests.test_name"
```

### For integration tests:

```sh
make check-integration parameters="<filepath>.py <Class>.<Function>"
```

#### Examples

- All integration tests
  ```sh
  make check-integration
  ```

- Test module
```sh
make check-integration parameters="/tests/integration/bla_test.py"
```

- Entire test class in a module
```sh
make check-integration parameters="/tests/integration/bla_test.py ClassTests"
```

- Especific test
```sh
make check-integration parameters="/tests/integration/bla_test.py ClassTests.test_name"
```

### For coverage tests:

```sh
make coverage
```

## Releasing
You can release creating a docker image or installing using setup.py

### Docker image

```sh
make release version=<version>
```

Creates a docker image, ready to be published in your registry and used in your preferred container orchestration tool.
Also it will create a git tag with the provided **<version>**

### Installing
To install in the system or using virtualenv

```
make install
```



