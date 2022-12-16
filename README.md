# go-fast-temp

#### What is this repository for?

* rest api template
* golang
* gin
* gorm
* postgres
* logging
* authentication
* Version 1.0

#### What to Focus on Folder Structure
* /cmd
is for main.go file
* /pkg
almost everything else is here with expressive names names /middleware, /data, /routes, /services, /dto, /controllers

routing: you can check /routes/initial.go which is kind of uncommon approch for routing. but it is very easy to understand and maintain. every controller defined its own routes after initial route.

logging: /middlewares/logging.go is a middleware for logging. it is very easy to use. just add it to your route. it will log every request and response. you may want to remove file logging.


#### How do I get set up? ###
for the very beginning check the `.env` file for configurations. for migration after changes set `ANY_MIGRATION`to true.

check `makefile` for debug and test purposes. it is easier to run with makefile from root directory.

```sh
make run
make test 
```


#### Dockerize

* docker build
```docker build --tag project:nametag .```
* docker run with mounted volume

obviously you can run without volume if you dont care about file logs

```docker run --name project-app -v /path/to/folder/outofdocker:/app/logs -d -p 8080:8080 project:nametag```

