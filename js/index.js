const _ = require('lodash');


let animals = ["gopher", "otter", "mole", "snake"]
let isCool = a => a.endsWith("er")

_.filter(animals, isCool) // [ 'gopher', 'otter' ]

