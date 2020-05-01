CREATE DATABASE `charts_db`;

USE `charts_db`;

CREATE TABLE `cats` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `category` VARCHAR(50),
    `count` INT
);
CREATE TABLE `dogs` (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `category` VARCHAR(50),
    `count` INT
);

INSERT INTO
    `cats` (`category`, `count`)
VALUES
    ('Cats', 45),
    ('Cats playing music instruments', 30),
    ('Awesome Code Examples', 10),
    ('Narwhal sightings', 5),
    ('Grahing libraries', 5),
    ('item 1', 5),
    ('item 2', 15),
    ('item 3', 50),
    ('Kittens', 5);

INSERT INTO
    `dogs` (`category`, `count`)
VALUES
    ('Dogs', 45),
    ('Puppies', 5);