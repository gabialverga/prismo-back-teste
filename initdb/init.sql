CREATE TABLE accounts (
  account_id INT AUTO_INCREMENT PRIMARY KEY,
  document_number VARCHAR(20) NOT NULL UNIQUE
);

CREATE TABLE operationTypes (
  operationType_ID INT AUTO_INCREMENT PRIMARY KEY,
  description VARCHAR(20) NOT NULL UNIQUE
);

CREATE TABLE transactions (
  transaction_id INT AUTO_INCREMENT PRIMARY KEY,
  account_id INT NOT NULL,
  operationType_ID INT NOT NULL,
  amount FLOAT NOT NULL,
  eventDate TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (account_id) REFERENCES accounts (account_id),
  FOREIGN KEY (operationType_ID) REFERENCES operationTypes (operationType_ID)
);

INSERT INTO operationTypes (description) VALUES ('COMPRA A VISTA');
INSERT INTO operationTypes (description) VALUES ('COMPRA PARCELADA');
INSERT INTO operationTypes (description) VALUES ('SAQUE');
INSERT INTO operationTypes (description) VALUES ('PAGAMENTO');
