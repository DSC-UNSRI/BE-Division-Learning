
DROP INDEX `idx_answers_question_id` ON `answers`;
DROP INDEX `idx_answers_user_id` ON `answers`;
DROP INDEX `idx_questions_user_id` ON `questions`;

DROP TABLE IF EXISTS `votes`;
DROP TABLE IF EXISTS `answers`;
DROP TABLE IF EXISTS `questions`;
DROP TABLE IF EXISTS `users`;