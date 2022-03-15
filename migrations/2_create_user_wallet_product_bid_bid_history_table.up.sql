CREATE TABLE users (
    id serial primary key,
    email varchar(255) unique not null,
    name varchar(255),
    password varchar(255)
);

CREATE TABLE wallet (
    id serial primary key,
    user_id int,
    balance int not null default 0,
    constraint fk_users
        foreign key(user_id)
            references users(id)
);

CREATE TABLE products (
    id serial primary key,
    name varchar(255),
    initial_price int not null default 0,
    start_bid_date timestamp,
    end_bid_date timestamp,
    owner_user_id int,
    last_bid_id int,
    image_url varchar(1024),
    status int not null default 0,
    bid_increment int not null default 0,
    constraint fk_users
        foreign key(owner_user_id)
            references users(id)
);

CREATE TABLE bid(
    id serial primary key,
    product_id int,
    user_id int,
    amount int not null default 0,
    status int not null default 0,
    constraint fk_users
        foreign key(user_id)
            references users(id),
    constraint fk_products
        foreign key(product_id)
            references products(id)
);

CREATE TABLE bid_history(
    id serial primary key,
    product_id int,
    user_id int,
    amount int not null default 0,
    status int not null default 0,
    bid_time timestamp,
    constraint fk_users
        foreign key(user_id)
            references users(id),
    constraint fk_products
        foreign key(product_id)
            references products(id)
);