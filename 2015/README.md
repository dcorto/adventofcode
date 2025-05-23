# 🎄 Advent of Code | int y=2015;

This repository contains my solutions for the [Advent of Code 2015](https://adventofcode.com/2015) challenges, implemented in **Go (Golang)**.

## 🚀 Usage

```bash
    # Add your input.txt file in the folder of the day you want to run

    # run specific day (x = number of day)
    $ make run-day day=x
    
    # run all days
    $ make run-all
    
    # Also you run directly using go:
    $ go run <day>/main.go
```

## ⭐ Solutions

| **Day**                                                                              | **Solution**       | **Stars** | **Runtime (A & B)**   | **Comments**                                                                                                                             |
|--------------------------------------------------------------------------------------|--------------------|-----------|-----------------------|------------------------------------------------------------------------------------------------------------------------------------------|
| [Day 1: Not Quite Lisp](https://adventofcode.com/2015/day/1)                         | [Link](1/main.go)  | ⭐️⭐️      | `313.42µs & 219.29µs` |                                                                                                                                          |
| [Day 2: I Was Told There Would Be No Math](https://adventofcode.com/2015/day/2)      | [Link](2/main.go)  | ⭐️⭐️      | `539.53µs & 466.85µs` |                                                                                                                                          |
| [Day 3: Perfectly Spherical Houses in a Vacuum](https://adventofcode.com/2015/day/3) | [Link](3/main.go)  | ⭐️⭐️      | `1.64ms & 1.42ms`     |                                                                                                                                          |
| [Day 4: The Ideal Stocking Stuffer](https://adventofcode.com/2015/day/4)             | [Link](4/main.go)  | ⭐️⭐️      | `765.43ms & 1.70s`    |                                                                                                                                          |
| [Day 5: Doesn't He Have Intern-Elves For This?](https://adventofcode.com/2015/day/5) | [Link](5/main.go)  | ⭐️⭐️      | `1.60ms & 1.13ms`     |                                                                                                                                          |
| [Day 6: Probably a Fire Hazard](https://adventofcode.com/2015/day/6)                 | [Link](6/main.go)  | ⭐️⭐️      | `73.30ms & 113.15ms`  |                                                                                                                                          |
| [Day 7: Some Assembly Required](https://adventofcode.com/2015/day/7)                 | [Link](7/main.go)  | ⭐️⭐️      | `29.44ms & 31.63ms`   | Very similar of [Day 24 of Year 2024](../2024/24/main.go) but in this case the input was different so the main idea is pretty the same   |
| [Day 8: Matchsticks](https://adventofcode.com/2015/day/8)                            | [Link](8/main.go)  | ⭐️⭐️      | `181.19µs & 1.075ms`  |                                                                                                                                          |
| [Day 9: All in a Single Night](https://adventofcode.com/2015/day/9)                  | [Link](9/main.go)  | ⭐️⭐️      | `130.59ms & 113.04ms` | Solved using force brute, will not work with large inputs                                                                                |             
| [Day 11: Corporate Policy](https://adventofcode.com/2015/day/11)                     | [Link](11/main.go) | ⭐️⭐️      | `13.93ms & 591.80ms`  |                                                                                                                                          |
| [Day 12: JSAbacusFramework.io](https://adventofcode.com/2015/day/12)                 | [Link](12/main.go) | ⭐️⭐️      | `32.52ms & 4.58ms`    | Found "inspiration" in the [reddit megathread](https://www.reddit.com/r/adventofcode/comments/3wh73d/day_12_solutions/)                  |
| [Day 13: Knights of the Dinner Table](https://adventofcode.com/2015/day/13)          | [Link](13/main.go) | ⭐️⭐️      | `331.15ms & 1.85s`    | Permutation function is the same as the [Day 9](./9/main.go)                                                                             |
| [Day 14: Reindeer Olympics](https://adventofcode.com/2015/day/14)                    | [Link](14/main.go) | ⭐️⭐️      | `591.41µs & 780.44µs` |                                                                                                                                          |
| [Day 15: Science for Hungry People](https://adventofcode.com/2015/day/15)            | [Link](15/main.go) | ⭐️⭐️      | `19.65ms & 13.96ms`   |                                                                                                                                          |
| [Day 16: Aunt Sue](https://adventofcode.com/2015/day/16)                             | [Link](16/main.go) | ⭐️⭐️      | `2.01ms & 2.22ms`     | You have 500 Aunts named "Sue". xd                                                                                                       |
| [Day 17: No Such Thing as Too Much](https://adventofcode.com/2015/day/17)            | [Link](17/main.go) | ⭐️⭐️      | `154.61ms & 146.78ms` |                                                                                                                                          |
| [Day 18: Like a GIF For Your Yard](https://adventofcode.com/2015/day/18)             | [Link](18/main.go) | ⭐️⭐️      | `116.90ms & 148.12ms` |                                                                                                                                          |
| [Day 19: Medicine for Rudolph](https://adventofcode.com/2015/day/19)                 | [Link](19/main.go) | ⭐️⭐️      | `1.28ms & 6.75ms`     | Took me a while the Part 2, finally using a shuffle replacements mapping i was able to found a valid solution in a right amount of time  |
| [Day 20: Infinite Elves and Infinite Houses](https://adventofcode.com/2015/day/20)   | [Link](20/main.go) | ⭐️️⭐️️    | `10.47s & 12.23s`     | Found "inspiration" in the [reddit megathread](https://www.reddit.com/r/adventofcode/comments/3xjpp2/day_20_solutions/)                  |
| [Day 21: RPG Simulator 20XX](https://adventofcode.com/2015/day/21)                   | [Link](21/main.go) | ⭐️️⭐️️    | `6.95ms & 7.81ms`     |                                                                                                                                          |
| [Day 22: Wizard Simulator 20XX](https://adventofcode.com/2015/day/22)                | [Link](22/main.go) | ⭐️️⭐️️    | `425.76ms & 8.31ms`   |                                                                                                                                          |
| [Day 23: Opening the Turing Lock](https://adventofcode.com/2015/day/23)              | [Link](23/main.go) | ⭐️️⭐️️    | `401.07µs & 462.47µs` |                                                                                                                                          |
| [Day 24: It Hangs in the Balance](https://adventofcode.com/2015/day/24)              | [Link](24/main.go) | ⭐️️⭐️️    | `2.72s & 729.72ms`    |                                                                                                                                          |
| [Day 25: Let It Snow](https://adventofcode.com/2015/day/25)                          | [Link](25/main.go) | ️⭐️️⭐️    | `250.43ms`            |                                                                                                                                          |

## 📝 Notes

- Each solution is designed to work with the input provided by the Advent of Code website.
- Make sure to place the input file `input.txt` (if required) in the day folder.
- I have not included my own input sets in the repository as [recommended](https://www.reddit.com/r/adventofcode/comments/e7khy8/comment/fa13hb9/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button) by one of the mods in the official subreddit.

## 🔄 My Other Years

[Advent of Code 2024](../2024)

[Advent of Code 2023](https://github.com/dcorto/adventofcode2023)

[Advent of Code 2022](https://github.com/dcorto/adventofcode2022)

[Advent of Code 2021](https://github.com/dcorto/adventofcode2021)

---

Happy coding and good luck with Advent of Code 2015! 🎉
