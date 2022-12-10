Given a JSON file with a list of people and their dates of birth, write a program to print out the people whose birthday is today.

If a person was born on Feb 29th, then their birthday during a non-leap year should be considered to be Feb 28th.

Input sample file:

```json
[
    ["Doe", "John", "1982/10/08"],
    ["Wayne", "Bruce", "1965/01/30"],
    ["Gaga", "Lady", "1986/03/28"],
    ["Curry", "Mark", "1988/02/29"]
]
```

If you make any assumptions, trade-offs or de-prioritise features for timeliness, please document these decisions.

Your submission must have:

- Instructions to run the code
- Tests to check if your code is correct, robust and complete.
- Edge cases handled and tested.
- Iterative development, backed by commit history.
- Modular, cohesive code with sensible separation of concerns.

Now we also want to be able to handle a CSV file with this additonal piece of data :arrow_double_down:

```
Doe,John,1982/10/08,john.doe@email.com,USA,
Wayne,Bruce,1965/01/30,bruce.wayne@email.com,Great Britain,
Gaga,Lady,1986/03/28,lady.gaga@email.com,Ireland,
Curry,Mark,1988/02/29,mark.curry@email.com,Spain,
```



Now that you have a working prototype, let us add analytics for the list of people.

:warning: *For this we will need a server that should be kinda RESTful* *!!!*

-  You will have to record the people thay have been sent a greeting. If already recorded the amount of emails sent should be recorded for data.
- For a given month, your manager wants to know how many people had their birthdays this month.
