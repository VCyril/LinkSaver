CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE lists
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255)
);

CREATE TABLE users_lists
(
    id      serial                                      not null unique,
    user_id int references users (id) on delete cascade not null,
    list_id int references lists (id) on delete cascade not null
);

CREATE TABLE links
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255)
);

CREATE TABLE lists_links
(
    id      serial                                      not null unique,
    list_id int references lists (id) on delete cascade not null,
    link_id int references links (id) on delete cascade not null
);