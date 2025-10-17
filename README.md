# Liatrio Apprenticeship Interview Excercise

This project is a Golang and Fiber web application that returns a JSON which will contain my name (Anmol Virdi) and a timestamp. It is deployed via Docker and Github Actions to a cloud platform.

## How To Run Locally

```bash
docker build -t liatrio_docker .
```
```bash
docker run -p 8080:80 liatrio_docker
```

## To View

```bash
curl http://localhost:8080
```
