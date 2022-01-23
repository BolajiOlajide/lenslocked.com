DROP TABLE IF EXISTS users;

CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	age INT,
	first_name VARCHAR(255),
	last_name VARCHAR(255),
	email TEXT UNIQUE NOT NULL
);

INSERT INTO users (
	age,
	first_name,
	last_name,
	email
) VALUES (
	9,
	'Bolaji',
	'Proton',
	'bolaji@olajide.com'
)
