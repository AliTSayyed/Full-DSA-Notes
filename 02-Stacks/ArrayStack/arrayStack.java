// create a stack using an array implementation 
package ArrayStack;

public class arrayStack<Item> {

  private Item[] stack; // declare an array of type Item
  private int N = 0;

  public arrayStack(int capacity) {
    stack = (Item[]) new Object[capacity]; // initilize array (must use a cast for a generic initilization)
  }

  public boolean isEmpty() {
    return N == 0;
  }

  public void push(Item item) {
    // if the array is 100% full, double the size of the array. 
    if (N == stack.length) {
      resize(2 * stack.length);
    }
    // stack[N] = item, then increment N
    stack[N++] = item;
  }

  public Item pop() {
    // Decrement N and store the item
    Item item = stack[--N];
    // Set the current array index to null to delete the item
    stack[N] = null;
    // if the array is less than 25% full, decrease the size of the array by half
    if (N > 0 && N == stack.length/4){
      resize(stack.length/2);
    }
    return item;
  }

  public void resize(int capacity) {
    Item[] copy = (Item[]) new Object[capacity]; // initilize a copy array of type Item, need to use a cast again
    for (int i = 0; i < N; i++) {
      copy[i] = stack[i];
    }
    stack = copy;
  }

  public static void main(String[] args){
    arrayStack<String> words = new arrayStack<>(1);
    words.push("Food");
    words.push("Water");
    System.out.println(words.pop());
  }
}
