ID: 46155c6dbed348cea6c0acf4fa8b4934
State: OPEN
User: nvc
Project: sit
Created: 2023-07-30T22:07:58+02:00
Tags: 

---

# Title
Move ticket struct and related code to a separate pkg 'ticket'

# Description
This will allow to concentrate all logic in one package.
It will improve naming:

```
ticket.New
ticket.State
ticket.Ticket
```
