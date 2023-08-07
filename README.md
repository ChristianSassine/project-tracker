# Project Manager
A web application to collaborate, manage projects and track tasks

## Description

A web application made for people to be able to manage their projects collaboratively while tracking their tasks. This application uses a PostgreSQL database to store data. It contains 3 main pages while managing projects: 

* An **Overview** page for a summary of the recent changes in the project.
* A **task board** page where you can create and modify tasks, change their states between **TODO**, **IN PROGRESS** and **DONE** by modifying them or with the drag and drop feature. It also supports comments on tasks by different users.
* A **History** page where logs of all events in the project are present

## Getting Started

### Dependencies

* [Docker](https://docs.docker.com/get-docker/)
* Web browser (Chrome, Firefox, Safari, etc.)

### Installing
* Clone the repo with the following command:
```bash
git clone https://github.com/ChristianSassine/project-tracker.git
```

### Usage
* Tweak the environmental variables in the [docker-compose.yml](https://github.com/ChristianSassine/project-tracker/blob/master/docker-compose.yml) (optional but **recommended**)
* Run the containers in the repository with the following command:
```shell
docker compose up
```
This step might take a couple of minutes as docker needs to create the client, server and database containers (normally a couple of minutes, the client specifically takes time to compile **Angular**).

* Access the web application on the port specified for the client in [docker-compose.yml](https://github.com/ChristianSassine/project-tracker/blob/master/docker-compose.yml) (by default 4200, replace if **changed**): `http://localhost:4200/`
