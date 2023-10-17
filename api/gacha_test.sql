-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- 主机： localhost
-- 生成日期： 2023-07-14 13:44:42
-- 服务器版本： 8.0.24
-- PHP 版本： 7.2.33

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `gacha`
--

-- --------------------------------------------------------

--
-- 表的结构 `burn_record`
--

CREATE TABLE `burn_record` (
  `id` int NOT NULL,
  `user_id` int DEFAULT NULL COMMENT '用户id',
  `address` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `hash` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '交易hash',
  `block_height` bigint DEFAULT NULL,
  `tx_status` tinyint DEFAULT NULL,
  `amount` decimal(22,4) DEFAULT NULL,
  `status` tinyint DEFAULT NULL COMMENT '是否开奖 0-否 1-是',
  `create_time` int DEFAULT NULL COMMENT '创建时间',
  `update_time` int DEFAULT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='销毁记录';


--
-- 表的结构 `deposit_record`
--

CREATE TABLE `deposit_record` (
  `id` int NOT NULL,
  `create_time` int DEFAULT NULL COMMENT '创建时间',
  `lock_day` int DEFAULT '0' COMMENT '锁定天数',
  `deposit_amount` decimal(22,4) DEFAULT '0.0000' COMMENT '质押数量',
  `hash` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '交易hash',
  `block_height` bigint DEFAULT NULL,
  `tx_status` tinyint DEFAULT NULL,
  `confirm_num` int DEFAULT '0' COMMENT '确认数',
  `user_id` int DEFAULT NULL COMMENT '用户id',
  `address` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status` tinyint DEFAULT '1' COMMENT '状态 1-待确认 2-已确认 3-已失败',
  `update_time` int DEFAULT NULL COMMENT '更新时间',
  `ticket_amount` decimal(22,4) DEFAULT '0.0000' COMMENT 'ticket奖励数量'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='质押记录';


--
-- 表的结构 `prize`
--

CREATE TABLE `prize` (
  `id` int NOT NULL,
  `user_id` int DEFAULT NULL COMMENT '用户id',
  `prize_id` int DEFAULT NULL COMMENT '奖品id',
  `amount` decimal(22,4) DEFAULT NULL,
  `status` tinyint DEFAULT NULL COMMENT '领取状态 0-未领取 1-已领取',
  `update_time` int DEFAULT NULL COMMENT '更新时间',
  `create_time` int DEFAULT NULL COMMENT '创建时间',
  `prize_type` int DEFAULT NULL COMMENT '奖品类型',
  `receive_hash` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '领奖hash',
  `token_id` int DEFAULT NULL COMMENT 'tokenid'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='奖品';



--
-- 表的结构 `prize_record`
--

CREATE TABLE `prize_record` (
  `id` int NOT NULL,
  `user_id` int DEFAULT NULL COMMENT '用户id',
  `prize_id` int DEFAULT NULL COMMENT '奖品id',
  `amount` decimal(22,4) DEFAULT NULL,
  `status` tinyint DEFAULT NULL COMMENT '领取状态 0-未领取 1-已领取',
  `update_time` int DEFAULT NULL COMMENT '更新时间',
  `create_time` int DEFAULT NULL COMMENT '创建时间',
  `burn_hash` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '销毁ticket交易hash',
  `prize_type` int DEFAULT NULL COMMENT '奖品类型',
  `proof` text COLLATE utf8mb4_general_ci COMMENT '默克尔树证明',
  `receive_hash` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '领奖hash'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='奖品记录';

--
-- 表的结构 `user`
--

CREATE TABLE `user` (
  `id` bigint NOT NULL COMMENT '主键',
  `address` varchar(191) NOT NULL COMMENT '用户钱包地址',
  `name` varchar(20) DEFAULT NULL COMMENT '用户名称',
  `avatar_uri` varchar(512) NOT NULL COMMENT '用户图像地址',
  `create_time` bigint DEFAULT NULL COMMENT '创建时间',
  `update_time` bigint NOT NULL COMMENT '更新时间',
  `lottery_num` int DEFAULT NULL COMMENT '抽奖次数',
  `version` bigint DEFAULT '0' COMMENT '乐观锁'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- 表的结构 `user_asset`
--

CREATE TABLE `user_asset` (
  `id` int NOT NULL,
  `user_id` int DEFAULT NULL COMMENT '用户id',
  `balance` decimal(22,4) DEFAULT '0.0000' COMMENT '余额',
  `type` int DEFAULT '0' COMMENT '资产类型 1-ticket',
  `create_time` int DEFAULT NULL COMMENT '创建时间',
  `update_time` int DEFAULT NULL COMMENT '更新时间',
  `freeze` decimal(22,4) DEFAULT '0.0000' COMMENT '冻结'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户资产';

--
-- 表的结构 `withdraw_record`
--

CREATE TABLE `withdraw_record` (
  `id` int NOT NULL,
  `user_id` int DEFAULT NULL COMMENT '用户id',
  `address` varchar(64) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `amount` decimal(22,4) DEFAULT '0.0000' COMMENT '提现总额',
  `hash` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '交易hash',
  `block_height` bigint DEFAULT NULL,
  `tx_status` tinyint DEFAULT NULL,
  `status` tinyint DEFAULT NULL COMMENT '状态 1-待确认 2-已确认 3-已失败',
  `create_time` int DEFAULT NULL COMMENT '创建时间',
  `update_time` int DEFAULT NULL COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='提现记录';

--
-- 转储表的索引
--

--
-- 表的索引 `burn_record`
--
ALTER TABLE `burn_record`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `deposit_record`
--
ALTER TABLE `deposit_record`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `prize`
--
ALTER TABLE `prize`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `prize_record`
--
ALTER TABLE `prize_record`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `user_asset`
--
ALTER TABLE `user_asset`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `withdraw_record`
--
ALTER TABLE `withdraw_record`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `burn_record`
--
ALTER TABLE `burn_record`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=82;

--
-- 使用表AUTO_INCREMENT `deposit_record`
--
ALTER TABLE `deposit_record`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=40;

--
-- 使用表AUTO_INCREMENT `prize`
--
ALTER TABLE `prize`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=95;

--
-- 使用表AUTO_INCREMENT `prize_record`
--
ALTER TABLE `prize_record`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=466;

--
-- 使用表AUTO_INCREMENT `user`
--
ALTER TABLE `user`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键', AUTO_INCREMENT=15;

--
-- 使用表AUTO_INCREMENT `user_asset`
--
ALTER TABLE `user_asset`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- 使用表AUTO_INCREMENT `withdraw_record`
--
ALTER TABLE `withdraw_record`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
