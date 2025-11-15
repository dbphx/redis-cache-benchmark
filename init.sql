CREATE TABLE IF NOT EXISTS users (
  id INT PRIMARY KEY,
  name VARCHAR(50)
);

INSERT INTO users (id, name)
SELECT seq, CONCAT('user_', seq)
FROM (SELECT @row := @row + 1 AS seq FROM mysql.help_topic, (SELECT @row := 0) init LIMIT 100000) t;
