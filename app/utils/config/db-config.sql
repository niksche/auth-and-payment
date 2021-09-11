create database test1
create user test3 with encrypted password 'test3'
grant all privileges on database test1 to test3;



create table users (user_id serial PRIMARY KEY, username varchar(50) UNIQUE NOT NULL, password varchar(60) not null);

create table accounts (user_id serial PRIMARY KEY, username varchar(50) UNIQUE NOT NULL, balance real DEFAULT 8.0);

GRANT ALL PRIVILEGES ON TABLE users TO test3;  
GRANT ALL PRIVILEGES ON TABLE accounts TO test3;  


GRANT USAGE, SELECT ON SEQUENCE users_user_id_seq TO test3; 

GRANT USAGE, SELECT ON SEQUENCE accounts_user_id_seq TO test3; 




`BEGIN; UPDATE accounts SET balance = balance - round(1.1 , 1) WHERE username = $1; COMMIT;`