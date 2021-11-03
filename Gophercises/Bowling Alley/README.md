# Bowling Scores

The goal of this program is to model a game of bowling. Given a series of input the program
should output the players final score.

# Specification/Rules of Bowling

- Each game, or line of bowling, includes ten turns, or frames for the bowler.
- In each frame, the bowler gets up to two tries to knock down all the pins.
- If in two tries, he fails to knock them all down, his score for that frame is the total number of
pins knocked down in his two tries.
- If in two tries he knocks them all down, this is called a spare and his score for the frame is
ten plus the number of pins knocked down on his next throw (in his next turn).
- If on his first try in the frame he knocks down all the pins, this is called a strike. His turn is
over, and his score for the frame is ten plus the simple total of the pins knocked down in his
next two rolls.
- If he gets a spare or strike in the last (tenth) frame, the bowler gets to throw one or two more
bonus balls, respectively. These bonus throws are taken as part of the same turn. If the
bonus throws knock down all the pins, the process does not repeat: the bonus throws are
only used to calculate the score of the final frame.
- The game score is the total of all frame scores.

> We will not check
> 1. Valid rolls
> 2. Correct number of rolls and frames.
> 3. We will not provide scores for intermediate frames.

# Example input and output

1. Gutter balls (all zero) 
> - [0,0],[0,0],[0,0],[0,0],[0,0],[0,0],[0,0],[0,0],[0,0],[0,0] 
> - 0 
 
2. All Threes 
> - [3,3],[3,3],[3,3],[3,3],[3,3],[3,3],[3,3],[3,3],[3,3],[3,3] 
> - 60 
 
3. All Spares with first ball a 4 
> - [4,6],[4,6],[4,6],[4,6],[4,6],[4,6],[4,6],[4,6],[4,6],[4,6,4] 
> - 140 
 
4. Nine Strikes followed by a gutter ball 
> - [10,0],[10,0],[10,0],[10,0],[10,0],[10,0],[10,0],[10,0],[10,0],[0,0] 
> - 240 
 
5. Perfect Game 
> - [10,0],[10,0],[10,0],[10,0],[10,0],[10,0],[10,0],[10,0],[10,0],[10,10,10]
> - 300