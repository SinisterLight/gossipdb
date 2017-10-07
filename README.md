# gossipdb
Distributed Embedded Key-Value store

## What it does?
An in memory cache which listens to neighbouring nodes and replicates
messages using gossip protocol.

gossipdb provides the replication and in-memory persistance layer for
the nodes of the cluster. It provides APIs to put and fetch data in
a key value pair. It also provides an API to fetch the active members of the cluster.

## Roadmap

GossipDb is currently an early alpha project. Api interface is likely to
change, and is not recommended for production usages.

Todos
1. Improve test coverage
2. CI pipeline and code linting tools
3. Config Objects for DB
4. Better capability from Key-Value store, like ttl, and deletes
5. Improved Object parser
6. Cluster health and Status
7. Instrumentation on replication
8. Benchmarking
9. Example usage

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
