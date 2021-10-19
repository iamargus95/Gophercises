# Problem Statement

There is a sequence of words in CamelCase as a string of letters, , having the following properties:

- It is a concatenation of one or more words consisting of English letters.
- All letters in the first word are lowercase.
- For each of the subsequent words, the first letter is uppercase and rest of the letters are lowercase.

Given a string, determine the number of words in string.

## Example

There are  words in the string: 'one', 'Two', 'Three'.

#### Function Description

camelcase has the following parameter(s):

- string s: the string to analyze

Returns

- int: the number of words in input string

Input Format

- A single line.

Constraints

- Sample Input

    - saveChangesInTheEditor

- Sample Output

    - 5

- Explanation

    - String  contains five words:

        - save
        - Changes
        - In
        - The
        - Editor