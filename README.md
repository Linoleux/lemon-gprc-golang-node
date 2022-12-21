# Description

Contoh client dan server gRPC menggunakan golang dan nodejs.

Baru ada contoh unary rpc doang bang.

# Run the code

Generate method berdasarkan proto yg ada:

```sh
$ make gen
```

Run server (go):

```sh
$ make server
```

Run client (go):

```sh
$ make client
```

Run client (nodejs):

Install library terlebih dahulu

```sh
$ make node-install
```

Kemudian:

```sh
$ make node
```
