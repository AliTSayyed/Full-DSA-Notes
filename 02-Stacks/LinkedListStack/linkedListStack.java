// create an empty stack using a linked list implementation 
package LinkedListStack;

public class linkedListStack<Item> {
  private Node first = null;
  private int size;

  private class Node{
    Item item;
    Node next;
  }
  public boolean isEmpty(){
    return first == null;
  }

  void push(Item item){
    // create a node to store the previous head node
    Node oldFirst = first;
    // overwrite first node with a new null Node
    first = new Node();
    // set value of new Node to the item
    first.item = item; 
    // point the new Node to the stored Node. First will now become the new Head. 
    first.next = oldFirst;
    // increase the size by 1
    size++;
  }

  public Item pop() {
    // retrieve the item at the first Node (Head)
    Item item = first.item;
    // delete the node by setting it equal to its next node. 
    first = first.next;
    // decrement the size 
    size--;
    return item;
  }

 
  public int size(){
    return size;
  }

  public static void main(String[] args) {
    linkedListStack<String> words = new linkedListStack<>();
    words.push("Food");
    words.push("Water");
    System.out.println(words.pop());
  }

}