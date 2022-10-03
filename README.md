# `cmd`

Simple command line interface template for creating internal CLI's 💻

## Structure

```text
├── cmd
│   └── {cmd_name}
│       └── main.go
├── {cmd_name}.go
├── {cmd_sub_name}.go
├── ...
├── go.mod
├── go.sum
├── LICENSE
├── Makefile
├── pkg
│   ├── config
│   │   └── config.go
│   ├── dir
│   │   └── dir.go
│   ├── editor
│   │   └── editor.go
│   ├── fs
│   │   └── fs.go
│   └── shell
│       └── shell.go
└── README.md
```

The layout consists on the `pkg` directory which has all the dependencies that I
need for creating custom CLI's and the `cmd` directory which just executes the
`cmd` package which is all the _{cmd_name}.go_, _{cmd_sub_name}.go_ files...

The _config_ pkg works as a _getter_ and _setter_ of json, where you can
`config.Query()` and `config.Set()` on a specific config path.

```go
conf := c.Conf{
  Id:   "{cmd_name}",
  Dir:  "{configs_path}",
  File: "config.json",
}
```

## References

1. <https://github.com/Odas0R/pomo-cmd>

## Credits

1. <https://github.com/rwxrob/conf>
2. <https://github.com/rwxrob>
