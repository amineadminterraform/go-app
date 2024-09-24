# Platform Iac controller


## Table of Contents
- [Installation](#installation)
- [Usage](#usage)



## Installation
1. Clone the repo:
    ```bash
    git clone https://github.com/amineadminterraform/go-app
    ```

2. Run docker compose up inside the go-app.
the app is exposed in localhost:8000 
You can check the apis at localhost:8000/docs/index.html


2. To Deploy on local machine (Assuming you already have go) you need to install go-migrate :

    ```bash
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz | tar xvz
     ```
3. Start Postgresql Service You can modify the app.env and change the env variables accordingly:
     ```bash
    docker run --name postgres -p5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
     ```
4. Create the db
     ```bash
    make createdb
     ```
5. Bootstrap the db
     ```bash
    make migrateup
     ```     
5. Start the api
     ```bash
    make server
     ```  
## Usage
This is api is used as a backend in a poc to create multiple environment layers from multiple iac inside the same environments and to add another level of automation in infrastructure deployement   
