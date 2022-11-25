CREATE TABLE notifications (
	id integer AUTO_INCREMENT PRIMARY KEY,
	customer_id integer,
	title varchar(25),
	content text NOT NULL,
	created_at datetime default NOW(),
	FOREIGN KEY (customer_id) REFERENCES customers(id)
);