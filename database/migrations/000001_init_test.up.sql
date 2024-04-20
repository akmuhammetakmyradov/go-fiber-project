-- Tables
CREATE TABLE IF NOT EXISTS users(
	"id" SERIAL PRIMARY KEY,
	"name" CHARACTER VARYING(300),
  "login" CHARACTER VARYING(100) UNIQUE NOT NULL,
  "password" CHARACTER VARYING(300) NOT NULL,
	"type" CHARACTER VARYING(5) DEFAULT 'user',       --2 types of user can be record: user, admin
	"created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS posts(
	"id" SERIAL PRIMARY KEY,
	"header" CHARACTER VARYING(300),
  "text" TEXT,
	"created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- superadmin created because of every admin can not create user without authorization. password is 1234
INSERT INTO users (name, login, password, type) 
VALUES ('first user', 'superadmin', '$2a$12$Eg2RiB4/zQIVObUWz2ZgeOzqVOm7pd3nSHswsPThclqn4OBaNAWse', 'admin');