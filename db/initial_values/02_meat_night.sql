
-- CREATE TABLE IF NOT EXISTS meat_night (
--   id      SERIAL      PRIMARY KEY NOT NULL,
--   date    DATE        NOT NULL DEFAULT NOW(),
--   host_id INTEGER     NOT NULL REFERENCES mateo(id),
--   venue   VARCHAR(32) NOT NULL
-- )

-- INSERT INTO mateo (first_name, last_name)
-- VALUES
--   ('Alex', 'Smith'),
--   ('Callum', 'Brown'),
--   ('James', 'Stewart'),
--   ('Nathan', 'Welch'),
--   ('Rav', 'Prasad'),
--   ('Tristan', 'Barr');

INSERT INTO meat_night (date, host_id, venue)
VALUES
  ('10/18/18', 5, 'Sendokgarpu'),
  ('15/11/18', 4, 'Taco Bell');
-- TODO: insert the rest from google drive
