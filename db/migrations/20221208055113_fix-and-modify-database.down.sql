ALTER TABLE merchant_details RENAME COLUMN display_profile_url TO diplay_profile_url;

ALTER TABLE merchant_details  ADD COLUMN phone_number text;

ALTER TABLE customer_details RENAME COLUMN display_profile_url TO diplay_profile_url;

ALTER TABLE notifications 
DROP FOREIGN KEY FK_notification_type;

ALTER TABLE notifications 
DROP COLUMN notification_type_id;

DROP TABLE notification_types;