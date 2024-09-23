create table rbac_user
(
    id          bigint                             auto_increment comment '主键',
    account     char(20)                           not null comment '账号',
    user_name   varchar(50) default ''             not null comment '昵称',
    password    varchar(64) charset utf8mb3        not null comment '密码',
    roles       varchar(255)                       null comment '角色',
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time datetime                           null on update CURRENT_TIMESTAMP comment '修改时间',
    constraint  AK_account unique (account),
    primary key (id)
) comment '用户信息';
