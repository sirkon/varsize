# varsize
Functions to compute varint and uvarint encoding size without an actual encoding.

## Install.

```shell
go get github.com/sirkon/varsize
```

## Usage.

```go
fmt.Println(varsize.Int(127), varsize.Int(128))
// Output: 1 2
```