# Cerberus
![img](https://img.shields.io/docker/cloud/build/furya/cerberus.svg) ![img](https://img.shields.io/github/go-mod/go-version/TessorNetwork/cerberus) ![img](https://img.shields.io/github/v/tag/TessorNetwork/cerberus?label=version)

Cerberus is a Furya oracle. Cerberus stores and validates PDV (private data value).

## cerberusd

cerberusd provides http API to pdv storing functionality. It receives and validates user PDV and sends it to SQS queue.

### Parameters

| CLI param         | Environment var          | Default | Description
|---------------|------------------|---------------|---------------------------------
| http.host         | HTTP_HOST         | 0.0.0.0  | host to bind server
| http.port    | HTTP_PORT    | 8080  | port to listen
| http.max-body-size    | HTTP_MAX_BODY_SIZE    | 8000000  | max requests' body size in bytes
| postgres    | POSTGRES    | host=localhost port=5432 user=postgres password=root sslmode=disable  | postgres dsn
| postgres.max_open_connections    | POSTGRES_MAX_OPEN_CONNECTIONS    | 0 | postgres maximal open connections count, 0 means unlimited
| postgres.max_idle_connections    | POSTGRES_MAX_IDLE_CONNECTIONS    | 5 | postgres maximal idle connections count
| postgres.migrations    | POSTGRES_MIGRATIONS    | /migrations/postgres | postgres migrations directory
| s3.endpoint    | S3_ENDPOINT    | localhost:9000  | s3 storage endpoint
| s3.region    | S3_REGION    |  | s3 storage region
| s3.access-key-id    | S3_ACCESS_KEY_ID    |  | Access KeyID for S3 storage
| s3.secret-access-key    | S3_SECRET_ACCESS_KEY    |   | Secret Key for S3 storage
| s3.use-ssl    | S3_USE_SSL    | false  | do use ssl for S3 storage connection?
| s3.bucket    | S3_BUCKET    | cerberus  | bucket name for S3 storage
| sqs.region    | SQS_REGION    |   | sqs region
| sqs.access-key-id | SQS_ACCESS_KEY_ID | | access key id for SQS
| sqs.secret-access-key | SQS_SECRET_ACCESS_KEY | | secret access key for SQS
| sqs.queue | SQS_QUEUE | testnet | SQS queue name
| save-pdv-throttle-period    | SAVE_PDV_THROTTLE_PERIOD    | 10m  | how often the user can send PDV to save
| reward-map-config | REWARD_MAP_CONFIG | configs/rewards.yml | path to yaml [config](configs/rewards.yml) with pdv rewards
| min-pdv-count | MIN_PDV_COUNT | 100 | minimal count of pdv to save
| max-pdv-count | MAX_PDV_COUNT | 100 | maximal count of pdv to save
| encrypt-key    | ENCRYPT_KEY    |   | private key for data encryption in hex
| sentry.dsn    | SENTRY_DSN    |  | sentry dsn
| log.level   | LOG_LEVEL   | info  | level of logger (debug,info,warn,error)
| pdv-rewards.pool-size | PDV_REWARDS_POOL_SIZE   | 100000000000  | PDV rewards (uFUR)
| pdv-rewards.interval  | PDV_REWARDS_INTERVAL  | 720h  | how often to pay PDV rewards
| hades.url | HADES_URL | | Hades service url

## processord

processord receives PDVs from SQS, rewards users and stores data into FileStorage

### Parameters

| CLI param         | Environment var          | Default | Description
|---------------|------------------|---------------|---------------------------------
| http.host         | HTTP_HOST         | 0.0.0.0  | host to bind server
| http.port    | HTTP_PORT    | 8080  | port to listen
| postgres    | POSTGRES    | host=localhost port=5432 user=postgres password=root sslmode=disable  | postgres dsn
| postgres.max_open_connections    | POSTGRES_MAX_OPEN_CONNECTIONS    | 0 | postgres maximal open connections count, 0 means unlimited
| postgres.max_idle_connections    | POSTGRES_MAX_IDLE_CONNECTIONS    | 5 | postgres maximal idle connections count
| postgres.migrations    | POSTGRES_MIGRATIONS    | /migrations/postgres | postgres migrations directory
| s3.endpoint    | S3_ENDPOINT    | localhost:9000  | s3 storage endpoint
| s3.region    | S3_REGION    |  | s3 storage region
| s3.access-key-id    | S3_ACCESS_KEY_ID    |  | Access KeyID for S3 storage
| s3.secret-access-key    | S3_SECRET_ACCESS_KEY    |   | Secret Key for S3 storage
| s3.use-ssl    | S3_USE_SSL    | false  | do use ssl for S3 storage connection?
| s3.bucket    | S3_BUCKET    | cerberus  | bucket name for S3 storage
| sqs.region    | SQS_REGION    |   | sqs region
| sqs.access-key-id | SQS_ACCESS_KEY_ID | | access key id for SQS
| sqs.secret-access-key | SQS_SECRET_ACCESS_KEY | | secret access key for SQS
| sqs.queue | SQS_QUEUE | testnet | SQS queue name
| blockchain.node   | BLOCKCHAIN_NODE    | http://zeus.testnet.furya.xyz:26657  | furya node address
| blockchain.from   | BLOCKCHAIN_FROM    |  | furya account name to send stakes
| blockchain.tx_memo   | BLOCKCHAIN_TX_MEMO    | | furya tx's memo
| blockchain.chain_id   | BLOCKCHAIN_CHAIN_ID    | testnet | furya chain id
| blockchain.client_home   | BLOCKCHAIN_CLIENT_HOME    | ~/.furyacli | furyacli home directory
| blockchain.keyring_backend   | BLOCKCHAIN_KEYRING_BACKEND    | test | furyacli keyring backend
| blockchain.keyring_prompt_input   | BLOCKCHAIN_KEYRING_PROMPT_INPUT    | | furyacli keyring prompt input
| blockchain.gas   | BLOCKCHAIN_GAS    | 10  | gas amount
| blockchain.fee   | BLOCKCHAIN_FEE    | 1ufury  | transaction fee
| sentry.dsn    | SENTRY_DSN    |  | sentry dsn
| log.level   | LOG_LEVEL   | info  | level of logger (debug,info,warn,error)

## rewardsd

rewardsd send PDV rewards

### Parameters

| CLI param         | Environment var          | Default | Description
|---------------|------------------|---------------|---------------------------------
| postgres    | POSTGRES    | host=localhost port=5432 user=postgres password=root sslmode=disable  | postgres dsn
| postgres.max_open_connections    | POSTGRES_MAX_OPEN_CONNECTIONS    | 0 | postgres maximal open connections count, 0 means unlimited
| postgres.max_idle_connections    | POSTGRES_MAX_IDLE_CONNECTIONS    | 5 | postgres maximal idle connections count
| postgres.migrations    | POSTGRES_MIGRATIONS    | /migrations/postgres | postgres migrations directory
| blockchain.node   | BLOCKCHAIN_NODE    | http://zeus.testnet.furya.xyz:26657  | furya node address
| blockchain.from   | BLOCKCHAIN_FROM    |  | furya account name to send stakes
| blockchain.tx_memo   | BLOCKCHAIN_TX_MEMO    | | furya tx's memo
| blockchain.chain_id   | BLOCKCHAIN_CHAIN_ID    | testnet | furya chain id
| blockchain.client_home   | BLOCKCHAIN_CLIENT_HOME    | ~/.furyacli | furyacli home directory
| blockchain.keyring_backend   | BLOCKCHAIN_KEYRING_BACKEND    | test | furyacli keyring backend
| blockchain.keyring_prompt_input   | BLOCKCHAIN_KEYRING_PROMPT_INPUT    | | furyacli keyring prompt input
| blockchain.gas   | BLOCKCHAIN_GAS    | 10  | gas amount
| blockchain.fee   | BLOCKCHAIN_FEE    | 1ufury  | transaction fee
| sentry.dsn    | SENTRY_DSN    |  | sentry dsn
| log.level   | LOG_LEVEL   | info  | level of logger (debug,info,warn,error)
| pdv-rewards.pool-size | PDV_REWARDS_POOL_SIZE   | 100000000000  | PDV rewards (uFUR)
| pdv-rewards.interval  | PDV_REWARDS_INTERVAL  | 720h  | how often to pay PDV rewards


## syncd

syncd binary listens to blockchain and reacts on `operations/ResetAccount` message.

### Parameters

| CLI param         | Environment var          | Default | Description
|---------------|------------------|---------------|---------------------------------
| blockchain.node   | BLOCKCHAIN_NODE    | zeus.testnet.furya.xyz:9090 | true | furya grpc node address
| blockchain.timeout   | BLOCKCHAIN_TIMEOUT    | 5s| true | timeout for requests to blockchain node
| blockchain.retry_interval   | BLOCKCHAIN_RETRY_INTERVAL    | 2s | true | interval to be waited on error before retry
| blockchain.last_block_retry_interval   | BLOCKCHAIN_LAST_BLOCK_RETRY_INTERVAL    | 1s | true | duration to be waited when new block isn't produced before retry
| postgres    | POSTGRES    | host=localhost port=5432 user=postgres password=root sslmode=disable  | postgres dsn
| postgres.max_open_connections    | POSTGRES_MAX_OPEN_CONNECTIONS    | 0 | postgres maximal open connections count, 0 means unlimited
| postgres.max_idle_connections    | POSTGRES_MAX_IDLE_CONNECTIONS    | 5 | postgres maximal idle connections count
| postgres.migrations    | POSTGRES_MIGRATIONS    | /migrations/postgres | postgres migrations directory
| s3.endpoint    | S3_ENDPOINT    | localhost:9000  | s3 storage endpoint
| s3.region    | S3_REGION    |  | s3 storage region
| s3.access-key-id    | S3_ACCESS_KEY_ID    |  | Access KeyID for S3 storage
| s3.secret-access-key    | S3_SECRET_ACCESS_KEY    |   | Secret Key for S3 storage
| s3.use-ssl    | S3_USE_SSL    | false  | do use ssl for S3 storage connection?
| s3.bucket    | S3_BUCKET    | cerberus  | bucket name for S3 storage
| sentry.dsn    | SENTRY_DSN    |  | sentry dsn
| log.level   | LOG_LEVEL   | info  | level of logger (debug,info,warn,error)

## Development
### Makefile
#### Update vendors
Use `make vendor`
#### Install required for development tools
You can check all tools existence with `make check-all` or force installing them with `make install-all` 
##### golangci-lint 1.29.0
Use `make install-linter`
##### swagger v0.25.0
Use `make install-swagger`
##### gomock v1.4.3
Use `make install-mockgen`
#### Build docker image
Use `make image` to build local docker image named `cerberus-local`
#### Build binary
Use `make build` to build for your OS or use `make linux` to build for linux(used in `make image`) 
#### Run tests
Use `make test` to run tests. Also you can run tests with `integration` tag with `make fulltest`
