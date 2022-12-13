ALTER TABLE invoices MODIFY COLUMN due_at date;

INSERT INTO payment_statuses(id, status_name) 
VALUES (1, "Unpaid"),(2, "Pending"),(3, "Paid");

INSERT INTO payment_types(id, payment_type_name)
VALUES (1, "Manual Transfer");
