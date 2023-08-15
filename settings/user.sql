DROP TABLE IF EXISTS users;
CREATE TABLE users (
                       user_id BIGINT PRIMARY KEY,
                       follow_count INT,
                       follower_count INT,
                       work_count INT,
                       favorite_count INT,
                       username VARCHAR(255),
                       password VARCHAR(255),
                       avatar VARCHAR(255),
                       background_image VARCHAR(255),
                       signature TEXT,
                       total_favorited VARCHAR(255),
                       is_follow BOOLEAN
);
