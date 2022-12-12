DROP TABLE merchant_notifications;
DROP TABLE customer_notifications;
DROP TABLE notification_titles ;
DROP TABLE notification_types ;
 
CREATE TABLE notifications (
	id integer AUTO_INCREMENT PRIMARY KEY,
	customer_id integer,
	title varchar(25),
	content text NOT NULL,
	created_at datetime default NOW(),
	FOREIGN KEY (customer_id) REFERENCES customers(id)
);

CREATE TABLE notification_types (
	id int PRIMARY KEY AUTO_INCREMENT,
	name text,
	created_at datetime DEFAULT NOW()
);

ALTER TABLE notifications 
ADD COLUMN notification_type_id int NOT NULL;

ALTER TABLE notifications 
ADD CONSTRAINT FK_notification_type 
FOREIGN KEY (notification_type_id) 
REFERENCES notification_types(id);