
CREATE TABLE IF NOT EXISTS attendee(
  id            SERIAL  PRIMARY KEY NOT NULL,
  attendee_id   INTEGER NOT NULL REFERENCES mateo(id),
  meat_night_id INTEGER NOT NULL REFERENCES meat_night(id)
)
