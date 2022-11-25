CREATE TABLE payment_types (
	id integer PRIMARY KEY AUTO_INCREMENT,
	payment_type_name VARCHAR(15)
);

CREATE TABLE payment_statuses (
	id integer PRIMARY KEY AUTO_INCREMENT,
	status_name VARCHAR(15)
);

CREATE TABLE invoices (
	id integer PRIMARY KEY AUTO_INCREMENT,
	merchant_id integer NOT NULL,
	customer_id integer NOT NULL,
	payment_type_id integer NOT NULL,
	payment_status_id integer NOT NULL,
	merchant_bank_id integer,
	due_at datetime NOT NULL,
	approval_document_url text,
	created_at datetime default NOW(),
	updated_at datetime default NOW() on update NOW(),
	FOREIGN KEY (merchant_id) REFERENCES merchants(id),
	FOREIGN KEY (customer_id) REFERENCES customers(id),
	FOREIGN KEY (payment_type_id) REFERENCES payment_types(id),
	FOREIGN KEY (payment_status_id) REFERENCES payment_statuses(id),
	FOREIGN KEY (merchant_bank_id) REFERENCES merchant_banks(id)
);

CREATE TABLE invoice_details (
	id integer PRIMARY KEY AUTO_INCREMENT,
	invoice_id integer NOT NULL,
	product text,
	quantity integer,
	price bigint,
	created_at datetime default NOW(),
	updated_at datetime default NOW() on update NOW(),
	FOREIGN KEY (invoice_id) REFERENCES invoices(id)
);