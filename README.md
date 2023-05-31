# haiku-me

A monolithic Go application closely tied to a MySQL database.
 
To run the application;
1. set up the database. The values have been hardcoded (yes, not soo nice!) `/database/sql/create-db`.
2. create the user to interact with the database. `/database/sql/create-user`.
3. set up other a user and sessions table. `/sql/user-table`, `/database/sql/sessions`.
4. run the application.
 
For a container approach, the commands used here are for Podman. 
Same commands can be used for Docker but will vary slightly.
Instruction are in the `/database/container/instruction` file.


### Run the application
Start the application with the `go run ./cmd/web`. It runs on port 4000.

