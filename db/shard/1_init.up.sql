CREATE TABLE `user` (
  `id` varchar(36) NOT NULL,
  `name` varchar(255) NOT NULL COMMENT '名前',
  `platform` int NOT NULL COMMENT '利用プラットフォーム',
  `created_at` datetime NOT NULL COMMENT '作成日時',
  `updated_at` datetime NOT NULL COMMENT '更新日時',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'ユーザ';

CREATE TABLE `user_card` (
  `id` varchar(36) NOT NULL,
  `user_id` varchar(36) NOT NULL COMMENT '所持ユーザ',
  `card_id` int unsigned NOT NULL COMMENT 'カード',
  `count` int unsigned NOT NULL COMMENT '所持数',
  `created_at` datetime NOT NULL COMMENT '作成日時',
  `updated_at` datetime NOT NULL COMMENT '更新日時',
  PRIMARY KEY (`id`),
  UNIQUE (`user_id`, `card_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '所持カード';
