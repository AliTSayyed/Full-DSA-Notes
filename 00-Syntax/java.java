public class java {
  // main method constructor that will execute the program
  public static void main(String[] args) {

    // instanatiate a variable
    int x = 1;
    String message = "hello";
    boolean bool = true;
    double decimal = 0.5;

    // print statement
    System.out.println("Hello, World!");
    System.out.println("The values are " + x + bool + decimal + message);

    // conditional statements
    int num = 10;
    if (num > 5) {
      System.out.println("num is greater than 5");
    } else if (num < 5) {
      System.out.println("num is less than 5");
    } else { // num == 5
      System.out.println("num is equal to 5");
    }

    int num2 = 10;
    int num3 = 40;
    if (num2 == 10 && num3 < 50) { // and operator
      System.out.print("true");
    } else if (num2 < 20 || num3 == 20) { // or operator
      System.out.println("true");
    }

    // loops
    int i = 5;
    while (i > 0) {
      if (i == 3) {
        i--; // must decrement before skipping printing 3 to avoid infinite loop
        continue;
      }
      if (i == 1) {
        break;
      }
      System.out.println(i);
      i--;
    }

    for (int j = 0; j < 5; j++) {
      System.out.println(j);
    }

    // math operations
    int add = 5 + 3;
    int substract = 5 - 3;
    int mult = 5 * 3;
    int divide1 = 5 / 3; // integer division, truncates decimal
    double divide2 = 5.0 / 3.0; // floating point division
    int remainder = 5 % 3;
    double exponent = Math.pow(5, 3); // 5 ^ 3
    double result = (5 - 3) / 2 * 5; // follows pemdas, does parenthesis first, then divides by 2, then multiplies by
                                     // 5
    System.out.println("The values are: " + add + substract + mult + divide1 + divide2 + remainder + exponent + result);
    
    greet("bob");
  }

  // class functions
  public static void greet(String name) { // public acess, static (do not need to create an object), void (return type),
                                          // String name (paramater type paramter name)
    System.out.println("Hello " + name);
  }
}
