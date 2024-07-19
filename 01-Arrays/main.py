# Python has Lists (no arrays). Lists are dynamic (no fixed size), can hold multiple types, and are passed as refrences. They can be sorted. Duplicate values are ok. 

# create a list by surrounding data types with [] and seperating values with ","
fruits = ["apple", "orange", "banana", "coconut"]
print(fruits)
# this is how to access an element in a List
print(fruits[0])
# you can choose what index to start the list from
print(fruits[1:])
# can create a step that prints every second element
print(fruits[::2])
# can reverse the order of the list by using a -1 step
print(fruits[::-1])

# use len function to get the length of a list
print(len(fruits))

# can iterate over a list using a for loop
for x in range(len(fruits)):
  print(fruits[x])

# can use a for each loop as well
for fruit in fruits:
  print(fruit)

# if you forget what methods you can use on a list, print the dir() function
# print(dir(fruits))

# if you forget what each method does, you can print the help()
# print(help(fruits))

# you can check if a value is contained in a list
contain = "apple" in fruits
print(contain)

# can change the element at an index
fruits[0] = "grape"

# can add items to a list
fruits.append("pineapple")
print(fruits)

# can remove an item at a certain index or the item itself if it is in the list 
fruits.remove(fruits[1])
fruits.remove("pineapple")
print(fruits)

# can retrieve the index of a element
print(fruits.index("grape"))

# can insert an item at a specific index
fruits.insert(1, "apple")
print(fruits)

# can sort the list
fruits.sort()
print(fruits)

# can reverse a list 
fruits.reverse()
print(fruits)

# can save a sorted list to a variable
sortedFruits = sorted(fruits)
sortedFruits[3] = "bob"
print(sortedFruits)
print(fruits) # sorted fruits is a seperate list from fruits not pointing to the same refrence. Changes to sortedFruits does not affect fruits

# lists are refrences, changing the value of a copy will affect the original list
number = [1,2,3,4]
number2 = number
number2[0] = 10 # even if we modify number2 it will affect number1 since they point to the same refrence in memory.
print(number) 

#-----------------------------------------------------------------------------------------------------------------------------#

# tupes are ordered lists but are unchangeable. They can store multiple types. Duplicate values are allowed. Faster than a list. 

# declare a tuple by using () instead of []
friends = ('Bob', 'Sam', 'Tim')
print(friends)

# can check if a tuple contains an element
print("Tom" in friends)

# can keep track of how many of the same items are in a list using the count method
print(friends.count("Bob"))

# can iterate over a tuple
for friend in friends:
  print(friend)