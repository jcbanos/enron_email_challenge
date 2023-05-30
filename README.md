# Enron Email Challenge
Project where we index Enron email database into ZincSearch. Then create a Vue.js interface powered by a Golang back end to query the database.

## How to use

There is a makefile included with the project that allow you to run the whole project.

To install zinc and have it up and running, run:

```
make install_zinc
make run_zinc
```

To run the indexer, execute the following:

```
make build_indexer
make run_indexer
```

(Note that zinc must be running before you are able to run the indexer)

To run the web-server:

```
make build_server
make run_server
```
(Note that zinc must be running before you are able to run the server)

