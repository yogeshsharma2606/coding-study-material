# Greedy — Stock Problems

**Greedy** means making the locally optimal choice at each step and trusting it leads to the global optimum. It only works when you can back it with an **exchange argument** ("no alternative choice could do better"). When greedy fails, you fall back to DP. Stock problems are the cleanest illustration of *when a constraint makes greedy valid*.

**The framing that ties them together:** the constraint on *how many transactions you may make* completely changes the strategy:
- **One transaction** → find the single best buy-low/sell-high pair → track the running minimum.
- **Unlimited transactions** → capture *every* upswing → sum all positive daily deltas.

## Programs

### [Best time to buy and sell once](best-time-buy-sell-once/best_time_buy_sell_once.go)

Maximize profit from **one** buy and one later sell.

**How to think about it:** you must buy **before** you sell, so as you scan left→right, the best sell price today is `today − (cheapest day seen so far)`. So carry `minPrice` (the best buy point up to now) and, at each day, either update the min or update the best profit. One pass, no need to try all pairs (that would be O(n²)).

**Dry run** — `[7,1,5,3,6,4]`:

```
minPrice=7
1 < 7 -> minPrice=1
5: profit 5-1=4 -> best=4
3: profit 3-1=2
6: profit 6-1=5 -> best=5   <- max
4: profit 4-1=3
```

Answer **5** (buy at 1, sell at 6).

- **O(n) time, O(1) space.**
- **Why greedy is valid:** the optimal sell day pairs with the minimum price *to its left*; scanning maintains exactly that.

### [Best time to buy and sell repeatedly](best-time-buy-sell-multiple/best_time_buy_sell_multiple.go)

Maximize profit with **unlimited** transactions (buy/sell as often as you like, one share at a time).

**The insight (why summing daily gains is optimal):** any profitable multi-day rise `p[j] − p[i]` telescopes into the sum of consecutive daily gains between `i` and `j`. So you can equivalently "buy at the close of every up-day and sell the next day". Just **add every positive `p[i] − p[i-1]`** and ignore the down days. You never lose by capturing each individual upswing.

**Dry run** — `[7,1,5,3,6,4]`: deltas `-6,+4,-2,+3,-2` → sum positives `4+3 = 7`. Answer **7**.

- **O(n) time, O(1) space.**
- **Contrast with the once version:** here the constraint is relaxed, so a *simpler* greedy (sum upswings) becomes optimal. Recognizing how the constraint changes the strategy is the whole lesson.
- **Where greedy stops working:** add a transaction *limit* (at most k trades) or a *cooldown/fee*, and greedy breaks — that's the DP family ("Best Time to Buy and Sell Stock III/IV").

## Review checklist

- State the exchange argument that justifies each greedy strategy.
- Why does "sum all positive daily deltas" equal the true max profit with unlimited trades?
- How does the number-of-transactions constraint flip the approach?
- At what added constraint does greedy fail and DP take over?
