# ChainPusher

[![codecov](https://codecov.io/github/chainpusher/chainpusher/graph/badge.svg?token=JZE9C481WY)](https://codecov.io/github/chainpusher/chainpusher)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/chainpusher/chainpusher)
[![Forum](https://img.shields.io/badge/community-forum-00afd1.svg?style=flat-square)](https://github.com/chainpusher/chainpusher/discussions)
[![Twitter](https://img.shields.io/badge/twitter-@yoichi24526-55acee.svg?style=flat-square)](https://twitter.com/yoichi24526)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/chainpusher/chainpusher/master/LICENSE.txt)




ChainPusher will immediately push the transaction to you when a transaction you care about occurs.

ChainPusher continuously monitors the blockchain network, and whenever a new block is written, it sends you the transaction data you are interested in.

ChainPusher supports multiple blockchain technologies and tokens, including Bitcoin, Ethereum's Ether, Tron's TRX, USDT, and Bitcoin Cash.

This feature is applicable in many scenarios, especially when real-time access to blockchain transaction data is needed:

1. **Wallet Activity Monitoring**: Users can use this feature to monitor activity in their cryptocurrency wallets, receiving timely notifications of deposits and transactions.

2. **Transaction Tracking**: Investigators can use this feature to monitor activity associated with specific addresses or transaction types, identifying potential fraud or anomalous behavior.

3. **Payment System Integration**: Merchants can integrate this feature into their payment systems to provide real-time notifications to their customers about the status of payments.

4. **Smart Contract Monitoring**: Developers can use this feature to monitor the execution of smart contracts, as well as transactions and events related to the contracts.

5. **Market Analysis**: Cryptocurrency exchanges and market analysts can use this feature to access real-time transaction data for market analysis and forecasting.

These are just a few examples of how this feature can be used. In reality, it can be employed in any scenario where real-time access to blockchain data is required.

# Features

- Cryto currency
  - Bitcoin
  - **Ether**
  - **Trx**
  - **USDT**
  - Bitcoin Cash

- Push supports
  - Email
  - **Telegram**
  - **HTTP**
  - gRPC
  - Function Call (Go)
  - Kafka
  - RabbitMQ

- Supported blockchains
  - Bitcoin
  - **Ethereum**
  - **TRON**

This project is currently under development.

# Getting started

```bash
$ ./chainpusher

Monitor blockchain data

Usage:
  chainpusher monitor [flags]

Flags:
  -b, --block-file string   File to write raw blockchain data to
  -h, --help                help for monitor
  -t, --trx-file string     File to write transactions to

Global Flags:
      --config string   config file (default is $HOME/.chainpusher.yaml) (default "c")
```

## Configuration

## `infrua_key`

This field is used to specify the Infura project ID or API key for accessing the Infura Ethereum and IPFS nodes. Infura is a service that provides remote Ethereum and IPFS nodes for developers.

- **Type**: String
- **Required**: Yes
- **Example**: 
  ```yml
  infura_key: your_infura_project_id
  ```

To obtain an Infura project ID or API key:
1. Go to the [Infura website](https://infura.io/).
2. Sign up for an account or log in if you already have one.
3. Create a new project.
4. Once the project is created, you will be provided with an Infura project ID or API key. Copy this ID/key and paste it as the value for `infura_key` in your configuration file. 

Ensure that your Infura project has the necessary permissions to access the Ethereum or IPFS networks depending on your application's requirements.

## `wallets`

This field is used to specify the list of cryptocurrency wallet addresses that will be monitored for activity.

- **Type**: Array of strings
- **Required**: No
- **Example**:
  ```yml
  wallets:
    - TDppqiBfL8u6VYWYDKa56UbdZ1QQFyKKgs
    - TFMK9fcEs9wc7dg86YqxTKP5eLxJhgGHKK
  ```

Each item in the `wallets` array should be a valid cryptocurrency wallet address.

## `telegram`

This section is used to configure integration with the Telegram messaging platform for receiving notifications.

- **Type**: Object
- **Required**: Yes
- **Example**:
  ```yml
  telegram:
    tokens:
      - token: "<token_1>"
      - token: "<token_2>"
        user:
          - "<user_id_1>"
          - "<user_id_2>"
  ```

### `tokens`

This field specifies the Telegram bot tokens used for authentication.

- **Type**: Array of objects
- **Required**: Yes
- **Example**:
  ```yml
  tokens:
    - token: "<bot_token_1>"
    - token: "<bot_token_2>"
      user:
        - "<user_id_1>"
        - "<user_id_2>"
  ```

Each item in the `tokens` array represents a Telegram bot token. If you have multiple bot tokens, you can specify them in this array.

#### `user`

This field specifies the user IDs or chat IDs that will receive notifications from the corresponding bot.

- **Type**: Array of int
- **Required**: Conditional (required if user-specific notifications are needed)
- **Example**:
  ```yml
  user:
    - "<user_id_1>"
    - "<user_id_2>"
  ```

Each item in the `user` array should be a valid Telegram user ID or chat ID. These IDs will receive notifications from the bot associated with the corresponding token.

Ensure that your Telegram bot has the necessary permissions to send messages to the specified users or chat groups. Depending on your application, you may need to configure privacy settings or grant additional permissions to the bot.

## `http`

This section is used to configure HTTP endpoints for sending data.

- **Type**: Array of strings
- **Required**: No
- **Example**:
  ```yml
  http:
    - url: "https://example.com/endpoint"
    - url: "https://another-example.com/api"
  ```

### `url`

This field specifies the URL of the HTTP endpoint where data will be sent.

- **Type**: String
- **Required**: Yes
- **Example**:
  ```yml
  url: "https://example.com/endpoint"
  ```

Each item in the `http` array represents a different HTTP endpoint. You can specify multiple endpoints if you need to send data to multiple locations.

Ensure that the specified URLs are valid and accessible. Depending on your application, you may need to configure additional settings such as authentication or headers for the HTTP requests.

## `logger`

This field is used to configure the logging level for the application.

- **Type**: String
- **Required**: No
- **Default**: `INFO`
- **Options**:
The `logger` field allows you to specify the level of detail for logging output. Here are the available options:

  - "TRACE": The most detailed level of logging.
  - "DEBUG": Detailed information useful for debugging purposes.
  - "INFO": Standard informational messages about the application's operation.
  - "WARN": Indicates potential issues or situations that may require attention.
  - "ERROR": Indicates errors that occurred during the application's execution but did not prevent it from continuing.
  - "FATAL": Indicates critical errors that may lead to the termination of the application.
  - "PANIC": Indicates critical errors that result in a panic.

Choose the appropriate logging level based on your debugging and monitoring needs. Ensure that the logging level selected provides sufficient information for troubleshooting while avoiding excessive verbosity.

## Example

```yml
infura_key: <infura_key>

wallets:
  - TDppqiBfL8u6VYWYDKa56UbdZ1QQFyKKgs
  - TFMK9fcEs9wc7dg86YqxTKP5eLxJhgGHKK

telegram:
  tokens:
    - token: <token_1>
    - token: <token_2>
      user:
        - <user_id_1>
        - <user_id_2>

http:
  - url: https://httpbin.org/post

logger: TRACE


```
## Contributing

Thank you for considering contributing! Here's how to get started:

1. **Fork and Clone**: Fork the repo on GitHub, then clone your fork locally.
    ```sh
    git clone https://github.com/your-username/repo-name.git
    ```
2. **Create a Branch**: 
    ```sh
    git checkout -b your-branch-name
    ```
3. **Make Changes**: Implement your feature or fix.
4. **Commit**:
    ```sh
    git commit -m "Describe your changes"
    ```
5. **Push**:
    ```sh
    git push origin your-branch-name
    ```
6. **Pull Request**: Open a pull request on GitHub.

#### Guidelines

- Follow code style.
- Write and run tests.
- Ensure all tests pass.

#### Reporting Issues

- Clearly describe the issue and steps to reproduce.

Thank you for contributing!

## License

This project is licensed under the MIT License. You are free to use, modify, and distribute this software in compliance with the terms of the MIT License.

For the full license text, see the [LICENSE](./LICENSE) file in the repository.