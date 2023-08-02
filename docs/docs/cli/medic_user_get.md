## medic user get

Get details for an user within a mediator control plane

### Synopsis

The medic user get subcommand lets you retrieve details for an user within a
mediator control plane.

```
medic user get [flags]
```

### Options

```
  -e, --email string      Email for the user to query
  -h, --help              help for get
  -i, --id int32          ID for the user to query (default -1)
  -o, --output string     Output format (json or yaml)
  -u, --username string   Username for the user to query
```

### Options inherited from parent commands

```
      --config string      config file (default is $PWD/config.yaml)
      --grpc-host string   Server host (default "localhost")
      --grpc-port int      Server port (default 8090)
```

### SEE ALSO

* [medic user](medic_user.md)	 - Manage users within a mediator control plane
