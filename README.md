# Efishery-Test

| Target Release 	| 1.0                                                                                            	|
|----------------	|------------------------------------------------------------------------------------------------	|
| Epic           	| Standalone Efishery-Test                                                                     	|
| Document Owner 	| [@bagustyo92](https://github.com/bagustyo92)                                         |
| Maintainer     	| [@bagustyo92](https://github.com/bagustyo92)|
| Developer      	| [@bagustyo92](https://github.com/bagustyo92)|
| Sprint         	| 2                                                                                              	|


## Goals Design and Architecture

![Diagram](/docs/efishery-test-diagram.png)
![Sequence Diagram](/docs/Sequence-Diagram-efishery.png)

1. Servers bisa dinyalakan di port berbeda
2. Semua endpoint berfungsi dengan semestinya (3 endpoint auth, 3 endpoint fetch)
3. Dokumentasi endpoint dengan format OpenAPI (API.md)
4. Dokumentasi system diagram-nya dalam format C4 Model (Context dan Deployment)
5. Pergunakan satu repo git untuk semua apps (mono repo)
6. Dockerfile untuk masing-masing app
7. Petunjuk penggunaan dan instalasi di README.md yang memudahkan

## Getting started

API Docs http://localhost:{defined-port}/swagger/index.html

*Notes
This can be access when our app properly running in your computer

### How to run in your local
- Clone branch develop from repo here https://github.com/bagustyo92/auth/tree/efishery-test
- Install docker linux here: [Linux Docker Installation](https://runnable.com/docker/install-docker-on-linux) / [Linux Docker Installation from docker.com](https://docs.docker.com/engine/install/ubuntu/)
- Install docker mac here: [Mac Docker Installation](https://docs.docker.com/docker-for-mac/install/)
- Install docker linux here: [Windows Docker Installation](https://docs.docker.com/docker-for-windows/install/)
- Run mysql on your docker you can see reference here: [mysql docker build](https://hub.docker.com/_/mysql)
- Make sure mysql run properly and port exposed
- Export env file with this command, you can check list of env here [NODE_ENV](NODE_ENV)
- I recommend used development env, by default if you not export it will used development
```
export NODE_ENV=development
```
- Finally we can run our app with this command
```
go run .
```
- Or you can used dockerfile builder to build our app
```
docker build -t {app-name} .
and
docker run -d {app-name}
```
- If found any error you can troubleshoot here [TROUBLESHOOT](https://www.google.co.id/) or you can ask to maintainer or developer who in charge in this project
- Or simply you can google it, yes google is your bestfriends :)