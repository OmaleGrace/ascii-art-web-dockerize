# ASCII-Art-Web—Dockerize

A containerized version of the **ascii-art-web** project. It takes the existing Go web server (which renders text into ASCII art) and packages it into a Docker image so it can be built, shipped, and run consistently on any machine — no local Go installation required.

## 📖 About

This project builds on the original `ascii-art-web` server and focuses on **containerization** rather than new application features. The goals are to:

- Write a Dockerfile that follows Docker best practices (small image, no unnecessary layers, no wasted build cache).
- Build a Docker **image** from that Dockerfile.
- Run the image as a Docker **container**.
- Apply **metadata** (labels) to Docker objects for traceability.
- Clean up unused Docker objects (dangling images, stopped containers, unused build cacfhe) — i.e. garbage collection.

## 🧰 Tech Stack

- **Go** (standard library only — no third-party packages)
- **Docker**

## 🚀 Getting Started

### Prerequisites

- [Docker](https://docs.docker.com/get-started/) installed and running
- Git

### Clone the repository

```bash
git clone https://github.com/omalegrace2009-g/ascii-art-web-dockerize.git
cd ascii-art-web-dockerize
```

### Build the image

```bash
docker build -t ascii-art-web-dockerize:latest .
```

### Run the container

```bash
docker run -d -p 8080:8080 --name ascii-art-web ascii-art-web-dockerize:latest
```

Then open your browser at [http://localhost:8080](http://localhost:8080).

### Stop and remove the container

```bash
docker stop ascii-art-web-dockerize
docker rm ascii-art-web-dockerize
```

## 🏷️ Metadata

Docker labels are applied to the image for identification and maintainability, e.g. maintainer, version, and description. See the `LABEL` instructions in the [Dockerfile](./Dockerfile).

## 🧹 Garbage Collection

To keep the local Docker environment clean of unused objects:

```bash
# Remove stopped containers, dangling images, unused networks, and build cache
docker system prune -f

# Remove dangling images only
docker image prune -f
```

## 📂 Project Structure

```
ascii-art-web-dockerize/
├── Dockerfile
├── go.mod
├── main.go
├── handlers/
├── templates/
├── static/
└── README.md
```

*(Adjust this tree to match your actual folder layout.)*

## 🎯 Learning Objectives

- Understanding client-server basics: HTTP, HTML, Go web servers
- Writing data to a response and receiving data from a request
- Learning Docker fundamentals: images, containers, layers
- Following Dockerfile best practices
- Managing Docker object lifecycle and cleanup

## 👥 Author

- Omale Grace 

## 📜 License

This project is part of the LEEF Centre / 01-edu curriculum and is for educational purposes.