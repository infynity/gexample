

DROP TABLE IF EXISTS `dcmc_area_setting`;
CREATE TABLE `dcmc_area_setting` (
                                     `id` int(11) NOT NULL AUTO_INCREMENT,
                                     `area_id` int(11) NOT NULL,
                                     `sys_software_config` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                     `sys_setting_config` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                     `area_name` varchar(30) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                     `scene_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                     `scene_code` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
                                     `scene_sort` int(11) DEFAULT NULL,
                                     `limit_policy` tinyint(3) DEFAULT NULL,
                                     `limit_software` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                     `pc_config` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                     `repair_zone_iptables` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                     `full_zone_switch` tinyint(3) DEFAULT NULL,
                                     `full_zone_verify_script` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                     `full_zone_fix_html` longtext COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                     `full_zone_fix_url` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                     `app_script` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                     `setting_script` text COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                     `created_at` datetime DEFAULT NULL,
                                     `updated_at` datetime DEFAULT NULL,
                                     `full_zone_verify_script_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                     `full_zone_fix_html_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                                     PRIMARY KEY (`id`),
                                     KEY `area_id` (`area_id`),
                                     KEY `scene_code` (`scene_code`),
                                     KEY `scene_sort` (`scene_sort`)
) ENGINE=InnoDB AUTO_INCREMENT=93 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


DROP TABLE IF EXISTS `dcmc_web_message_to`;
CREATE TABLE `dcmc_web_message_to` (
                                       `id` int(11) NOT NULL AUTO_INCREMENT,
                                       `created_at` timestamp NULL DEFAULT NULL,
                                       `updated_at` timestamp NULL DEFAULT NULL,
                                       `deleted_at` timestamp NULL DEFAULT NULL,
                                       `web_message_id` int(11) DEFAULT NULL,
                                       `to` int(11) DEFAULT NULL,
                                       `status` varchar(255) DEFAULT NULL,
                                       `receiver_id` int(11) DEFAULT NULL,
                                       `receiver_deleted_at` timestamp NULL DEFAULT NULL,
                                       PRIMARY KEY (`id`),
                                       KEY `idx_dcmc_web_message_to_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8;