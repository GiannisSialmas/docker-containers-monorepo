# Chinook database

[Dockerfile](https://github.com/GiannisSialmas/docker/blob/master/postgres/chinook/Dockerfile)


The chinook database implemented in a docker postgres docker container.
This is a bad example because it loads the data when the container is started and not during the build so it takes a lot of time
```
Database: postgres
User: postgres
Password: chinook
```

![Chinook database schema](chinook.jpg)