# bitcoind-healthcheck

Use REST (cookie/auth) to verify if a bitcoind is synced

Exposes an HTTP endpoint (default `:8080`) that returns `200` when the node
passes all enabled checks and `503` otherwise, with a JSON body listing each
check result. Intended for use as a Kubernetes readiness probe sidecar for
`bitcoind`-compatible nodes (including Elements).

## Checks

| Check | Enabled | Passes when |
|---|---|---|
| `verificationprogress` | always | `getblockchaininfo.verificationprogress > 0.9999` |
| `syncheight` | default on | `headers - blocks <= SYNC_HEIGHT_MAX_LAG` from `getblockchaininfo` |
| `gettxindexinfo` | `TXINDEX_ENABLED=true` | txindex reports synced |
| `estimatesmartfee` | `FEE_ESTIMATION_ENABLED=true` | a fee estimate is available |

`verificationprogress` alone is not enough after a restart: a node replaying
blocks can report `>0.9999` while its active chain still lags its known
headers. The `syncheight` check keeps the endpoint unhealthy until the active
chain catches up. Set `SYNC_HEIGHT_CHECK_ENABLED=false` to restore the old
behavior, or `SYNC_HEIGHT_MAX_LAG` to tolerate a small lag.

## Configuration

| Variable | Default | Description |
|---|---|---|
| `RPC_HOST` | `localhost:8332` | Node RPC host:port |
| `RPC_USER` / `RPC_PASS` | — | RPC credentials (basic auth) |
| `RPC_COOKIE_PATH` | — | RPC cookie file, used when `RPC_USER` is unset |
| `PORT` | `8080` | HTTP listen port |
| `CACHE_EXPIRE_SECONDS` | `14` | Per-check RPC result cache TTL. Keep this low; a high value delays readiness transitions by up to the TTL. |
| `SYNC_HEIGHT_CHECK_ENABLED` | `true` | Require active chain height to match known headers |
| `SYNC_HEIGHT_MAX_LAG` | `0` | Allowed number of blocks behind headers |
| `TXINDEX_ENABLED` | `false` | Also require a synced txindex |
| `FEE_ESTIMATION_ENABLED` | `false` | Also require fee estimation to be available |
| `FEE_ESTIMATION_TARGET` | `1` | Confirmation target for the fee estimation check |
