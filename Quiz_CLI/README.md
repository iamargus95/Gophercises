# Exercise #1: Quiz Game

## Exercise details

### Part 1

- [x] Create a program that will read in a quiz provided via a CSV file
- [x] Ask quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.

- [x] The CSV file should default to `problems.csv` (example shown below), but the user should be able to customize the filename via a flag.

The CSV file will be in a format like below, where the first column is a question and the second column in the same row is the answer to that question.
```
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

- [x] At the end of the quiz the program should output the total number of questions correct and how many questions there were in total. Questions given invalid answers are considered incorrect.

**NOTE:** *CSV files may have questions with commas in them. Eg: `"what is 2+2, sir?",4` is a valid row in a CSV. Look into the CSV package in Go*

### Part 2

- [x] Adapt program from part 1 to add a timer. The default time limit should be 30 seconds, but should also be customizable via a flag.

- [x] The quiz should stop as soon as the time limit has exceeded.

- [x] Users should be asked to press enter (or some other key) before the timer starts, and then the questions should be printed out to the screen one at a time until the user provides an answer. Regardless of whether the answer is correct or wrong the next question should be asked.

- [x] At the end of the quiz the program should still output the total number of questions correct and how many questions there were in total. Questions given invalid answers or unanswered are considered incorrect.