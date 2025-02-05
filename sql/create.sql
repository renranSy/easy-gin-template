create table admin
(
    id         int auto_increment
        primary key,
    username   varchar(255) default '' not null,
    password   varchar(255) default '' not null,
    email      varchar(255) default '' not null,
    status     tinyint      default 1  not null comment '0.disable, 1.enable',
    created_at datetime                not null,
    updated_at datetime                not null,
    deleted_at datetime                null
);

create table admin_role
(
    id         int auto_increment
        primary key,
    admin_id   int      not null,
    role_id    int      not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime null
);

create table menu
(
    id          int auto_increment
        primary key,
    name        varchar(255) default '' not null,
    code        varchar(255) default '' not null,
    description varchar(255) default '' not null,
    sequence    int          default 0  not null,
    type        varchar(10)  default '' not null comment 'page/button',
    path        varchar(255) default '' not null,
    status      tinyint      default 1  not null comment '0.disable, 1.enable',
    parent_id   int          default 0  not null,
    created_at  datetime                not null,
    updated_at  datetime                not null,
    deleted_at  datetime                null
);

create table menu_resource
(
    id         int auto_increment
        primary key,
    menu_id    int                     not null,
    method     varchar(255) default '' not null,
    path       varchar(255) default '' not null,
    created_at datetime                not null,
    updated_at datetime                not null,
    deleted_at datetime                null
);

create table role
(
    id          int auto_increment
        primary key,
    name        varchar(255) default '' not null,
    description varchar(255) default '' not null,
    sequence    int          default 0  not null,
    status      tinyint      default 1  not null comment '0.disable, 1.enable',
    created_at  datetime                not null,
    updated_at  datetime                not null,
    deleted_at  datetime                null
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

