create table menu
(
    id          int auto_increment
        primary key,
    name        varchar(255)      not null,
    label       varchar(255)      null,
    description varchar(255)      null,
    sequence    int     default 0 not null,
    type        varchar(10)       not null comment 'page/button',
    path        varchar(255)      null,
    status      tinyint default 1 comment '0.disable, 1.enable',
    parent_id   int               not null,
    created_at  datetime          not null,
    updated_at  datetime          not null,
    deleted_at  datetime          null
);

create table role
(
    id          int auto_increment
        primary key,
    name        varchar(255)      not null,
    description varchar(255)      null,
    sequence    int     default 0 not null,
    status      tinyint default 1 comment '0.disable, 1.enable',
    parent_id   int               not null,
    created_at  datetime          not null,
    updated_at  datetime          not null,
    deleted_at  datetime          null
);

create table role_menu
(
    id         int auto_increment
        primary key,
    role_id    int      not null,
    menu_id    int      not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime null
);

create table user
(
    id         int auto_increment
        primary key,
    username   varchar(255) not null,
    password   varchar(255) not null,
    status     tinyint default 1 comment '0.disable, 1.enable',
    created_at datetime     not null,
    updated_at datetime     not null,
    deleted_at datetime     null
);

create table user_role
(
    id         int
        primary key,
    user_id    int      not null,
    role_id    int      not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime null
);

create table menu_resource
(
    id         int
        primary key,
    menu_id    int      not null,
    method     varchar(255),
    path       varchar(255),
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime null
);

