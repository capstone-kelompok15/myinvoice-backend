INSERT INTO notification_types(id,name) VALUES (1,"info");
INSERT INTO notification_types(id,name) VALUES (2,"payment");
INSERT INTO notification_types(id,name) VALUES (3,"invoice");

INSERT INTO notification_titles(id,name,notification_type_id) VALUES (1,"Payment is Overdue",1);
INSERT INTO notification_titles(id,name,notification_type_id) VALUES (2,"Payment is Due Soon",1);
INSERT INTO notification_titles(id,name,notification_type_id) VALUES (3,"Payment Success",2);
INSERT INTO notification_titles(id,name,notification_type_id) VALUES (4,"Payment Failed",2);
INSERT INTO notification_titles(id,name,notification_type_id) VALUES (5,"Payment Pending",2);
INSERT INTO notification_titles(id,name,notification_type_id) VALUES (6,"Payment Done",2);
INSERT INTO notification_titles(id,name,notification_type_id) VALUES (7,"New Bill Issued",3);