# pwd_hash

The pwd_hash package provides two functions for password hashing and verification in Go applications: Hash and Verify. These functions use a cryptographic hashing algorithm to securely transform a plaintext password into a hash value that can be stored and compared against later.

## Installation

```bash
go get -u github.com/ginger-go/pwd_hash
```

## Documentation

Please refer to [https://ginger-go.gitbook.io/pwd_hash/](https://ginger-go.gitbook.io/pwd_hash/) for the documentation.

## Perform tests

Run ginkgo tests with the following command:
```bash
ginkgo -r -v -p --cover --coverpkg=github.com/ginger-go/pwd_hash
```

Read the coverage report with the following command:
```bash
go tool cover -html=coverprofile.out
```

