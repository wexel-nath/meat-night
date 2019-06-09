
-- CREATE TABLE IF NOT EXISTS attendee(
--   id            SERIAL  PRIMARY KEY NOT NULL,
--   attendee_id   INTEGER NOT NULL REFERENCES mateo(id),
--   meat_night_id INTEGER NOT NULL REFERENCES meat_night(id)
-- )

-- INSERT INTO mateo (first_name, last_name)
-- VALUES
--   ('Alex', 'Smith'),
--   ('Callum', 'Brown'),
--   ('James', 'Stewart'),
--   ('Nathan', 'Welch'),
--   ('Rav', 'Prasad'),
--   ('Tristan', 'Barr');

INSERT INTO attendee (attendee_id, meat_night_id)
VALUES
  (1, 1),
  (2, 1),
  (5, 1);
-- TODO: insert the rest from google drive
