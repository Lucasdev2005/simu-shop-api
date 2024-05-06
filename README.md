# simu-shop-api

## About
  This project is designed as a portfolio, simulating the flow of publications and product sales in a fictitious e-commerce. Here, users can publish products, while other users can search for products by theme or best-selling, in addition to making purchaes.


## Technologies used

<div align="center"> 
  <img width="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original.svg" />
  <img width="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/postgresql/postgresql-plain-wordmark.svg" />
  <img width="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/docker/docker-original.svg" />
</div>

- ```Postgres```
- ```Go```
- ```Docker```

## Compiling locally

1) install ```docker``` and ```docker compose``` on your machine for external dependencies.
2) install a ```1.2X``` version from Go on your machine.
3) run ``` docker compose up -d ``` from up all external dependencies (Database, MQ, Grafana)
4) create a ```.env``` file with the variables from the ```.env.example``` file.
5) run ``` go run cmd/api/api.go ``` to start API.
