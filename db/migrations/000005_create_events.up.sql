CREATE TABLE IF NOT EXISTS `events`
(
    `id`                INT      NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name`              TEXT     NOT NULL,
    `start_date`        DATETIME NOT NULL,
    `end_date`          DATETIME NOT NULL,
    `total_expenditure` DOUBLE   NOT NULL,
    `created_at`        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);