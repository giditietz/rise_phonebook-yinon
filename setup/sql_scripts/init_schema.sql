DROP TABLE IF EXISTS addresses;
DROP TABLE IF EXISTS phones;
DROP TABLE IF EXISTS contacts;

CREATE TABLE contacts (
    contact_id INT AUTO_INCREMENT NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    PRIMARY KEY (contact_id)
);

CREATE TABLE addresses (
    address_id INT AUTO_INCREMENT NOT NULL,
    contact_id INT NOT NULL,
    description VARCHAR(10),
    city VARCHAR(40) NOT NULL,
    street VARCHAR(20),
    home_number VARCHAR(20),
    apartment VARCHAR(20),

    PRIMARY KEY (address_id),
    FOREIGN KEY (contact_id)
        REFERENCES contacts(contact_id)
        ON DELETE CASCADE
);

CREATE TABLE phones (
    phone_id INT AUTO_INCREMENT NOT NULL,
    contact_id INT NOT NULL,
    description VARCHAR(10),
    phone_number VARCHAR(15) NOT NULL,

    PRIMARY KEY (phone_id),
    FOREIGN KEY (contact_id)
        REFERENCES contacts(contact_id)
        ON DELETE CASCADE
);

INSERT INTO contacts VALUES (1, "Yinon", "Yishay");
INSERT INTO contacts VALUES (2, "Talya", "Yishay");
INSERT INTO addresses VALUES (1, 1, "Home", "Revava", "bla", "10", "9");
INSERT INTO addresses VALUES (2, 1, "Work", "Tel Aviv", "bla", "10", "9");
INSERT INTO phones VALUES (1, 1, "Mobile", "052-8119308");
INSERT INTO phones VALUES (2, 1, "Mobile", "052-8119308");
INSERT INTO phones VALUES (3, 1, "Mobile", "052-8119308")