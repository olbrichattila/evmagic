CREATE TABLE event_replay (
    event_id varchar(64) NOT NULL,
    event_payload MEDIUMTEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    KEY idx_event_replay_id (event_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
