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

INSERT INTO contacts VALUES (1,"Rachel", "Green");
INSERT INTO addresses VALUES (1, 1, "Home", "New Haven", "Cook Hill Road", "312", "9");
INSERT INTO addresses VALUES (2, 1, "Work", "San Jose", "Friendship Lane", "4746", "42");
INSERT INTO phones VALUES (1, 1, "Mobile", "050-1234567");
INSERT INTO phones VALUES (2, 1, "Mobile", "051-7654321");

INSERT INTO contacts VALUES (2, "Monica", "Geller");
INSERT INTO addresses VALUES (3, 2, "Work", "Tomato", "cucmber avenue", "12", "34");
INSERT INTO addresses VALUES (4, 2, "Home", "New York", "High road", "312", "10");
INSERT INTO phones VALUES (3, 2, "Mobile", "055-5555555");
INSERT INTO phones VALUES (4, 2, "Mobile", "051-1111111");

INSERT INTO contacts VALUES (3, "Phoebe", "Buffay-Hannigan");
INSERT INTO addresses VALUES (5, 3, "Work", "Massage", "stam", "12", "34");
INSERT INTO phones VALUES (5, 3, "Mobile", "052-2222222");

INSERT INTO contacts VALUES (4, "Joey", "Tribbiani");
INSERT INTO addresses VALUES (6, 4, "Home", "TV", "series", "12", "34");
INSERT INTO phones VALUES (6, 4, "Mobile", "056-6666666");

INSERT INTO contacts VALUES (5, "Chandler", "Bing");
INSERT INTO addresses VALUES (7, 5, "Home", "TV", "series", "12", "34");
INSERT INTO phones VALUES (7, 5, "Mobile", "054-4444444");

INSERT INTO contacts VALUES (6, "Ross", "Geller");

INSERT INTO contacts VALUES (7, "Marshall", "Eriksen");
INSERT INTO addresses VALUES (8, 7, "Home", "Marathon", "42k", "12", "34");
INSERT INTO phones VALUES (8, 7, "Mobile", "053-3333333");

INSERT INTO contacts VALUES (8, "Ted", "Mosby");
INSERT INTO contacts VALUES (9, "Barney", "Stinson");
INSERT INTO contacts VALUES (10, "Lily", "Aldrin");
INSERT INTO contacts VALUES (11, "Leonard", "Hofstadter");
INSERT INTO contacts VALUES (12, "Sheldon", "Cooper");
INSERT INTO contacts VALUES (13, "Penny", "Hofstadter");
INSERT INTO contacts VALUES (14, "Rajesh", "Koothrappali");
INSERT INTO contacts VALUES (15, "Leslie", "Winkle");
INSERT INTO contacts VALUES (16, "Bernadette", "Rostenkowski");
INSERT INTO contacts VALUES (17, "Amy", "Farrah Fowler");
INSERT INTO contacts VALUES (18, "Stuart", "Bloom");
INSERT INTO contacts VALUES (19, "Frank", "Costanza");
INSERT INTO contacts VALUES (20, "James", "Wilson");
INSERT INTO contacts VALUES (21, "Gregory", "House");
INSERT INTO contacts VALUES (22, "Lisa", "Cuddy");
INSERT INTO contacts VALUES (23, "Eric", "Foreman");
INSERT INTO contacts VALUES (24, "Robert", "Chase");
INSERT INTO contacts VALUES (25, "Allison", "Cameron");
INSERT INTO contacts VALUES (26, "Chris", "Taub");