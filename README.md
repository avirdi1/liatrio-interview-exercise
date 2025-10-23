# Liatrio Apprenticeship Interview Exercise
[![CI & CD](https://github.com/avirdi1/liatrio-interview-exercise/actions/workflows/ci-cd-workflow.yml/badge.svg)](https://github.com/avirdi1/liatrio-interview-exercise/actions/workflows/ci-cd-workflow.yml)

This project is a **Golang + Fiber** web application that returns a JSON response with my name (**Anmol Virdi**) and a Unix timestamp in milliseconds.  
The app is **containerized using Docker**, automatically **built and tested through GitHub Actions**, and **deployed to an AWS EC2 instance** using Docker Hub as the image registry.

---

## 🌐 Live Instance
The live application can be accessed (when the EC2 instance is running) at:  
👉 [http://54.176.194.82/](http://54.176.194.82/)

---

## What This Project Demonstrates
1. Building a simple **Go Fiber API** that dynamically returns a JSON object.  
2. **Containerizing** the application with Docker.  
3. Using **GitHub Actions** to automate building, testing, versioning, and pushing to Docker Hub.  
4. **Deploying** the versioned Docker image to a cloud platform (AWS EC2).  
5. Showing **Continuous Delivery** by updating the app and seeing the change live.

---

## How To Run Locally

```bash
# 1. Build the Docker image
docker build -t liatrio_app .
```
```bash
# 2. Run the container (maps local port 8080 to container port 80)
docker run -p 8080:80 liatrio_app
```
```bash
# 3. To view locally
curl http://localhost:8080
```
