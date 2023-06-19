# haiku-me

## A monolithic Go application closely tied to a MySQL database.
 
To run the application, there will be no need to set up the database, it will be taken care of when building the MySQL image to be used.


### Run the application
In this branch, Dockerfiles are available to run to setup the environments.
For these to work, clone the repo
1. in to root directory, run `docker build -f Dockerfile.db -t mysql-db:latest --build-arg MYSQL_USER=ernest --build-arg MYSQL_PASSWORD=connect@db1 --build-arg MYSQL_ROOT_PASSWORD=p@ssw0r/d .` to build the MySQL docker image.
2. run the `docker build -f Dockerfile.app -t app:latest .` to build the Go app image.
3. start a container from the MySQL image and Go images with the following commands respectively;
```
docker run -p 3306:3306 \
--name db \
mysql-db:latest
```
and
```
docker run -p 4000:4000 \
--name app \
app:latest
```
4. Note that the MySQL image takes approximately 130 seconds to properly initialize, run the app container only when the MySQL container has finished running. Check the logs using `docker logs -f db` to confirm.