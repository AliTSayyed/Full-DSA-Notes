// binary search algorithim

import java.util.Arrays;

public class BinarySearch {
  public static void main(String[] args) {
    int[] array = { 1, 3, 5, 7, 8, 10 };
    int index = binarySearch(array, 10);
    System.out.println(Arrays.toString(array));
    System.out.println("The number 10 is found at index: " + index); // prints 5
  }

  public static int binarySearch(int[] nums, int target) {
    int high = nums.length - 1; // max array index
    int low = 0; // min array index

    while (low <= high) {
      int mid = (high + low) / 2; // find the average of the 2 pointers
      if (nums[mid] < target) {
        low = mid + 1;
      } else if (nums[mid] > target) {
        high = mid - 1;
      } else if (nums[mid] == target) {
        return mid;
      }
    }
    return -1; // return -1 if not found
  }
}