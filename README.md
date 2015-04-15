# peer
A peer network for any application


Peers connect to the "network" and maintain a common list of active nodes. 
If a new node has no knowledge of any peers to get synchronisation started, 
it uses a hardcoded fallback.


## Todos

1. tcp client and server goroutines

2. a node runs a server and has a list of peers that it intermittently connects 
to and pings