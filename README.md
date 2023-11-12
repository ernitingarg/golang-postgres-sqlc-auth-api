# Overview

## Prerequisites installation

### Tools

- Dbeaver Universal database tool: [Download](https://dbeaver.io/download/)
- Make: [Download](https://gnuwin32.sourceforge.net/packages/make.htm) or `choco install make`

## Run db docker locally

To run database directly inside container, please follow below steps:

- Download docker

  - [Windows](https://docs.docker.com/desktop/install/windows-install/)
  - [Ubuntu](https://docs.docker.com/engine/install/ubuntu/#install-using-the-repository)

- Run db docker container

```bash
make start_db
```

- To stop docker container

```bash
make stop_db
```

- To see containers logs

```bash
make logs_db
```

- To connect to Postgres database manually

```bash
 docker exec -it postgres bash
 psql -U admin -d postgresdb
```

- To see all tables in database

```bash
\dt;
```
