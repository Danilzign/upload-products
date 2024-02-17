CREATE TABLE products
(
    id            serial       not null unique,
    name          varchar(255) not null,
    price         serial not null,
    amount        serial not null,
    imageName        varchar(255) not null
);
