# Assignment:

## Tech Stack:
1. GO
2. Docker

## Install GO:
https://golang.org/doc/install

## Install Docker in Ubuntu:
https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-20-04

## To run Docker in Detached mode (Background):
```bash
docker run --name greedy_games_assignment -d -p 9000:9000 greedy_games_assignment
```

## To check logs:
```bash
docker logs -f greedy_games_assignment
```

## To stop server:
```bash
docker stop greedy_games_assignment
```

## Basic concept:
1. Used nested struct to mimic a tree structure
2. Leaves of tree are represented as a map
