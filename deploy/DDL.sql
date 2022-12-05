create table if not exists comment
(
    id          int auto_increment
    primary key,
    context     tinytext      not null,
    user_id     int           not null,
    moment_id   int           not null,
    create_time datetime      not null,
    update_time datetime      not null,
    status      int default 1 not null
)
    engine = InnoDB
    collate = utf8mb4_bin;

create table if not exists moment
(
    id          int auto_increment
    primary key,
    context     text          not null,
    `like`      int default 0 not null,
    user_id     int           not null,
    create_time datetime      not null,
    update_time datetime      not null,
    status      int default 1 not null
)
    engine = InnoDB
    collate = utf8mb4_bin;

create table if not exists thumb_up
(
    id          int auto_increment
    primary key,
    moment_id   int           not null,
    user_id     int           not null,
    create_time datetime      not null,
    update_time datetime      not null,
    status      int default 1 not null
)
    engine = InnoDB
    collate = utf8mb4_bin;

create table if not exists user
(
    id          int auto_increment
    primary key,
    email       varchar(255)  not null,
    name        varchar(255)  not null,
    password    varchar(255)  not null,
    create_time datetime      not null,
    update_time datetime      not null,
    status      int default 1 not null,
    constraint email_unique
    unique (email)
    )
    engine = InnoDB
    collate = utf8mb4_bin;

