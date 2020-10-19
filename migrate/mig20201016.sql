CREATE TABLE `dcmc_area_pc` (
                                `id` int(11) NOT NULL AUTO_INCREMENT,
                                `area_id` int(11) NOT NULL,
                                `pc_id` int(11) NOT NULL,
                                `area_name` varchar(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                `status` tinyint(3) DEFAULT NULL,
                                `need_upgrade` tinyint(3) DEFAULT NULL,
                                `machine_id` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                `created_at` datetime DEFAULT NULL,
                                `updated_at` datetime DEFAULT NULL,
                                PRIMARY KEY (`id`),
                                KEY `area_id` (`area_id`),
                                KEY `pc_id` (`pc_id`),
                                KEY `created_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=123 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



CREATE TABLE `dcmc_area_software` (
                                      `id` int(11) NOT NULL AUTO_INCREMENT,
                                      `area_id` int(11) NOT NULL,
                                      `software_id` int(11) NOT NULL,
                                      `area_name` varchar(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                      `name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                      `package_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                      `version` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                      `architecture` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                      `shelf` tinyint(3) DEFAULT NULL,
                                      `icon` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                      `created_at` datetime DEFAULT NULL,
                                      `updated_at` datetime DEFAULT NULL,
                                      PRIMARY KEY (`id`),
                                      KEY `area_id` (`area_id`),
                                      KEY `software_id` (`software_id`),
                                      KEY `created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;



CREATE TABLE `dcmc_area_user` (
                                  `id` int(11) NOT NULL AUTO_INCREMENT,
                                  `area_id` int(11) NOT NULL,
                                  `user_id` int(11) NOT NULL,
                                  `area_name` varchar(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                  `utype` tinyint(3) DEFAULT NULL,
                                  `ldap_username` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                  `full_name` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                  `job_number` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                  `status` tinyint(3) DEFAULT NULL,
                                  `created_at` datetime DEFAULT NULL,
                                  `updated_at` datetime DEFAULT NULL,
                                  PRIMARY KEY (`id`),
                                  KEY `area_id` (`area_id`),
                                  KEY `user_id` (`user_id`),
                                  KEY `created_at` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE `dcmc_asset` (
                              `id` int(11) NOT NULL AUTO_INCREMENT,
                              `cpu` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'CPU硬件码',
                              `hash` char(64) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT 'sha256摘要',
                              `host_account` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '主机账号',
                              `create_admin_id` smallint(5) DEFAULT NULL COMMENT '创建管理ID',
                              `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                              PRIMARY KEY (`id`),
                              KEY `idx_cpu` (`cpu`),
                              KEY `idx_hash` (`hash`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE `dcmc_scene` (
                              `id` int(11) NOT NULL AUTO_INCREMENT,
                              `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
                              `code` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
                              `sort` int(11) DEFAULT 0,
                              `created_at` datetime DEFAULT NULL,
                              `updated_at` datetime DEFAULT NULL,
                              PRIMARY KEY (`id`),
                              KEY `sort` (`sort`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE `dcmc_sys_config` (
                                   `id` int(11) NOT NULL AUTO_INCREMENT,
                                   `sys_key` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
                                   `sys_value` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                   `sys_group` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT 'System',
                                   `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                   `created_at` datetime DEFAULT NULL,
                                   `updated_at` datetime DEFAULT NULL,
                                   PRIMARY KEY (`id`),
                                   KEY `sys_group` (`sys_group`),
                                   KEY `sys_key` (`sys_key`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


CREATE TABLE `dcmc_admin_user` (
                                      `admin_id` int(11) NOT NULL,
                                      `user_id` int(11) NOT NULL,
                                      PRIMARY KEY (`admin_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



-- ----------------------------
-- Records of dcmc_admin
-- ----------------------------
-- INSERT INTO `dcmc_admin` VALUES (1, '00', '', 'admin', '33275a8aa48ea918bd53a9181aa975f15ab0d0645398f5918a006d08675c1cb27d5c645dbd084eee56e675e25ba4019f2ecea37ca9e2995b49fcb12c096a032e', '超级管理员', '', '', '', '', '2019-06-05 08:35:54', '2020-08-12 06:42:56');

-- ----------------------------
-- Records of dcmc_auth
-- ----------------------------
-- INSERT INTO `dcmc_auth` VALUES (1, '', '创建子管理员', 'admin', 'post');
-- INSERT INTO `dcmc_auth` VALUES (2, '', '创建子区域', 'area', 'post');

-- ----------------------------
-- Records of dcmc_scene
-- ----------------------------
INSERT INTO `dcmc_scene` VALUES (1, '内网接入', 'intranet-scene', 0, '2020-07-20 01:29:32', '2020-07-20 01:29:32');


-- 软件类别表
-- INSERT INTO `dcmc`.`dcmc_software_category`(`id`,`name`)VALUES(1,'应用软件');
-- INSERT INTO `dcmc`.`dcmc_software_category`(`id`,`name`)VALUES(2,'系统软件');

-- ----------------------------
-- Records of dcmc_sys_config
-- ----------------------------
INSERT INTO `dcmc_sys_config` VALUES (1, 'AssetsPath', '', 'Assets', '内部资产设备清单文件地址', '2020-07-20 01:29:32', '2020-08-11 02:47:03');
INSERT INTO `dcmc_sys_config` VALUES (2, 'DeepinAppStoreUrl', 'http://icbcarm-dstore.chinauos.com/api/public/packages', 'DeepinAppStore', '内网仓库服务地址', '2020-07-20 01:29:32', '2020-08-10 05:58:54');
INSERT INTO `dcmc_sys_config` VALUES (3, 'DeepinAppStoreIconUrl', 'http://icbcarm-developer.chinauos.com/api/public/blob', 'DeepinAppStore', '内网仓库静态资源地址', '2020-07-20 01:29:32', '2020-08-10 05:58:54');
INSERT INTO `dcmc_sys_config` VALUES (4, 'DeepinAppStoreSourceList', '[\"deb [trusted=yes] http://icbcarm-repo.chinauos.com/appstore/ eagle appstore\"]', 'DeepinAppStore', '内网仓库源地址', '2020-07-20 01:29:32', '2020-08-10 05:58:54');
INSERT INTO `dcmc_sys_config` VALUES (5, 'SysAppConfig', '[{\"app_id\":1,\"title\":\"浏览器\",\"package_name\":\"uos-browser-stable\",\"clip_script\":\"apt-get remove -y uos-browser-stable\",\"rec_script\":\"apt-get install -y uos-browser-stable\",\"display_order\":1,\"whether_clip\":0},{\"app_id\":2,\"title\":\"应用商店\",\"package_name\":\"deepin-appstore\",\"clip_script\":\"apt-get remove -y deepin-appstore\",\"rec_script\":\"apt-get install -y deepin-appstore\",\"display_order\":2,\"whether_clip\":0},{\"app_id\":3,\"title\":\"音乐\",\"package_name\":\"deepin-music\",\"clip_script\":\"apt-get remove -y deepin-music\",\"rec_script\":\"apt-get install -y deepin-music\",\"display_order\":3,\"whether_clip\":0},{\"app_id\":4,\"title\":\"影院\",\"package_name\":\"deepin-movie\",\"clip_script\":\"apt-get remove -y deepin-movie\",\"rec_script\":\"apt-get install -y deepin-movie\",\"display_order\":4,\"whether_clip\":0},{\"app_id\":5,\"title\":\"文件管理器\",\"package_name\":\"dde-file-manager\",\"clip_script\":\"apt-get remove -y dde-file-manager\",\"rec_script\":\"apt-get install -y dde-file-manager\",\"display_order\":5,\"whether_clip\":0},{\"app_id\":6,\"title\":\"截图录屏\",\"package_name\":\"deepin-screen-recorder\",\"clip_script\":\"apt-get remove -y deepin-screen-recorder\",\"rec_script\":\"apt-get install -y deepin-screen-recorder\",\"display_order\":6,\"whether_clip\":0},{\"app_id\":7,\"title\":\"看图\",\"package_name\":\"deepin-image-viewer\",\"clip_script\":\"apt-get remove -y deepin-image-viewer\",\"rec_script\":\"apt-get install -y deepin-image-viewer\",\"display_order\":7,\"whether_clip\":0},{\"app_id\":8,\"title\":\"相册\",\"package_name\":\"deepin-album\",\"clip_script\":\"apt-get remove -y deepin-album\",\"rec_script\":\"apt-get install -y deepin-album\",\"display_order\":8,\"whether_clip\":0},{\"app_id\":9,\"title\":\"画板\",\"package_name\":\"deepin-draw\",\"clip_script\":\"apt-get remove -y deepin-draw\",\"rec_script\":\"apt-get install -y deepin-draw\",\"display_order\":9,\"whether_clip\":0},{\"app_id\":10,\"title\":\"文档查看器\",\"package_name\":\"deepin-reader\",\"clip_script\":\"apt-get remove -y deepin-reader\",\"rec_script\":\"apt-get install -y deepin-reader\",\"display_order\":10,\"whether_clip\":0},{\"app_id\":11,\"title\":\"编辑器\",\"package_name\":\"deepin-editor\",\"clip_script\":\"apt-get remove -y deepin-editor\",\"rec_script\":\"apt-get install -y deepin-editor\",\"display_order\":11,\"whether_clip\":0},{\"app_id\":12,\"title\":\"雷鸟邮件\",\"package_name\":\"thunderbird\",\"clip_script\":\"apt-get remove -y thunderbird\",\"rec_script\":\"apt-get install -y thunderbird\",\"display_order\":12,\"whether_clip\":0},{\"app_id\":13,\"title\":\"终端\",\"package_name\":\"gnome-terminal\",\"clip_script\":\"apt-get remove -y gnome-terminal\",\"rec_script\":\"apt-get install -y gnome-terminal\",\"display_order\":13,\"whether_clip\":0},{\"app_id\":14,\"title\":\"联系人\",\"package_name\":\"deepin-contacts\",\"clip_script\":\"apt-get remove -y deepin-contacts\",\"rec_script\":\"apt-get install -y deepin-contacts\",\"display_order\":14,\"whether_clip\":0},{\"app_id\":15,\"title\":\"语音记事本\",\"package_name\":\"deepin-voice-note\",\"clip_script\":\"apt-get remove -y deepin-voice-note\",\"rec_script\":\"apt-get install -y deepin-voice-note\",\"display_order\":15,\"whether_clip\":0},{\"app_id\":16,\"title\":\"帮助手册\",\"package_name\":\"deepin-manual\",\"clip_script\":\"apt-get remove -y deepin-manual\",\"rec_script\":\"apt-get install -y deepin-manual\",\"display_order\":16,\"whether_clip\":0},{\"app_id\":17,\"title\":\"扫描管理\",\"package_name\":\"simple-scan\",\"clip_script\":\"apt-get remove -y simple-scan\",\"rec_script\":\"apt-get install -y simple-scan\",\"display_order\":17,\"whether_clip\":0},{\"app_id\":18,\"title\":\"控制中心\",\"package_name\":\"dde-control-center\",\"clip_script\":\"apt-get remove -y dde-control-center\",\"rec_script\":\"apt-get install -y dde-control-center\",\"display_order\":18,\"whether_clip\":0},{\"app_id\":19,\"title\":\"输入法配置\",\"package_name\":\"fcitx-data\",\"clip_script\":\"apt-get remove -y fcitx-data\",\"rec_script\":\"apt-get install -y fcitx-data\",\"display_order\":19,\"whether_clip\":0},{\"app_id\":20,\"title\":\"安全中心\",\"package_name\":\"deepin-defender\",\"clip_script\":\"apt-get remove -y deepin-defender\",\"rec_script\":\"apt-get install -y deepin-defender\",\"display_order\":20,\"whether_clip\":0},{\"app_id\":21,\"title\":\"系统监视器\",\"package_name\":\" deepin-system-monitor\",\"clip_script\":\"apt-get remove -y  deepin-system-monitor\",\"rec_script\":\"apt-get install -y  deepin-system-monitor\",\"display_order\":21,\"whether_clip\":0},{\"app_id\":22,\"title\":\"启动盘制作工具\",\"package_name\":\"deepin-boot-maker\",\"clip_script\":\"apt-get remove -y deepin-boot-maker\",\"rec_script\":\"apt-get install -y deepin-boot-maker\",\"display_order\":22,\"whether_clip\":0},{\"app_id\":23,\"title\":\"设备管理器\",\"package_name\":\"deepin-devicemanager\",\"clip_script\":\"apt-get remove -y deepin-devicemanager\",\"rec_script\":\"apt-get install -y deepin-devicemanager\",\"display_order\":23,\"whether_clip\":0},{\"app_id\":24,\"title\":\"日志收集工具\",\"package_name\":\"deepin-log-viewer\",\"clip_script\":\"apt-get remove -y deepin-log-viewer\",\"rec_script\":\"apt-get install -y deepin-log-viewer\",\"display_order\":24,\"whether_clip\":0},{\"app_id\":25,\"title\":\"打印管理器\",\"package_name\":\"dde-printer \",\"clip_script\":\"apt-get remove -y dde-printer \",\"rec_script\":\"apt-get install -y dde-printer \",\"display_order\":25,\"whether_clip\":0},{\"app_id\":26,\"title\":\"日历\",\"package_name\":\"dde-calendar\",\"clip_script\":\"apt-get remove -y dde-calendar\",\"rec_script\":\"apt-get install -y dde-calendar\",\"display_order\":26,\"whether_clip\":0},{\"app_id\":27,\"title\":\"计算器\",\"package_name\":\"deepin-calculator\",\"clip_script\":\"apt-get remove -y deepin-calculator\",\"rec_script\":\"apt-get install -y deepin-calculator\",\"display_order\":27,\"whether_clip\":0},{\"app_id\":28,\"title\":\"字体管理器\",\"package_name\":\"deepin-font-manager\",\"clip_script\":\"apt-get remove -y deepin-font-manager\",\"rec_script\":\"apt-get install -y deepin-font-manager\",\"display_order\":28,\"whether_clip\":0},{\"app_id\":29,\"title\":\"归档管理器\",\"package_name\":\"deepin-compressor\",\"clip_script\":\"apt-get remove -y deepin-compressor\",\"rec_script\":\"apt-get install -y deepin-compressor\",\"display_order\":29,\"whether_clip\":0},{\"app_id\":30,\"title\":\"软件包安装器\",\"package_name\":\"deepin-deb-installer\",\"clip_script\":\"apt-get remove -y deepin-deb-installer\",\"rec_script\":\"apt-get install -y deepin-deb-installer\",\"display_order\":30,\"whether_clip\":0},{\"app_id\":31,\"title\":\"分区编辑器\",\"package_name\":\"gparted\",\"clip_script\":\"apt-get remove -y gparted\",\"rec_script\":\"apt-get install -y gparted\",\"display_order\":31,\"whether_clip\":0},{\"app_id\":32,\"title\":\"欢迎\",\"package_name\":\"dde-introduction\",\"clip_script\":\"apt-get remove -y dde-introduction\",\"rec_script\":\"apt-get install -y dde-introduction\",\"display_order\":32,\"whether_clip\":0},{\"app_id\":33,\"title\":\"服务与支持\",\"package_name\":\"uos-service-support\",\"clip_script\":\"apt-get remove -y uos-service-support\",\"rec_script\":\"apt-get install -y uos-service-support\",\"display_order\":33,\"whether_clip\":0},{\"app_id\":35,\"title\":\"授权管理\",\"package_name\":\"deepin-license-activator\",\"clip_script\":\"apt-get remove -y deepin-license-activator\",\"rec_script\":\"apt-get install -y deepin-license-activator\",\"display_order\":35,\"whether_clip\":0},{\"app_id\":37,\"title\":\"桌面智能助手\",\"package_name\":\"deepin-aiassistant\",\"clip_script\":\"apt-get remove -y deepin-aiassistant\",\"rec_script\":\"apt-get install -y deepin-aiassistant\",\"display_order\":37,\"whether_clip\":0},{\"app_id\":38,\"title\":\"茄子摄像头\",\"package_name\":\"cheese\",\"clip_script\":\"apt-get remove -y cheese\",\"rec_script\":\"apt-get install -y cheese\",\"display_order\":38,\"whether_clip\":0}]', 'SystemCustom', '系统预装应用配置清单', '2020-07-20 01:29:32', '2020-07-20 01:29:32');
INSERT INTO `dcmc_sys_config` VALUES (6, 'SysSetting', '[{"checked_option":"禁用","cn_name":"快捷键","disable_script":"gsettings set com.deepin.dde.daemon keybinding false ; qdbus org.kde.kglobalaccel /kglobalaccel blockGlobalShortcuts true","display_order":1,"en_name":"app","enable_script":"gsettings set com.deepin.dde.daemon keybinding true ; qdbus org.kde.kglobalaccel /kglobalaccel blockGlobalShortcuts false","group_name":"","id":1,"options":"启用,禁用","template":0,"tool_tip":"可配置是否允许使用系统快捷键","whether_enable":0,"relation_keys":""},{"checked_option":"高效模式","cn_name":"显示模式","disable_script":"gsettings set com.deepin.dde.dock display-mode efficient","display_order":2,"en_name":"display-mode","enable_script":"gsettings set com.deepin.dde.dock display-mode fashion","group_name":"任务栏配置","id":2,"options":"时尚模式,高效模式","template":0,"tool_tip":"可配置任务栏显示模式 高效模式 时尚模式","whether_enable":0,"relation_keys":""},{"checked_option":"隐藏","cn_name":"启动器","disable_script":"gsettings set com.deepin.dde.dock.module.launcher enable false","display_order":2,"en_name":"launcher","enable_script":"gsettings set com.deepin.dde.dock.module.launcher enable true","group_name":"任务栏配置","id":2,"options":"显示,隐藏","template":0,"tool_tip":"可配置是否显示左下角启动器,显示应用列表","whether_enable":0,"relations_keys":"menu,dockapp,network,disk-mount"},{"checked_option":"禁用","cn_name":"任务栏右键菜单","disable_script":"gsettings set com.deepin.dde.dock.module.menu enable false","display_order":3,"en_name":"menu","enable_script":"gsettings set com.deepin.dde.dock.module.menu enable true","group_name":"任务栏配置","id":3,"options":"启用,禁用","template":0,"tool_tip":"可配置是否可以右键任务栏展开右键菜单","whether_enable":0,"relation_keys":""},{"checked_option":"隐藏","cn_name":"驻留任务栏应用","disable_script":"gsettings set com.deepin.dde.dock.module.dockapp enable false","display_order":4,"en_name":"dockapp","enable_script":"gsettings set com.deepin.dde.dock.module.dockapp enable true","group_name":"任务栏配置","id":4,"options":"显示,隐藏","template":0,"tool_tip":"可配置是否显示驻留在任务栏上的应用","whether_enable":0,"relation_keys":""},{"checked_option":"禁用","cn_name":"网络托盘图标可点击","disable_script":"gsettings set com.deepin.dde.dock.module.network control true","display_order":5,"en_name":"network ","enable_script":"gsettings set com.deepin.dde.dock.module.network control false","group_name":"托盘图标配置","id":5,"options":"启用,禁用","template":0,"tool_tip":"可配置是否可以点击托盘上的网络图标进行网络设置","whether_enable":0,"relation_keys":""},{"checked_option":"禁用","cn_name":"USB托盘图标双击打开访问","disable_script":"gsettings set com.deepin.dde.dock.module.disk-mount filemanager-integration false","display_order":6,"en_name":"disk-mount","enable_script":"gsettings set com.deepin.dde.dock.module.disk-mount filemanager-integration true","group_name":"托盘图标配置","id":6,"options":"启用,禁用","template":0,"tool_tip":"可配置是否可以双击打开USB托盘图标","whether_enable":0,"relation_keys":""},{"checked_option":"隐藏","cn_name":"回收站","disable_script":"gsettings set com.deepin.dde.dock.module.trash enable false","display_order":7,"en_name":"trash","enable_script":"gsettings set com.deepin.dde.dock.module.trash enable true","group_name":"应用配置","id":7,"options":"显示,隐藏","template":0,"tool_tip":"可配置是否显示回收站图标","whether_enable":0,"relation_keys":""},{"checked_option":"显示","cn_name":"账户","disable_script":"","display_order":8,"en_name":"accounts","enable_script":"","group_name":"控制中心配置","id":8,"options":"显示,隐藏","template":1,"tool_tip":"可配置是否显示本地账户一级菜单","whether_enable":1,"relation_keys":""},{"checked_option":"显示","cn_name":"Union ID","disable_script":"","display_order":9,"en_name":"cloudsync,unionid","enable_script":"","group_name":"控制中心配置","id":9,"options":"显示,隐藏","template":1,"tool_tip":"可配置是否显示网络帐户一级菜单","whether_enable":1,"relation_keys":""},{"checked_option":"显示","cn_name":"显示","disable_script":"","display_order":10,"en_name":"display","enable_script":"","group_name":"控制中心配置","id":10,"options":"显示,隐藏","template":1,"tool_tip":"可配置是否显示显示一级菜单","whether_enable":1,"relation_keys":""},{"checked_option":"显示","cn_name":"默认程序","disable_script":"","display_order":11,"en_name":"defapp","enable_script":"","group_name":"控制中心配置","id":11,"options":"显示,隐藏","template":1,"tool_tip":"可配置是否显示默认程序一级菜单","whether_enable":1,"relation_keys":""},{"checked_option":"显示","cn_name":"个性化","disable_script":"","display_order":12,"en_name":"personalization","enable_script":"","group_name":"控制中心配置","id":12,"options":"显示,隐藏","template":1,"tool_tip":"可配置是否显示个性化一级菜单","whether_enable":1,"relation_keys":""},{"checked_option":"显示","cn_name":"网络","disable_script":"","display_order":13,"en_name":"network,","enable_script":"","group_name":"控制中心配置","id":13,"options":"显示,隐藏","template":1,"tool_tip":"可配置是否显示网络一级菜单","whether_enable":1,"relation_keys":""},{"checked_option":"显示","cn_name":"声音","disable_script":"","display_order":14,"en_name":"sound","enable_script":"","group_name":"控制中心配置","id":14,"options":"显示,隐藏","template":1,"tool_tip":"可配置是否显示声音一级菜单","whether_enable":1,"relation_keys":""},{"checked_option":"显示","cn_name":"蓝牙","disable_script":"","display_order":15,"en_name":"bluetooth","enable_script":"","group_name":"控制中心配置","id":15,"options":"显示,隐藏","template":1,"tool_tip":"可配置是否显示蓝牙一级菜单","whether_enable":1,"relation_keys":""},{"checked_option":"显示","cn_name":"日期时间","disable_script":"","display_order":16,"en_name":"datetime","enable_script":"","group_name":"控制中心配置","id":16,"options":"显示,隐藏","template":1,"tool_tip":"可配置是否显示日期时间一级菜单","whether_enable":1,"relation_keys":""},{"checked_option":"显示","cn_name":"通知","disable_script":"","display_order":17,"en_name":"notification","enable_script":"","group_name":"控制中心配置","id":17,"options":"显示,隐藏","template":1,"tool_tip":"可配置是否显示通知一级菜单","whether_enable":1,"relation_keys":""},{"checked_option":"显示","cn_name":"电源管理","disable_script":"","display_order":17,"en_name":"power","enable_script":"","group_name":"控制中心配置","id":17,"options":"显示,隐藏","template":1,"tool_tip":"可配置是否显示电源一级菜单","whether_enable":1,"relation_keys":""},{"checked_option":"显示","cn_name":"鼠标","disable_script":"","display_order":18,"en_name":"mouse","enable_script":"","group_name":"控制中心配置","id":18,"options":"显示,隐藏","template":1,"tool_tip":"可配置是否显示鼠标一级菜单","whether_enable":1,"relation_keys":""},{"checked_option":"显示","cn_name":"数位板","disable_script":"","display_order":18,"en_name":"wacom","enable_script":"","group_name":"控制中心配置","id":18,"options":"显示,隐藏","template":1,"tool_tip":"可配置是否显示数位板一级菜单","whether_enable":1,"relation_keys":""},{"checked_option":"显示","cn_name":"键盘和语言","disable_script":"","display_order":19,"en_name":"keyboard","enable_script":"","group_name":"控制中心配置","id":19,"options":"显示,隐藏","template":1,"tool_tip":"可配置是否显示键盘和语言一级菜单","whether_enable":1,"relation_keys":""},{"checked_option":"显示","cn_name":"辅助功能","disable_script":"","display_order":20,"en_name":"voice","enable_script":"","group_name":"控制中心配置","id":20,"options":"显示,隐藏","template":1,"tool_tip":"可配置是否显示辅助功能一级菜单","whether_enable":1,"relation_keys":""},{"checked_option":"显示","cn_name":"系统信息","disable_script":"","display_order":22,"en_name":"systeminfo","enable_script":"","group_name":"控制中心配置","id":22,"options":"显示,隐藏","template":1,"tool_tip":"可配置是否显示系统信息一级菜单","whether_enable":1,"relation_keys":""},{"checked_option":"显示","cn_name":"通用","disable_script":"","display_order":23,"en_name":"commoninfo","enable_script":"","group_name":"控制中心配置","id":23,"options":"显示,隐藏","template":1,"tool_tip":"可配置是否显示通用一级菜单","whether_enable":1,"relation_keys":""}]', 'SystemCustom', '系统功能定制配置清单', '2020-07-20 01:29:32', '2020-07-20 01:29:32');
INSERT INTO `dcmc_sys_config` VALUES (7, 'SysPcConfig', '{\"iptables\":{\"content\":[\"iptables -A INPUT -p tcp --dport 80 -j DROP\",\"iptables -A OUTPUT -p tcp --sport 80 -j DROP\"]},\"screensaver\":{\"enable\":true,\"delay\":0},\"power\":{\"display_standby\":0},\"pam_cracklib\":{\"minlen\":8,\"reject_username\":false,\"minclass\":0},\"wallpaper\":{\"image\":\"/api/data/default_img/wallpaper.jpg\"},\"script\":{\"name\":\"\",\"content\":\"\",\"upload_at\":{}},\"can_login\":true}', 'SystemCustom', '系统区域基础配置', '2020-07-20 01:29:32', '2020-07-20 01:29:32');
INSERT INTO `dcmc_sys_config` VALUES (8, 'SysClientControlLimitNum', '1000', 'System', '系统终端控制最大数量限制', '2020-07-20 01:29:32', '2020-07-20 01:29:32');
INSERT INTO `dcmc_sys_config` VALUES (9, 'SysInitFixHtml', '<html>...</html>', 'System', '系统初始修复页面', '2020-07-20 01:29:32', '2020-07-20 01:29:32');
INSERT INTO `dcmc_sys_config` VALUES (10, 'SysInitFixScript', 'echo \"fix script ...\"', 'System', '系统初始修复脚本', '2020-07-20 01:29:32', '2020-07-20 01:29:32');
INSERT INTO `dcmc_sys_config` VALUES (11, 'SysSoftwareEtagTime', '', 'System', '系统软件上一次更新时间（初始化为空）', '2020-07-20 01:29:32', '2020-08-07 08:52:00');
INSERT INTO `dcmc_sys_config` VALUES (12, 'SysInitRecScript', 'gsettings set com.deepin.dde.daemon keybinding true ; qdbus org.kde.kglobalaccel /kglobalaccel blockGlobalShortcuts false ; gsettings set com.deepin.dde.dock.module.launcher enable true ; gsettings set com.deepin.dde.dock.module.menu enable true ; gsettings set com.deepin.dde.dock.module.dockapp enable true ; gsettings set com.deepin.dde.dock.module.network control false ; gsettings set com.deepin.dde.dock.module.disk-mount filemanager-integration true ; gsettings set com.deepin.dde.dock.module.trash enable true ; gsettings set com.deepin.dde.control-center hide-module \"[]\" ; gsettings set com.deepin.dde.Launcher filter-keys \"[]\"', 'System', '系统初始恢复脚本(解除管控)', '2020-07-20 01:29:32', '2020-07-20 01:29:32');
INSERT INTO `dcmc_sys_config` VALUES (13, 'SysAreaConfigKeys', 'pc_config setting_script area_app_list', 'System', '区域场景配置key', '2020-07-22 14:29:16', '2020-07-22 14:29:19');
INSERT INTO `dcmc_sys_config` VALUES (14, 'SysDdeDisEnableAppScriptTemplate', 'gsettings set com.deepin.dde.Launcher filter-keys \"[\'%s\']\"', 'System', 'dde应用禁用脚本模板', '2020-07-22 14:32:51', '2020-07-22 14:32:54');
INSERT INTO `dcmc_sys_config` VALUES (15, 'DeepinAppStoreBackgroundUrl', 'http://icbc-developer.chinauos.com', 'DeepinAppStore', '内网商店后台地址', '2020-07-29 20:57:11', '2020-08-10 05:58:54');
INSERT INTO `dcmc_sys_config` VALUES (16, 'usbconf', 1, 'System', '', '2020-07-29 20:57:11', '2020-08-10 05:58:54');




