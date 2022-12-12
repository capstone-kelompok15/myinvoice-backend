DELETE FROM payment_types;

DELETE FROM payment_statuses;

ALTER TABLE invoices MODIFY COLUMN due_at datetime;