import java.util.ArrayList;
import java.util.Arrays;

public class Main {
  public static void main(String[] args) {
    // An array is used to store multiple values in a single variable (must be of the same type).
    // Arrays are of fixed size. Can not increase the amount of space they have after declaring. 
    // Typically the memeory adresses of the actual values (elements) in the array themselves are stored in contiguous (adjacent) blocks of memory for O(1) efficiency when acessing the array. 
    // But the variable that stores the array is a refrence to the memeory adress of the first element in the array. The array variable is NOT a value itself.
    // This is the same for ArrayLits. They are refrences values. 

    int[] numbers = new int[5]; // how to declare and allocate memory for an array. Must declare the data type as well. 
    String[] cars = {"Truck", "SUV", "Cedan"}; // second way of creating an array. Create a list of values. 
    
    // this is how to assign elements to index values. Index always starts with 0.
    numbers[0] = 1;
    numbers[1] = 2;
    numbers[2] = 3;
    numbers[3] = 4;
    numbers[4] = 5;

    // this is how to loop through an array and acess each value
    for(int i = 0; i < numbers.length; i++) {
      System.out.println(numbers[i]);
    }

    // this is how to update a element at a certain index. Overwrites pervious element.
    cars[0] = "Electric Truck"; 

    // for each loop to iterate through each item in a list (another way)
    for (String car : cars) {
      System.out.println(car);
    }
    
    // this is how to get the size of an array. It is a property of the array. 
    System.out.println(numbers.length);

    // printing out an array gives the memeory adress since an array is a refrence. Need to use the toString method to actually print out the values. 
    System.out.println(cars);
    System.out.println(Arrays.toString(cars));

    // arrays are refrences so if you set another variable equal to an array it will still point to the same array and changes will affect the original array
    int[] moreNumbers = numbers;
    moreNumbers[4] = 10;
    // even though we modified moreNumbers, numbers is still changed. Need to manually copy each value of the numbers array into the more Numbers array to create an actual copy. 
    System.out.println(Arrays.toString(numbers)); 
    
    
    /* --------------------------------------------------------------------------------------------------------------------------------------------- */ 
    
    // a resizable array is called an ArrayList. This can have dynamic space. You can add and remove elements from the ArrayList unlike the array. 
    ArrayList<String> friends = new ArrayList<>();
    // this is how to add items to the ArrayList. Can also declare at the beginning using Arrays.asList method. 
    friends.add("Bob");
    friends.add("Sam");
    friends.add("Tim");
    // this is how to get an item from an ArrayList. 
    System.out.println(friends.get(2));

    // this is how to remove an element from the list. 
    friends.remove(2);

    // this is how to update an elment at a certein index in a list. 
    friends.set(0,"Boby");

    // can still iterate through the array list using a for loop or a for each loop (like below)
    for (String friend : friends){
      System.out.println(friend);
    }

    // this is how to acess the size of the array list. It is a method of the array list. 
    System.out.println(friends.size());

    // you can normally print out the array list, no need to use a to string method
    System.out.println(friends);

    /* --------------------------------------------------------------------------------------------------------------------------------------------- */ 

    // simple way of creating a resizable array
    int[] newNumbers = new int[5];
    for (int i = 0; i < 10; i++){
      // if there is not enough space in the original array, use the resize method to increase capacity. 
      if (i >= newNumbers.length){
        int newSize = newNumbers.length * 2;
        resizeArray(newNumbers,newSize);
      }
      newNumbers[i] = i;
    }
    System.out.println(Arrays.toString(newNumbers));
  
  }

  // this is the underlying way to create a resizable array (ArrayList)
  // resizing an array uses constant amortized time
  public static void resizeArray(int[] oldArray, int capacity) {
    int[] copy = new int[capacity];
    // simply declare a new array with a larger capacity. Then copy all the values of the old array into the new array with a for loop. 
    for (int i = 0; i < oldArray.length; i++){
      copy[i] = oldArray[i];
    }
    oldArray = copy;
  }
}