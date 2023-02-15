-- insert into user (id, name, email, login, password, location_lat, location_lng, location_address, location_city, location_state, location_zip_code, created_at, updated_at) values  ('7954dfe7-0da2-458e-8fc1-5bfd6d627b7f', 'John Doe', 'john.doe@vibbra.com.br', 'john.doe', '$2a$10$n3/j01lxnX/BNydBNC//GewW5VAkO2jVNzio2uNnQJfukTCk4AYHa', -23.5506507, -46.6333824, 'Rua Vergueiro, 3185', 'São Paulo', 'SP', 65000000, '2023-01-25 18:43:04', '2023-01-25 18:43:04');

create table deal
(
    id                 varchar(36)  not null primary key,
    user_id            varchar(36)  not null,
    type               INT,
    value              decimal(10, 2) not null,
    description        VARCHAR(500),
    trade_for          VARCHAR(100),
    location_lat       double       null,
    location_lng       double       null,
    location_address   varchar(255) not null,
    location_city      varchar(255) not null,
    location_state     varchar(255) not null,
    location_zip_code  bigint       not null,
    urgency_type       INT         not null,
    urgency_limit_date datetime     not null,
    photos             VARCHAR(500),
    created_at         datetime     not null,
    updated_at         datetime     not null,
    constraint fk_deal_user
        foreign key (user_id) references user (id)
);