# chainpusher

Chainpusher will immediately push the transaction to you when a transaction you care about occurs.

# Features

- Cryto currency
  - Bitcoin
  - Ethereum
  - Tron
  - Usdt
  - Bitcon Cash

- Push supports
  - email
  - telegram
  - http
  - grpc
  - function call (golang)
  - kafka
  - rabbitmq

This project is currently under development.

# Getting started

```bash
$ ./chainpusher run
```

## Watch wallet

```yml
wallets:
  - TDppqiBfL8u6VYWYDKa56UbdZ1QQFyKKgs
  - TFMK9fcEs9wc7dg86YqxTKP5eLxJhgGHKK
```

## Push

```yml
receive:
  telegram:
    token: TOKEN
```