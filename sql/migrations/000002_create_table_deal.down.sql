create table user
(
    id                varchar(36)   not null primary key,
    name              varchar(255) not null,
    email             varchar(255) not null unique,
    login             varchar(255) not null unique,
    password          varchar(255) not null,
    location_lat      double       null,
    location_lng      double       null,
    location_address  varchar(255) not null,
    location_city     varchar(255) not null,
    location_state    varchar(255) not null,
    location_zip_code bigint       not null,
    created_at        datetime     not null,
    updated_at        datetime     not null
);

