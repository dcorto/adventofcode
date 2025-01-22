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

| **Day**                                                                              | **Solution**                                                                                                                          | **Runtime(A & B)**                                            | **Comments**                                                                                                                            |
|--------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------|
| [Day 1: Not Quite Lisp](https://adventofcode.com/2015/day/1)                         | <div style='display:block;white-space: nowrap;'>[Link](https://github.com/dcorto/adventofcode/blob/master/2015/1/main.go) ⭐️⭐️</div>  | <span style='white-space: nowrap;'>313.42µs & 219.29µs</span> |                                                                                                                                         |
| [Day 2: I Was Told There Would Be No Math](https://adventofcode.com/2015/day/2)      | <div style='display:block;white-space: nowrap;'>[Link](https://github.com/dcorto/adventofcode/blob/master/2015/2/main.go) ⭐️⭐️</div>  | <span style='white-space: nowrap;'>539.53µs & 466.85µs</span> |                                                                                                                                         |
| [Day 3: Perfectly Spherical Houses in a Vacuum](https://adventofcode.com/2015/day/3) | <div style='display:block;white-space: nowrap;'>[Link](https://github.com/dcorto/adventofcode/blob/master/2015/3/main.go) ⭐️⭐️</div>  | <span style='white-space: nowrap;'>1.64ms & 1.42ms</span>     |                                                                                                                                         |
| [Day 4: The Ideal Stocking Stuffer](https://adventofcode.com/2015/day/4)             | <div style='display:block;white-space: nowrap;'>[Link](https://github.com/dcorto/adventofcode/blob/master/2015/4/main.go) ⭐️⭐️</div>  | <span style='white-space: nowrap;'>765.43ms & 1.70s</span>    |                                                                                                                                         |
| [Day 5: Doesn't He Have Intern-Elves For This?](https://adventofcode.com/2015/day/5) | <div style='display:block;white-space: nowrap;'>[Link](https://github.com/dcorto/adventofcode/blob/master/2015/5/main.go) ⭐️⭐️</div>  | <span style='white-space: nowrap;'>1.60ms & 1.13ms</span>     |                                                                                                                                         |
| [Day 6: Probably a Fire Hazard](https://adventofcode.com/2015/day/6)                 | <div style='display:block;white-space: nowrap;'>[Link](https://github.com/dcorto/adventofcode/blob/master/2015/6/main.go) ⭐️⭐️</div>  | <span style='white-space: nowrap;'>73.30ms & 113.15ms</span>  |                                                                                                                                         |
| [Day 7: Some Assembly Required](https://adventofcode.com/2015/day/7)                 | <div style='display:block;white-space: nowrap;'>[Link](https://github.com/dcorto/adventofcode/blob/master/2015/7/main.go) ⭐️⭐️</div>  | <span style='white-space: nowrap;'>29.44ms & 31.63ms</span>   | Very similar of [Day 24 of Year 2024](../2024/24/main.go) but in this case de input was different, but the main idea is pretty the same |
| [Day 8: Matchsticks](https://adventofcode.com/2015/day/8)                            | <div style='display:block;white-space: nowrap;'>[Link](https://github.com/dcorto/adventofcode/blob/master/2015/8/main.go) ⭐️⭐️</div>  | <span style='white-space: nowrap;'>181.19µs & 1.075ms</span>  |                                                                                                                                         |
| [Day 9: All in a Single Night](https://adventofcode.com/2015/day/9)                  | <div style='display:block;white-space: nowrap;'>[Link](https://github.com/dcorto/adventofcode/blob/master/2015/9/main.go) ⭐️⭐️</div>  | <span style='white-space: nowrap;'>130.59ms & 113.04ms</span> | Solved using force brute, will not work with large inputs                                                                               |             
| [Day 11: Corporate Policy](https://adventofcode.com/2015/day/11)                     | <div style='display:block;white-space: nowrap;'>[Link](https://github.com/dcorto/adventofcode/blob/master/2015/11/main.go) ⭐️⭐️</div> | <span style='white-space: nowrap;'>13.93ms & 591.80ms</span>  |                                                                                                                                         |
| [Day 12: JSAbacusFramework.io](https://adventofcode.com/2015/day/12)                 | <div style='display:block;white-space: nowrap;'>[Link](https://github.com/dcorto/adventofcode/blob/master/2015/12/main.go) ⭐️⭐️</div> | <span style='white-space: nowrap;'>32.52ms & 4.58ms</span>    | Found "inspiration" in the [reddit megathread](https://www.reddit.com/r/adventofcode/comments/3wh73d/day_12_solutions/)                 |
| [Day 13: Knights of the Dinner Table](https://adventofcode.com/2015/day/13)          | <div style='display:block;white-space: nowrap;'>[Link](https://github.com/dcorto/adventofcode/blob/master/2015/13/main.go) ⭐️⭐️</div> | <span style='white-space: nowrap;'>331.15ms & 1.85s</span>    | Permutation function is the same as the [Day 9](./9/main.go)                                                                            |
| [Day 14: Reindeer Olympics](https://adventofcode.com/2015/day/14)                    | <div style='display:block;white-space: nowrap;'>[Link](https://github.com/dcorto/adventofcode/blob/master/2015/14/main.go) ⭐️⭐️</div> | <span style='white-space: nowrap;'>591.41µs & 780.44µs</span> |                                                                                                                                         |

## 📝 Notes

- Each solution is designed to work with the input provided by the Advent of Code website.
- Make sure to place the input file `input.txt` (if required) in the day folder.
- I have not included my own input sets in the repository as [recommended](https://www.reddit.com/r/adventofcode/comments/e7khy8/comment/fa13hb9/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button) by one of the mods in the official subreddit.

## 🔄 My Previous Years

[Advent of Code 2023](https://github.com/dcorto/adventofcode2023)

[Advent of Code 2022](https://github.com/dcorto/adventofcode2022)

[Advent of Code 2021](https://github.com/dcorto/adventofcode2021)

---

Happy coding and good luck with Advent of Code 2015! 🎉
