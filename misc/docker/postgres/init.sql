CREATE TABLE example_users (
  id INT NOT NULL,
  username VARCHAR(32) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE example_messages (
  id INT NOT NULL,
  user_id INT NOT NULL,
  text VARCHAR(128) NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO example_users
VALUES
  (1, 'ondrejsika'),
  (2, 'alice'),
  (3, 'bob')
;

INSERT INTO example_messages
VALUES
  (1, 1, 'Hello World!'),
  (2, 1, 'Ahoj Svete!'),
  (3, 1, 'Hallo Welt!')
;
