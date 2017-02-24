Run:

```
go build cmd/main.go
./main tests/case-2-input.txt
```

This will generate two files `case-2-input.txt-P1-status` `case-2-input.txt-P2-status` for each Player
you can tweek the `case-2-input.txt` and the similarly names files are generated for the Players.

`case-1-input.txt` holds the sample input provided in the questionb.

But the sample has some issues.

 1. (0 < s <= M/2) but while M = 5, s cannot be > 2. But in the example it is given as 5
 2. If we have to increase the M = 10, that is not possible. a (0 < M < 10) as per the problem statement.
 3. On the similar not the co-ordinates for the ship has to be tweeked.

 Please reach out at 9895865899 if you have questions.

 Make your test files and run as

 ```
 ./main tests/your-file
 ```


.