# Luanch Blockchain
## Generate genesis file
```
./oce init
```
And you should see something like this:
```
{
  "chain_id": "oce",
  "node_id": "e14c5056212b5736e201dd1d64c89246f3288129",
  "app_message": {
    "secret": "pluck life bracket worry guilt wink upgrade olive tilt output reform census member trouble around abandon"
  }
}
```
## Start up the blockchain
```
./oce start
```
## Reset the blockchain data
```
./oce unsafe-reset-all
```