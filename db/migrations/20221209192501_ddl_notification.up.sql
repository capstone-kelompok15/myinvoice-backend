
DROP TABLE notifications;
DROP TABLE notification_types;

CREATE TABLE notification_types (
	id integer PRIMARY KEY,
	name VARCHAR(255) NOT NULL
);
CREATE TABLE notification_titles (
	id integer PRIMARY KEY,  
	name VARCHAR(255) NOT NULL,
    notification_type_id integer,
	FOREIGN KEY (notification_type_id) REFERENCES notification_types(id)
);
CREATE TABLE customer_notifications (
	id integer AUTO_INCREMENT PRIMARY KEY,
	to_customer_id integer ,
	from_merchant_id integer ,
	invoice_id integer ,
	notification_title_id integer ,
	is_read boolean DEFAULT false,
	created_at datetime default NOW(),
	FOREIGN KEY (to_customer_id) REFERENCES customer_details(customer_id),
	FOREIGN KEY (from_merchant_id) REFERENCES merchants(id),
	FOREIGN KEY (invoice_id) REFERENCES invoices(id),
	FOREIGN KEY (notification_title_id) REFERENCES notification_titles(id)
);
CREATE TABLE merchant_notifications (
	id integer AUTO_INCREMENT PRIMARY KEY,
	from_customer_id integer ,
	to_merchant_id integer ,
	invoice_id integer ,
	notification_title_id integer ,
	is_read boolean DEFAULT false,
	created_at datetime default NOW(),
	FOREIGN KEY (from_customer_id) REFERENCES customer_details(customer_id),
	FOREIGN KEY (to_merchant_id) REFERENCES merchants(id),
	FOREIGN KEY (invoice_id) REFERENCES invoices(id),
	FOREIGN KEY (notification_title_id) REFERENCES notification_titles(id)
);

 