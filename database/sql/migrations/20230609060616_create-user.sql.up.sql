-- create the user
-- CREATE USER 'web'@'localhost';
-- this is to create a user that can connect from any host, choose this method for container to guarantee a connection
CREATE USER 'web'@'%';

-- Grant privileges to the user
-- GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'%';

-- add the password
-- ALTER USER 'web' IDENTIFIED BY 'password1';
ALTER USER 'web' IDENTIFIED BY 'Mysql/pass1';