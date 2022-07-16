# TODO App

A multi-purpose todo application 📋

## Getting Started

### Docker

This project relies on a postgres instance to store user and todo data. So that it is easy to get started, the postgres instance
has been dockerized and a set of scripts created for easily working with it.

That said, you will need to have docker installed on your machine. It can be found [here](https://docs.docker.com/get-docker/) 🦾

### Steps

For each of these sections, the assumption is that you are starting from the base of the repo!

**Start Postgres**

```sh
# Run docker-compose to start postgres instance
yarn dev:db:up
```

**Start the API**

For local development, we are using [air](https://github.com/cosmtrek/air) for live reloading (ya know... better dev sanity 😄). That said,
make sure that you go through the docs and ensure you have everything needed for getting it running 🙂

```sh

# Move into api directory
cd api

# Start the server in development with hot reloading 🔥
air
```

**Start the client**

```sh
# Move into the client directory
cd client

# Install dependencies
yarn

# Start dev server
yarn dev

```