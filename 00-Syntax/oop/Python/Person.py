class Person: # create person class
  def __init__(self, name, age): # object constructor, uses __init__ and self 
    self.name = name
    self.age = age
  
  def greet(self): # needs self param to be a method instead of a function 
    print(f"Hello my name is {self.name} and I am {self.age} years old!")
