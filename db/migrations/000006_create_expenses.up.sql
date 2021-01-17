CREATE TABLE `expenses`
(
    `id`         INT      NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `event_id`   INT      NOT NULL REFERENCES `events` (id) ON DELETE CASCADE,
    `item`       TEXT     NOT NULL,
    `cost`       DOUBLE   NOT NULL,
    `bill`       BLOB     NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);