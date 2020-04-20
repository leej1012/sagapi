DROP TABLE IF EXISTS `tbl_qr_code`;
DROP TABLE IF EXISTS `tbl_api_key`;
DROP TABLE IF EXISTS `tbl_order`;
DROP TABLE IF EXISTS `tbl_error_code`;
DROP TABLE IF EXISTS `tbl_request_param`;
DROP TABLE IF EXISTS `tbl_api_detail_info`;
DROP TABLE IF EXISTS `tbl_api_basic_info`;

create table tbl_api_basic_info
(
 ApiId int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
 ApiLogo varchar(100) not null default '' COMMENT '',
 ApiName varchar(100) not null  default '' COMMENT '',
 ApiProvider varchar(100) not null default '' COMMENT '',
 ApiUrl varchar(100) not null  default '' COMMENT '',
 ApiPrice varchar(100) not null  default '' COMMENT '',
 ApiDesc varchar(100) not null  default '' COMMENT '',
 Specifications int(11) not null  default 0 COMMENT '规格',
 Popularity int(11) not null default 0 COMMENT '流行度',
 Delay int(11) not null default 0 COMMENT '',
 SuccessRate int(11) not null default 0 COMMENT '',
 InvokeFrequency int(11) not null default 0 COMMENT '',
 PRIMARY KEY (ApiId)
);


create table tbl_api_detail_info
(
 Id int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
 ApiId int(11) unique not null,
 Mark varchar(100) not null default '' COMMENT '',
 ResponseParam varchar(100) not null default ''  COMMENT '',
 ResponseExample varchar(100) not null default ''  COMMENT '',
 DataDesc varchar(100) not null default '' COMMENT '',
 DataSource varchar(100) not null default ''  COMMENT '',
 ApplicationScenario varchar(100) not null default '' COMMENT '',
 PRIMARY KEY (Id),
 foreign key(ApiId) references tbl_api_basic_info(ApiId)
);


create table tbl_request_param (
  Id int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  ApiDetailInfoId int(11) not null,
  ParamName varchar(50) not null default '',
  Required  tinyint(1) not null,
  ParamType varchar(10) not null default '',
  Note varchar(50) not null default '',
  PRIMARY KEY (Id),
  CONSTRAINT FK_request_param_id FOREIGN KEY (ApiDetailInfoId) REFERENCES tbl_api_detail_info(Id)
);


create table tbl_error_code (
  Id int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  ApiDetailInfoId int(11) not null,
  ErrorCode int(11) not null,
  ErrorDesc varchar(50) not null default '',
  PRIMARY KEY (Id),
  CONSTRAINT FK_error_code_id FOREIGN KEY (ApiDetailInfoId) REFERENCES tbl_api_detail_info(Id)
);


create table tbl_order (
  OrderId varchar(50) unique not null COMMENT '',
  ProductName varchar(50) not null default '' COMMENT '',
  OrderType varchar(50) not null default ''  COMMENT '',
  OrderTime int(11) not null default 0 COMMENT '下单时间',
  PayTime  int(11) not null default 0  COMMENT '支付时间',
  OrderStatus tinyint(1) not null default 0,
  Amount varchar(50) not null default '' COMMENT '',
  OntId varchar(50) not null default '' COMMENT '用户ontid',
  UserName varchar(50) not null default '' COMMENT '',
  TxHash varchar(50) not null  default '' COMMENT '支付交易hash',
  Price varchar(50) not null default ''  COMMENT '',
  ApiId int(11) NOT NULL COMMENT '',
  Specifications int(11) NOT NULL COMMENT '规格',
  PRIMARY KEY (OrderId),
  CONSTRAINT FK_tbl_order_id FOREIGN KEY (ApiId) REFERENCES tbl_api_basic_info(ApiId)
);


create table tbl_api_key (
  Id int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  ApiKey varchar(50) unique not null  default '',
  ApiId int(11) not null,
  OrderId varchar(50) unique not null COMMENT '',
  RequestLimit int(11) not null default 0,
  UsedNum int(11) not null default 0,
  OntId varchar(50) not null default '',
  PRIMARY KEY (Id),
  foreign key(OrderId) references tbl_order(OrderId)
);

CREATE TABLE `tbl_qr_code` (
  Id int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  QrCodeId varchar(50) unique not null default '',
  Ver varchar(50) not null default '',
  OrderId varchar(50) unique not null default '' ,
  Requester varchar(50) not null default '',
  Signature varchar(50) not null default '',
  Signer varchar(50) not null default '',
  QrCodeData varchar(50) not null default '',
  Callback varchar(50) not null default '',
  Exp varchar(50) not null default '',
  Chain varchar(50) not null default '',
  QrCodeDesc varchar(50) not null default '',
  PRIMARY KEY (Id),
  foreign key(OrderId) references tbl_order(OrderId)
);