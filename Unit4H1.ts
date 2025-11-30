/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-30
 * @fileoverview This program stores 10 decimal numbers in an array and finds the largest.
 */

// variables
const numbers: number[] = new Array(10);
let largest: number;

// get 10 numbers from user
for (let i = 0; i < 10; i++) {
  numbers[i] = parseFloat(prompt(`Enter number ${i + 1}: `) || "0");
}

// find largest number
largest = numbers[0];
for (let i = 1; i < numbers.length; i++) {
  if (numbers[i] > largest) {
    largest = numbers[i];
  }
}

// display largest number
console.log(`\nThe largest value in your list is: ${largest}`);

// display all numbers
console.log("Here is the list of numbers: " + numbers.join(", "));

console.log("\nDone.");
