
DROP TABLE IF EXISTS customer;

CREATE TABLE customer (
  id     INTEGER PRIMARY KEY,
  firstname VARCHAR(100) NOT NULL,
  lastname   VARCHAR(100) NOT NULL,
  birthdate  DATE NOT NULL,
  gender     VARCHAR(6) NOT NULL CHECK( gender = 'Male' OR gender = 'Female'),
  email      VARCHAR(255) NOT NULL,
  address    VARCHAR(200)
);

DROP SEQUENCE IF EXISTS customer_id_seq;
CREATE SEQUENCE customer_id_seq START 101;

INSERT INTO customer VALUES( nextval('customer_id_seq'), 'Rustam', 'Novikov', to_date('13.11.1980', 'DD.MM.YYYY'), 'Male', 'novikovrustam@gmail.com', 'Muuga');