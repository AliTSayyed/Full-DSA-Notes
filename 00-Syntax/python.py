# instaitate a variable (dynamically typed)
x = 1
message = 'hello'
bool = True
double = 0.5

print('Hello World')
print('The value of x is', x) # type of x is int
print(f'The value of x is {x}') # formated print statement

# conditional statements
num = 10
if num > 5:
  print("num is greater than 5")
elif num < 5:
  print("num is less than 5")
else:  #num == 5
  print("num is equal to 5")

num2 = 10
num3 = 40
if num2 == 10 and num3 < 50: # and operator
  print("true")
elif num2 < 20 or num3 == 20: # or operator
  print('true')

# loops
i = 5
while i > 0:
  if i == 3:
    i -= 1 # must decrement before skipping printing 3 to avoid infinite loop
    continue
  if i == 1:
    break
  print(i)
  i -= 1

for j in range(5):
  print(j)

# math operations
add = 5 + 3
substract = 5 -3
mult = 5 * 3
divide1 = 5 / 3 # floating point division
divide2 = 5 // 3 # integer division, truncates decimal 
remainder = 5 % 3 
exponent = 5**3 # 5 ^ 3
decrement = 1
decrement -= 1 # no -- or ++ shortcut for a 1 digit decrement/increment
result = (5-3) / 2 * 5 # follows pemdas, does parenthesis first, then divides by 2, then multiplies by 5

# function
def greet(name: str): # def decalres function, greet is function name, and name is the paramter passed in. Must be declared above where it is used. :str is a type hint
  print("Hello" + name)
