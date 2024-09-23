DROP TABLE IF EXISTS sys_user;
CREATE TABLE IF NOT EXISTS sys_user (
    id bigint auto_increment comment '主键',
    username varchar(50) not null comment '用户名',
    password varchar(50) not null comment '密码',
    nickname varchar(50) default '' not null comment '昵称',
    status tinyint default 1 not null comment '状态(1:正常,0:禁用)',
    sort int default 1 not null comment '排序',
    remark varchar(255) null comment '备注',
    roles varchar(255) null comment '角色',
    create_time datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    update_time datetime null on update CURRENT_TIMESTAMP comment '修改时间',
    constraint AK_username unique (username),
    primary key (id)
) comment '用户信息';