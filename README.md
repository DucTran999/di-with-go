# DI With GO - Dependency Injection Examples

[![CI](https://github.com/DucTran999/di-with-go/actions/workflows/ci.yml/badge.svg)](https://github.com/DucTran999/di-with-go/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/DucTran999/di-with-go)](https://goreportcard.com/report/github.com/DucTran999/di-with-go)
[![Go](https://img.shields.io/badge/Go-1.23-blue?logo=go)](https://golang.org)
[![codecov](https://codecov.io/gh/DucTran999/di-with-go/branch/master/graph/badge.svg)](https://codecov.io/gh/DucTran999/di-with-go)
[![Known Vulnerabilities](https://snyk.io/test/github/ductran999/di-with-go/badge.svg)](https://snyk.io/test/github/ductran999/di-with-go)
[![License](https://img.shields.io/github/license/DucTran999/di-with-go)](LICENSE)

🧩 **Some example dependency injection in Go**

This project demonstrates how to implement Dependency Injection (DI) in Go using idiomatic, clean architecture principles.

- Manual DI
- DI with Wire
- DI with Fx and dig

---

## 📦 Project Structure

```csharp
di-with-go/
├── cmd/                # Entry points (e.g., main.go, wire.go)
├── internal/
│   ├── entity/         # Entities
│   ├── repository/     # Data access layer
│   ├── usecase/        # Business logic
│   └── handler/        # HTTP handlers (delivery layer)
│
├── go.mod
└── README.md

```

## 🔧 Setup

- Clone repository

```sh
git clone https://github.com/DucTran999/di-with-go.git
cd di-with-go
go mod tidy
```

- Install task for run scripts

```sh
go install github.com/go-task/task/v3/cmd/task@latest
```

## 🚀 Run the App

```sh
task run
```

- The main program

```sh
  ____ ___    ____  ___
 |  _ \_ _|  / ___|/ _ \
 | | | | |  | |  _| | | |
 | |_| | |  | |_| | |_| |
 |____/___|  \____|\___/

=== DI Menu ===
1) Example DI manual
2) Example DI with wire
3) Exit
Select option:
```

- make api call with curl

```sh
curl http://localhost:9420/join/daniel
```

- server will response

```json
{
  "code": 200,
  "message": "daniel(9420): playing di"
}
```

## 🧪 Run Tests

```sh
task coverage
```

## 📎 Related

- [**Google Wire**](https://github.com/google/wire) – Compile-time dependency injection tool for Go.
- [**Uber Fx**](https://github.com/uber-go/fx) – A powerful Go framework for dependency injection and lifecycle management.

## 🤝 Contributing

Contributions are welcome! If you find a bug, have a suggestion, or want to add an example, feel free to open an issue or submit a pull request.

Steps to contribute:

1. Fork the repository
2. Create your feature branch: `git checkout -b feature/my-feature`
3. Commit your changes: `git commit -am 'Add new feature'`
4. Push to the branch: `git push origin feature/my-feature`
5. Create a new Pull Request

Please follow the existing code style and write tests if applicable.

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).
