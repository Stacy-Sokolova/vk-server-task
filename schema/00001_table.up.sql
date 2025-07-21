CREATE TABLE users
(
    id          SERIAL PRIMARY KEY,
    login       varchar(255) NOT NULL,
    password    varchar(255) NOT NULL
);

CREATE TABLE ads
(
    id          SERIAL PRIMARY KEY,
    user_id     int NOT NULL,
    title       varchar(255) NOT NULL,
    description varchar(255) NOT NULL,
    image_url     varchar(255) NOT NULL,
    price       int not null NOT NULL,
    created_at  TIMESTAMP DEFAULT now()
);