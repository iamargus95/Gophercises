# Caesar Cipher

**What is a Caesar Cipher?**
- Julius Caesar protected his confidential information by encrypting it using a cipher.
- Caesar's cipher shifts each letter by a number of letters.
- If the shift takes you past the end of the alphabet, just rotate back to the front of the alphabet.
- In the case of a rotation by 3, w, x, y and z would map to z, a, b and c.

Example:

> Original alphabet:-        abcdefghijklmnopqrstuvwxyz

> Alphabet rotated +3:-      defghijklmnopqrstuvwxyzabc

## Usage
`go run caesarCipher.go <input-string> <int-delta-value>`

> **NOTE**: All non-alphabetic charachters remain the same. 