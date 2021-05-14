USE `chat`;

INSERT INTO `user` (`name`) VALUES ('じろう');
INSERT INTO `user` (`name`) VALUES ('たろう');

INSERT INTO `message` (`user_id`, `content`, `created_at`) VALUES 
  (2, "あ〜〜〜〜〜", FROM_UNIXTIME(1617371835)),
  (1, "う〜〜〜", FROM_UNIXTIME(1617372835)),
  (2, "画像が入る", FROM_UNIXTIME(1617373835));
