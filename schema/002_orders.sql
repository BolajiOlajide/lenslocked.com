DROP TABLE IF EXISTS orders;

CREATE TABLE orders (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL,
	amount INT,
	description TEXT
);
