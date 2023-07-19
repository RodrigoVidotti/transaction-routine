CREATE TABLE Account (
	account_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    document_number INT
);

INSERT INTO `transaction_routine`.`Account` (`document_number`) VALUES (1111111111);
INSERT INTO `transaction_routine`.`Account` (`document_number`) VALUES (987654321);
INSERT INTO `transaction_routine`.`Account` (`document_number`) VALUES (123456789);

CREATE TABLE OperationType (
	operation_type_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(20)
);

INSERT INTO `transaction_routine`.`OperationType` (`name`) VALUES ('COMPRA A VISTA');
INSERT INTO `transaction_routine`.`OperationType` (`name`) VALUES ('COMPRA PARCELADA');
INSERT INTO `transaction_routine`.`OperationType` (`name`) VALUES ('SAQUE');
INSERT INTO `transaction_routine`.`OperationType` (`name`) VALUES ('PAGAMENTO');

CREATE TABLE Transaction (
	transaction_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    account_id INT NOT NULL,
    operation_type_id INT NOT NULL,
    amount DECIMAL(15,2),
    event_datetime DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (account_id) REFERENCES Account(account_id),
    FOREIGN KEY (operation_type_id) REFERENCES OperationType(operation_type_id)
);
