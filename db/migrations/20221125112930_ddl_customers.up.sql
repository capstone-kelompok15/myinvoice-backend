CREATE TABLE customers (
	id integer PRIMARY KEY AUTO_INCREMENT,
	email VARCHAR(255) NOT NULL,
	customer_password text NOT NULL,
	created_at datetime default NOW(),
	updated_at datetime default NOW() on update NOW(),
	UNIQUE KEY unique_email(email)
);

CREATE TABLE customer_details (
	customer_id integer PRIMARY KEY,
	full_name text NOT NULL,
	display_profile_url text,
	FOREIGN KEY (customer_id) REFERENCES customers(id)
);

CREATE TABLE customer_settings (
	customer_id integer PRIMARY KEY,
	is_verified boolean,
	is_deactivated boolean,
	FOREIGN KEY (customer_id) REFERENCES customers(id)
);

CREATE TABLE customer_tokens(
	id integer PRIMARY KEY AUTO_INCREMENT,
	customer_id integer,
	token text,
	device_id text,
	is_login boolean,
	created_at datetime default NOW(),
	FOREIGN KEY (customer_id) REFERENCES customers(id)
);