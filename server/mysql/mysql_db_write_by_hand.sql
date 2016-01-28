SET @ORIG_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `table_authorization`;
DROP TABLE IF EXISTS `table_app`;
DROP TABLE IF EXISTS `table_mac`;
DROP TABLE IF EXISTS `table_coordinate_record`;
DROP TABLE IF EXISTS `table_running_record`;
DROP TABLE IF EXISTS `table_users`;
DROP view IF EXISTS `view_mac_app`;
DROP view IF EXISTS `view_mac_coordinate`;
DROP view IF EXISTS `view_current_mac`
SET FOREIGN_KEY_CHECKS=@ORIG_FOREIGN_KEY_CHECKS;

CREATE TABLE `table_authorization` (
  `Authorization` int(3) unsigned zerofill NOT NULL DEFAULT '010' COMMENT '权限（范围001-999，不含小数）',
  `Authorization_Info` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '权限说明（长度限制255）',
  `Authorization_Password` int(8) unsigned zerofill NOT NULL DEFAULT '00000000' COMMENT '权限密码（不加密，长度限制8）',
  PRIMARY KEY (`Authorization`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=FIXED COMMENT='权限表';

CREATE TABLE `table_app` (
  `App_Name` varchar(30) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '应用名称（长度限制30）',
  `App_Path` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '应用路径（长度限制255）',
  `App_Authorization` int(3) unsigned zerofill NOT NULL DEFAULT '999' COMMENT '应用权限（默认为999）',
  `App_Info` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '应用信息（长度限制255）',
  PRIMARY KEY (`App_Name`),
  KEY `App_Authorization_Cnstraint` (`App_Authorization`),
  CONSTRAINT `App_Authorization_Cnstraint` FOREIGN KEY (`App_Authorization`) REFERENCES `table_authorization` (`Authorization`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=FIXED COMMENT='应用表';

CREATE TABLE `table_mac` (
  `MAC_address` varchar(12) COLLATE utf8_unicode_ci NOT NULL COMMENT 'MAC地址（12位字符）',
  `MAC_name` varchar(50) CHARACTER SET utf8 DEFAULT NULL COMMENT '设备名称或说明（长度限制50）',
  `MAC_Authorization` int(3) unsigned zerofill NOT NULL DEFAULT '010' COMMENT '设备权限（默认为010）',
  `Date_Add` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '添加日期（自动生成，请勿指定）',
  PRIMARY KEY (`MAC_address`),
  KEY `MAC_Authorization_Constraint` (`MAC_Authorization`),
  CONSTRAINT `MAC_Authorization_Constraint` FOREIGN KEY (`MAC_Authorization`) REFERENCES `table_authorization` (`Authorization`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=FIXED COMMENT='设备表';

CREATE TABLE `table_coordinate_record` (
  `Coordinate_Num` int(11) unsigned zerofill NOT NULL AUTO_INCREMENT COMMENT '坐标记录流水号（自动生成，勿指定）\r\n\r\n',
  `Coordinate_Longitude` double(11,8) unsigned zerofill NOT NULL DEFAULT '00.00000000' COMMENT '坐标经度（长度11\r\n\r\n，小数点后8）',
  `Coordinate_Latitude` double(10,8) unsigned zerofill NOT NULL DEFAULT '0.00000000' COMMENT '坐标纬度（长度10，小数点后8）',
  `MAC_address` varchar(12) COLLATE utf8_unicode_ci NOT NULL COMMENT 'MAC地址（12位字符，需检查范围）',
  `Date_Add` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '添加日期（自动生成，请勿指定）',
  PRIMARY KEY (`Coordinate_Num`),
  KEY `MAC_Coordinate_Constraint` (`MAC_address`),
  CONSTRAINT `MAC_Coordinate_Constraint` FOREIGN KEY (`MAC_address`) REFERENCES `table_mac` (`MAC_address`)
) ENGINE=InnoDB AUTO_INCREMENT=249 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=FIXED COMMENT='坐标记录表';

CREATE TABLE `table_running_record` (
  `Num` int(11) unsigned zerofill NOT NULL AUTO_INCREMENT COMMENT '流水号',
  `Date_Add` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '添加日期（自动生成，请勿指定）',
  `MAC_address` varchar(12) COLLATE utf8_unicode_ci NOT NULL COMMENT 'MAC地址（12位字符，需检查范围）',
  `Data` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '操作内容（长度限制255）',
  PRIMARY KEY (`Num`),
  KEY `Mac_Running_Constraint` (`MAC_address`),
  CONSTRAINT `Mac_Running_Constraint` FOREIGN KEY (`MAC_address`) REFERENCES `table_mac` (`MAC_address`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=FIXED COMMENT='流水';

CREATE TABLE `table_users` (
  `User_Name` varchar(50) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '登陆用户名（长度限制50）',
  `User_Password` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '管理员登陆密码（md5加密后长度限制255）',
  `User_Email` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '邮件（长度限制255）',
  PRIMARY KEY (`User_Name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci ROW_FORMAT=FIXED COMMENT='管理员信息表';

create view view_mac_app as
select MAC_Address,Authorization_Password,App_Name,App_Path
from table_authorization inner join table_mac inner join table_app
where Authorization=MAC_Authorization and Authorization>=App_Authorization;

create view view_mac_coordinate as
select MAC_Address,Date_Add,Coordinate_Longitude,Coordinate_Latitude
from table_coordinate_record 
where date_add(Date_Add,interval 0 day)>date_add(NOW(),interval -1 hour);

create view view_current_mac as
select MAC_Address,max(Date_Add) as Date_Add
from view_mac_coordinate
group by MAC_Address;
