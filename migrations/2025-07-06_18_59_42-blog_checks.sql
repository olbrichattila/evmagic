CREATE TABLE blog_checks (
    blog_id INT UNSIGNED NOT NULL,
    check_type VARCHAR(64),
    reason VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    KEY idx_blog_check_blog_id (blog_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
