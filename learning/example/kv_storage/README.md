Test KV storage
---

- Load GBs input and fast checking

- Libs
  + badgerDB <<
  + lmdb
  + boltDB


Build
---

```
$ cd learning/example/kv_storage
$ go build .
$ chmod +x kv_storage
```

Usage
---

```
$ cd learning/example/kv_storage
$ ./kv_storage -h

Available Commands:
  db_viewer   DB viewer
  gen_input   Generate IP to file
  help        Help about any command
  run         Run test
  save_db     Save input data to DB
```

-  Generate IP to file
```
# this will gen 100 ips to file "input_raw/list_ips_100.txt"
$ ./kv_storage gen_input --max_len=100

$ ./kv_storage gen_input -h

Flags:
  -f, --from_ip string   generate IP (default "192.10.10.0")
  -h, --help             help for gen_input
  -l, --max_len int      Number of IPs to generate (default 10)
```

- Save input data to DB
```
$ ./kv_storage save_db -i=input_raw/list_ips_100.txt -c=false -m=false

$ ./kv_storage save_db -h

Flags:
  -c, --check             Check existed each value before add
  -d, --db_path string    DB path (default "data_storage/badger")
  -h, --help              help for save_db
  -i, --inp_path string   Input file path (default "input_raw/list_ips_5.txt")
  -m, --mem               Show memory usage
```

- DB viewer
```
# check ip in db
$ ./kv_storage db_viewer

$ ./kv_storage db_viewer -h

Flags:
  -d, --db_path string   DB path (default "data_storage/badger")
  -h, --help             help for db_viewer
```


Benchmark
---

- Using BadgerDB
  + Import 5m ips to DB took 130s (check exists before adding)
  + CPU ~50-90 %
  + Memory ~190-350 M
  + Try `./kv_storage save_db -i=input_raw/list_ips_1m.txt -c=true -m=true`

- Using Map in go
  + Memory:
    - 70->270 M (input file 1m ips with size 14M)
    - 70->800 M (input file 5m ips with size 70M)
  + Try `./kv_storage test_map -i=input_raw/list_ips_1m.txt`

Refs
---

- https://github.com/gostor/awesome-go-storage
