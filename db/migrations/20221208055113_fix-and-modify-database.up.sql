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

ALTER TABLE customer_details RENAME COLUMN diplay_profile_url TO display_profile_url;

ALTER TABLE merchant_details  DROP COLUMN phone_number;

ALTER TABLE merchant_details RENAME COLUMN diplay_profile_url TO display_profile_url;