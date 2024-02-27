CREATE TABLE products
(
    id            uuid DEFAULT gen_random_uuid(),
    name          varchar(255) not null,
    price         serial not null,
    quantity      serial not null,
    image_name    varchar(255) not null
);
