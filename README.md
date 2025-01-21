# journee

## Overview
1. Journal
2. Reflect
3. Connect / Visualize

Submit your daily journal entry -> Reflect on goals, screentime, etc. express gratitude -> See your friends emotion/goals/status and send encouragement

## Features
- [x] Track emotions and energy levels
- [~] Set and reflect on goals, screentime, etc
  - "I would like to [spend more time/spend less time/build a habit/break a habit/etc]"
- [ ] Visualize your journey on the dashboard
- [ ] Share one reflection, excerpt, achievement, etc each day

## Parking lot
- Markdown support
- Quotes API
- Daily checkins via Discord integration?
- LLM integration for memories?

## TODO
- [x] More date input margin-bottom
- [ ] Encrypt journal content with server-side key
- [D] Impl cursor to go to next/prev date with content
- [ ] put slider/emotion selection in div with vh height
- [ ] Impl basic client-side LRU cache for journal data
- [A] Impl logout
- [ ] Impl signup

- [ ] Impl forgot password?
- [ ] .air.toml

## Before Production
- `make gen` to generate static files from templates
- `docker build --platform linux/amd64 -t journee .`
