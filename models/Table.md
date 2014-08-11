|——table1:
|————intelligencedata
|————————id
|————————domain
|————————path
|————————flag
|————————discreption
|
|——table2:
|————users
|————————id
|————————username
|————————passwd
|————————permission
|————————discreption
|
|——table2:
|————permission
|————————pid
|————————pname
|————————discreption





CREATE TABLE `intelligencedata` (
`gid` INT(10) NOT NULL AUTO_INCREMENT,
`domain` VARCHAR(64) NOT NULL,
`path` VARCHAR(64) NOT NULL,
`discreption` VARCHAR(255) NULL DEFAULT NULL,
`flag` TINYINT NOT NULL DEFAULT 1,
PRIMARY KEY (`gid`)
)ENGINE=MyISAM DEFAULT CHARSET=utf8 

CREATE TABLE `users` (
`uid` INT(10) NOT NULL AUTO_INCREMENT,
`username` VARCHAR(64) NOT NULL,
`passwd` VARCHAR(64) NOT NULL,
`discreption` VARCHAR(64) NOT NULL,
`pid` INT(10) NOT NULL,
PRIMARY KEY (`uid`)
)ENGINE=MyISAM DEFAULT CHARSET=utf8 

CREATE TABLE `permission` (
`pid` INT(10) NOT NULL AUTO_INCREMENT,
`discreption` VARCHAR(64) NOT NULL,
PRIMARY KEY (`pid`)
)ENGINE=MyISAM DEFAULT CHARSET=utf8 