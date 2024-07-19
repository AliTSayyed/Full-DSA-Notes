from Person import Person # from the person file import the person object 

class Main: 
  @staticmethod # make the main method a static method so no main object needs to be created 
  def main(): 
    person1 = Person("Bob", 24)
    person1.greet()

if __name__ == "__main__": # call the main function only when this line has been reached
  Main.main()