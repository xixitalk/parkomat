
# parkomat.tiny


### What is it?

parkomat.tiny forked from [github.com/parkomat/parkomat](https://github.com/parkomat/parkomat), ONLY Web + WebDav server.

### Features

**ONLY Web and WebDav server. NO DNS, NO HTTPS/SSL/TLS, NO domains**

### Installation

Parkomat at the moment doesn't provide pre-built binaries, so you need to have Go 1.5+ installed. Latest version of Go is recommended.

To build, issue:

```
go get github.com/xixitalk/parkomat
```

### Setting up

As a configuration format Parkomat uses [TOML](https://github.com/toml-lang/toml)

### Example Configuration:

Note: instead of `127.0.0.1` use your external IP.

```
[web]
ip = "127.0.0.1"
port = 80
path = "./www"

# make sure that path exists
# for example issue mkdir -p /var/log/parkomat
#access_log = "/var/log/parkomat/access.log"

[webdav]
enabled = true
username = "hello"
password = "world"
# your share will be under http://example.domain/dav/
mount = "/webdav/"

```



```
./parkomat -config_file=/path/to/config.toml
```

You can also use following environment variables, that will overwrite passed arguments:

`PARKOMAT_CONFIG_FILE` - path to the configuration file, for example `/path/to/config.toml`

### Web server directory structure

You `./www` path could look like this:

```
.
├── default
│   └── public_html
│       └── index.html
└── config.toml
```

All your html and other files go to `public_html` directory.
