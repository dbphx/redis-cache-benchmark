CREATE TABLE IF NOT EXISTS users (
                                     id INT PRIMARY KEY,
                                     name VARCHAR(50)
    );

SET @row := 0;

INSERT INTO users (id, name)
SELECT @row := @row + 1 AS seq,
       CONCAT('user_', @row)
FROM mysql.help_topic
    LIMIT 100000;
