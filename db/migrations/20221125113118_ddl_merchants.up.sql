CREATE TABLE merchants (
	id integer PRIMARY KEY AUTO_INCREMENT,
	merchant_name varchar(255) NOT NULL,
	created_at datetime default NOW(),
	updated_at datetime default NOW() on update NOW(),
	UNIQUE KEY unique_merchant_name(merchant_name)
);

CREATE TABLE merchant_details (
	merchant_id integer PRIMARY KEY,
	merchant_address text,
	phone_number text,
	display_profile_url text,
	FOREIGN KEY (merchant_id) REFERENCES merchants(id)
);

CREATE TABLE banks (
	id integer PRIMARY KEY AUTO_INCREMENT,
	bank_name VARCHAR(50),
	code integer
);

CREATE TABLE merchant_banks (
	id integer PRIMARY KEY AUTO_INCREMENT,
	merchant_id integer,
	bank_id integer,
	on_behalf_of text,
	bank_number text,
	FOREIGN KEY (merchant_id) REFERENCES merchants(id),
	FOREIGN KEY (bank_id) REFERENCES banks(id)
);

CREATE TABLE admins (
	id integer PRIMARY KEY AUTO_INCREMENT,
	merchant_id integer NOT NULL,
	username varchar(50),
	admin_password text,
	email varchar(255),
	created_at datetime default NOW(),
	updated_at datetime default NOW() on update NOW(),
	FOREIGN KEY (merchant_id) REFERENCES merchants(id),
	UNIQUE KEY unique_username(username),
	UNIQUE KEY unique_email(email)
);

CREATE TABLE refresh_tokens (
	id integer PRIMARY KEY AUTO_INCREMENT,
	admin_id integer,
	token text, 
	is_valid bool,
	created_at datetime default NOW(),
	FOREIGN KEY (admin_id) REFERENCES admins(id)
);