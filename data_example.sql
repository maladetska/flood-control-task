DROP TABLE user_calls;
DROP TABLE flood;

CREATE TABLE user_calls
(
    user_id     BIGINT PRIMARY KEY,
    calls_count BIGINT
);

INSERT INTO user_calls (user_id, calls_count)
VALUES (1, 6);
INSERT INTO user_calls (user_id, calls_count)
VALUES (2, 7);
INSERT INTO user_calls (user_id, calls_count)
VALUES (3, 5);

CREATE TABLE flood
(
    user_id BIGINT PRIMARY KEY,
    end_at  TIMESTAMP
);

INSERT INTO flood (user_id, end_at)
VALUES (1, '2024-03-24T18:00:00');
INSERT INTO flood (user_id, end_at)
VALUES (2, '2024-03-24T18:06:00');
INSERT INTO flood (user_id, end_at)
VALUES (3, '2024-03-24T15:00:00');