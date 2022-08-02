# Quiz Game
Quiz Game exercise from https://gophercises.com/

- Read in questions from CSV file (problems.csv)in format `question, answer` Present questions to user, prompt for answer, track how many they get right/wrong.
- allow user to pass in flag to use a custom csv input file
- flag to optionally shuffle quiz order
- add in a timer, default 30 seconds, customizable via flag, quiz stops immediatly when out of time

## v1
- quiz without the timer 

## v2
- refactored after watching solution video to incorporate some of instructor ideas
  - create a struct for the Q&A instead of leaving as a nested slice
  - trim space from the answer on file read

## v3
- implement timer after watching solution video