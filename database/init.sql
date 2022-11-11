CREATE DATABASE canyes;

\connect canyes;

CREATE TABLE users (
	user_id serial PRIMARY KEY,
	mail VARCHAR ( 50 ) UNIQUE NOT NULL
);

CREATE TABLE sessions (
    session_id serial PRIMARY KEY,
    user_id INT NOT NULL,
    expiration TIMESTAMP NOT NULL,
    token VARCHAR ( 128 ) UNIQUE NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE TABLE elements (
    element_id serial PRIMARY KEY,
    emoji VARCHAR ( 50 ) UNIQUE NOT NULL
);

CREATE TABLE interactions (
    interaction_id serial PRIMARY KEY,
    user_id INT NOT NULL,
    element_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id),
    FOREIGN KEY (element_id) REFERENCES elements (element_id)
);

INSERT INTO users (user_id, mail)
VALUES (DEFAULT, 'amorell@cifpfbmoll.eu');