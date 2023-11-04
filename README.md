# PikPakHash

It is a package for calculating the hash of pikpak.

## Usage

```go
ph := pikpakhash.Default()
hash, err := ph.HashFromPath(path)
if err != nil {
    panic(err)
}
fmt.Println(hash)
```
