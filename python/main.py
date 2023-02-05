import random

animals = ["gopher", "otter", "mole", "snake"]
is_cool = lambda a: "er" in a

numbers = [0, 100, 50, -10, -365.42]
is_pos = lambda n: n > 0

# Using built-in filter
print(list(filter(is_cool, animals)))
print(list(filter(is_pos, numbers)))

# Using list comprehension
print([a for a in animals if is_cool(a)])
print([n for n in numbers if is_pos(n)])

# map cool animals to random shuffling of their names
print(
    {a: "".join(random.sample(list(a), len(a)))
     for a in animals
     if is_cool(a)}
)
