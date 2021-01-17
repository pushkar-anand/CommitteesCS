CREATE TABLE IF NOT EXISTS `committee_members`
(
    `id`           INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `committee_id` INT UNSIGNED NOT NULL REFERENCES committees (id) ON DELETE CASCADE,
    `faculty_id`   INT UNSIGNED NOT NULL REFERENCES faculties (id) ON DELETE CASCADE,
    `created_at`   DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UNIQUE (committee_id, faculty_id)
);