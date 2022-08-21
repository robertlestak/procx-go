# procx-go

`procx-go` is a simple wrapper library for [procx](https://github.com/robertlestak/procx). `procx-go` exports a single function, `Procx`, which accepts one argument, `args`, and produces two return values, `io.Reader` and `error`. This exposes a versatile DAO interface to all procx drivers which can be imported into an application for more integrated usage.

You must have `procx` installed in order to use `procx-go`.

See `examples/main.go` for an example of using one codebase and dynamically pulling from multiple services with a single configuration array.

Here is a basic example:

```golang
args := []string{
    "-driver",
    "redis-list",
    "-redis-host",
    "localhost",
    "-redis-port",
    "6379",
    "-redis-key",
    "test-key",
}
data, err := procx.Procx(args)
if err != nil && err != procx.ErrNoData {
    panic(err)
} else if err == procx.ErrNoData {
    println("no data")
    os.Exit(0)
}
bd, err := ioutil.ReadAll(data)
if err != nil {
    panic(err)
}
// you can unmarshal the data into a struct or map, here we just print it
println(string(bd))
```

Since the entire data layer configuration is contained within the `args` string slice, this can be moved to a configuration layer such as Vault or Consul. If you ever need to change the data layer configuration, you can simply update the `args` configuration, and your code will remain entirely the same.