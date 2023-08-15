CREATE TABLE videos (
                        id INT PRIMARY KEY,
                        favorite_count INT,
                        comment_count INT,
                        play_url VARCHAR(255),
                        cover_url VARCHAR(255),
                        title VARCHAR(255),
                        is_favorite BOOLEAN,
                        author_id BIGINT,
                        FOREIGN KEY (author_id) REFERENCES users(id)
);
