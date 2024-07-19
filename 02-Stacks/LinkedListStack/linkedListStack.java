// create an empty stack using a linked list implementation 
package LinkedListStack;

public class linkedListStack<Item> {
  private Node first = null;
  private int size;

  private class Node {
    Item item;
    Node next;
  }

  public boolean isEmpty() {
    return first == null;
  }

  void push(Item item) {
    // store the head
    Node oldFirst = first;
    // create a new head
    first = new Node();
    // set value of new Node to the item
    first.item = item;
    // point the new head to the old head.
    first.next = oldFirst;
    // keep track of the size
    size++;
  }

  public Item pop() {
    if (isEmpty()) {
      return null;
    }
    // remove the first item on the list
    // retrieve the item at the first Node (Head)
    Item item = first.item;
    // delete the node by setting it equal to its next node.
    first = first.next;
    // decrement the size
    size--;
    return item;
  }

  public int size() {
    return size;
  }

  public static void main(String[] args) {
    linkedListStack<String> words = new linkedListStack<>();
    words.push("Food");
    words.push("Water");
    System.out.println(words.pop());
  }

}