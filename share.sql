

CREATE TABLE IF NOT EXISTS users(
    id UUID NOT NULL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(500) NOT NULL,
    phone VARCHAR(50)
);


DROP TABLE IF EXISTS users;