const _ = require('lodash');


let animals = ["gopher", "otter", "mole", "snake"]
let isCool = a => a.endsWith("er")

let numbers = [0, 100, 50, -10, -365.42]
let isPos = n => n > 0

console.log(_.filter(animals, isCool)) // [ 'gopher', 'otter' ]
console.log(_.filter(numbers, isPos))

