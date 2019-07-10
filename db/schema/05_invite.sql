
CREATE TABLE IF NOT EXISTS invite (
  invite_id     VARCHAR(32) PRIMARY KEY NOT NULL,
  invite_type   VARCHAR(16) NOT NULL,                            -- host | guest
  invite_status VARCHAR(16) NOT NULL DEFAULT 'pending',          -- pending | accepted | declined
  mateo_id      INTEGER     NOT NULL REFERENCES mateo(mateo_id),
  dinner_id     INTEGER     REFERENCES dinner(dinner_id),
  timestamp     TIMESTAMP   NOT NULL DEFAULT NOW(),
  UNIQUE (mateo_id, dinner_id)
);
