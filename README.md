# `go-gen-lo`

Generate lo methods for slice struct.

## Usage

```go
// TODO:
```

## Help

```shell
Generate lo methods for slice struct.

Usage:
  gen-slice-accessors [flags]

Flags:
  -e, --entity string   target entity name. e.g. User or *User
  -h, --help            help for gen-slice-accessors
  -i, --input string    input file name
  -o, --output string   output file name
  -s, --slice string    target slice name. e.g. Users
```

## [Example](./example)

- Common case ([src](./example/user.go) / [generated code 1](./example/users_list_gen.go) / [generated code 2](./example/users_ptr_gen.go))

## Spec

- ✅ Support
- 🚫 Cannot Support(TODO: add issue)
- ☑️ Not supported yet

| samber/lo | Basic | Extend |
| --------- | ----- | ------ |
| Filter    | ✅    | ✅     |
| Map       | ✅    | 🚫     |
| KeyBy     | 🚫    | ✅     |
| GroupBy   | ☑️    | ☑️     |
| Find      | ☑️    | ☑️     |

- TODO: https://github.com/samber/lo?tab=readme-ov-file#-spec

## E2E

```shell
$ go generate ./example
$ go run ./example
```

## LICENSE

[MIT](./LICENSE)
