INSERT INTO users (name, nick, email, password)
VALUES
    ('user 1', 'user 1', 'user1@gmail.com', '$2a$10$K0nmGpx5mKKmbrjuYrXgHuiMttcUsjDCUwm7EgaHfzw7TACONYMBG'),
    ('user 2', 'user 2', 'user2@gmail.com', '$2a$10$K0nmGpx5mKKmbrjuYrXgHuiMttcUsjDCUwm7EgaHfzw7TACONYMBG'),
    ('user 3', 'user 3', 'user3@gmail.com', '$2a$10$K0nmGpx5mKKmbrjuYrXgHuiMttcUsjDCUwm7EgaHfzw7TACONYMBG');

INSERT INTO followers (user_id, follower_id)
VALUES
    (1, 2),
    (3, 1),
    (1, 3);