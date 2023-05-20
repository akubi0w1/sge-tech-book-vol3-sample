CREATE TABLE `character` (
  `id` int unsigned NOT NULL,
  `name` varchar(255) NOT NULL COMMENT '名前',
  `gender` int NOT NULL COMMENT '性別',
  `age` int NOT NULL COMMENT '年齢',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'キャラクター';

CREATE TABLE `card` (
  `id` int unsigned NOT NULL,
  `name` varchar(255) NOT NULL COMMENT 'カード名',
  `rarity_type` int NOT NULL COMMENT 'レアリティ',
  `character_id` int unsigned NOT NULL COMMENT 'キャラクターID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'カード';