CREATE TABLE `users` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL UNIQUE,
    `password` VARCHAR(255) NOT NULL, -- Akan menyimpan hash bcrypt
    `tier` ENUM('free', 'premium') NOT NULL DEFAULT 'free',
    `security_question` TEXT NOT NULL,
    `security_answer` VARCHAR(255) NOT NULL, -- Akan menyimpan hash bcrypt
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` DATETIME DEFAULT NULL
);

CREATE TABLE `questions` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `title` VARCHAR(255) NOT NULL,
    `body` TEXT NOT NULL,
    `user_id` INT NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
);

CREATE TABLE `answers` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `body` TEXT NOT NULL,
    `user_id` INT NOT NULL,
    `question_id` INT NOT NULL,
    `upvotes` INT NOT NULL DEFAULT 0,
    `downvotes` INT NOT NULL DEFAULT 0,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`question_id`) REFERENCES `questions`(`id`) ON DELETE CASCADE
);

CREATE TABLE `votes` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `answer_id` INT NOT NULL,
    `vote_type` TINYINT NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    UNIQUE KEY `unique_user_answer_vote` (`user_id`, `answer_id`),

    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE,
    FOREIGN KEY (`answer_id`) REFERENCES `answers`(`id`) ON DELETE CASCADE
);

CREATE INDEX `idx_questions_user_id` ON `questions`(`user_id`);
CREATE INDEX `idx_answers_user_id` ON `answers`(`user_id`);
CREATE INDEX `idx_answers_question_id` ON `answers`(`question_id`);