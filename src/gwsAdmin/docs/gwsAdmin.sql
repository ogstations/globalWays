drop table `gws_admin`
    DROP TABLE IF EXISTS `gws_admin`

drop table `gws_adminPanel`
    DROP TABLE IF EXISTS `gws_adminPanel`

drop table `gws_menu`
    DROP TABLE IF EXISTS `gws_menu`

drop table `gws_role`
    DROP TABLE IF EXISTS `gws_role`

drop table `gws_logs`
    DROP TABLE IF EXISTS `gws_logs`

create table `gws_admin`
    -- --------------------------------------------------
    --  Table Structure for `gwsAdmin/models.Admin`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `gws_admin` (
        `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `user_name` varchar(255) NOT NULL UNIQUE,
        `password` varchar(255) NOT NULL,
        `email` varchar(32) NOT NULL,
        `mobile` varchar(255) NOT NULL,
        `real_name` varchar(32) NOT NULL,
        `role_id` bigint NOT NULL,
        `lastLogin_ip` varchar(32) NOT NULL,
        `lastLogin_time` datetime NOT NULL,
        `lang` char(5) NOT NULL,
        `status` bool NOT NULL,
        `created` datetime NOT NULL,
        `updated` datetime NOT NULL
    ) ENGINE=InnoDB;
    CREATE INDEX `gws_admin_role_id` ON `gws_admin` (`role_id`);

create table `gws_adminPanel`
    -- --------------------------------------------------
    --  Table Structure for `gwsAdmin/models.AdminPanel`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `gws_adminPanel` (
        `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `mid` bigint NOT NULL,
        `aid` bigint NOT NULL,
        `name` varchar(40) NOT NULL,
        `url` varchar(100) NOT NULL,
        `created` datetime NOT NULL,
        `updated` datetime NOT NULL
    ) ENGINE=InnoDB;

create table `gws_menu`
    -- --------------------------------------------------
    --  Table Structure for `gwsAdmin/models.Menu`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `gws_menu` (
        `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `pid` bigint NOT NULL,
        `name` varchar(40) NOT NULL,
        `url` varchar(100) NOT NULL,
        `data` varchar(60) NOT NULL,
        `order` bigint NOT NULL,
        `display` bool NOT NULL
    ) ENGINE=InnoDB;

    INSERT INTO `gws_menu` (`id`, `pid`, `name`, `url`, `data`, `order`, `display`) VALUES
    (1, 0, '我的面板', 'panel', '', 10000, 1),
    (2, 0, '设置', 'setting', '', 20000, 1),
    (3, 0, '会员', 'member', '', 30000, 1),
    (4, 0, '会员卡', 'memberCard', '', 40000, 1),
    (5, 2, '菜单设置', 'javascript:;', '', 20100, 1),
    (6, 5, '菜单管理', 'menu', '', 20101, 1),
    (7, 1, '个人设置', 'javascript:;', '', 10100, 1),
    (8, 7, '个人信息', 'editInfo', '', 10101, 1),
    (9, 7, '修改密码', 'editPwd', '', 10102, 1),
    (10, 2, '管理员管理', 'javascript:;', '', 20200, 1),
    (11, 10, '管理员管理', 'admin', '', 20201, 1),
    (12, 10, '角色管理', 'role', '', 20202, 1),
    (13, 2, '日志管理', 'javascript:;', '', 20300, 1),
    (14, 13, '日志管理', 'logs', '', 20301, 1),
    (15, 1, '快捷面板', 'javascript:;', '', 10200, 1);

create table `gws_role`
    -- --------------------------------------------------
    --  Table Structure for `gwsAdmin/models.Role`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `gws_role` (
        `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `role_name` varchar(255) NOT NULL UNIQUE,
        `desc` varchar(255) NOT NULL,
        `data` longtext NOT NULL,
        `status` bool NOT NULL,
        `created` datetime NOT NULL,
        `updated` datetime NOT NULL
    ) ENGINE=InnoDB;


INSERT INTO `gws_role` (`id`, `role_name`, `desc`, `data`, `status`, `created`, `updated`) VALUES
(1, '超级管理员', '超级管理员', '1,11,12,13,2,9,10,14,15,16,17,18,3,4,5,6,7,8', 1, '2014-06-28 20:20:09', '2014-06-28 20:20:09');


create table `gws_logs`
    -- --------------------------------------------------
    --  Table Structure for `gwsAdmin/models.Logs`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `gws_logs` (
        `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `uid` bigint NOT NULL,
        `module` varchar(50) NOT NULL,
        `action` varchar(100) NOT NULL,
        `ip` varchar(15) NOT NULL,
        `desc` longtext NOT NULL,
        `created` datetime NOT NULL
    ) ENGINE=InnoDB;


    --创始人
    INSERT INTO `gws_admin` (`user_name`, `password`, `email`, `mobile`, `real_name`, `role_id`, `lastLogin_ip`, `lastLogin_time`, `lang`, `status`, `created`, `updated`) VALUES
    ('admin', '$2a$10$yNj7fzAZ5J6EmEW17q7R7OaE7bRF1a3FvpXgr3l/QGTLTYFm2Apq2', 'zi__chen@163.com', '18610889275', '赵铭', 1, '127.0.0.1', '2014-06-28 15:00:20', 'zh-CN', true, '2014-10-21 20:40:20', '2014-10-21 20:40:20');

