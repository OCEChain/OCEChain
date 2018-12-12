# Luanch Client
## Generating keys
```
./ocecli keys add <key name>
./ocecli keys add <key name> --recover
./ocecli keys list
```

## Send transactions
```
./ocecli send --from=<from address> --amount=6mycoin --to=<to address> --chain-id=oce
```

## Query transaction
```
./ocecli tx <TX HASH> --chain-id=oce
```

## Query account
```
./ocecli account <address> --chain-id=oce
```