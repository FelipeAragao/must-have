create table user
(
    id                binary(16)   not null primary key,
    name              varchar(255) not null,
    email             varchar(255) not null,
    login             varchar(255) not null,
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

