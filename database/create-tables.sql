DROP TABLE IF EXISTS todo;
CREATE TABLE todo (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  description     VARCHAR(255) NOT NULL,
  priority      DECIMAL(5,2) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO todo
  (title, description, priority)
VALUES
  ('Clean the room', 'To clean my room: my table, my bed, the bathroom and the balcony', 1),
  ('Buy foods for the next week', 'Pork, Beef, Salmon, Potato, Vegies', 2),
  ('Join birthday party', 'My nephew is almost 1 year old', 1);