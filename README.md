# meat-night [![Build Status](https://travis-ci.com/wexel-nath/meat-night.svg?branch=master)](https://travis-ci.com/wexel-nath/meat-night)
Tracking Mateo's Meat Night hosting and attendance.

## Scheduled Messages
Messages are scheduled to run on a cron set up through [Temporize](https://devcenter.heroku.com/articles/temporize)

To manually set up a cron:

```
curl -X 'POST' 'https://user:pass@api.temporize.net/v1/events/{cron}/{callback}'
```
