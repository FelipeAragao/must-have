create table if not exists users
(
    id                  binary(16)   not null
        primary key,
    name                varchar(255) null,
    email               varchar(255) null,
    login               varchar(255) null,
    `location.lat`      double       null,
    `location.lng`      double       null,
    `location.address`  varchar(255) null,
    `location.city`     varchar(255) null,
    `location.state`    varchar(255) null,
    `location.zip_code` bigint       null,
    created_at          datetime     null,
    updated_at          datetime     null
);

