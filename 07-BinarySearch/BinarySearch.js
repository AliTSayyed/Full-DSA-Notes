// binary search algorithim
// run in terminal using node BinarySearch.js

const nums = [1, 3, 5, 7, 8, 10];
const index = binarySearch(nums, 10);
console.log("The number 10 is found at index: " + index);


function binarySearch(array, target) {
  let high = array.length - 1;
  let low = 0;

  while (low <= high) {
    let mid = Math.floor((high + low) / 2); // js does not truncate, need to use math.floor
    if (array[mid] < target) {
      low = mid + 1;
    } else if (array[mid] > target) {
      high = mid - 1;
    } else if (array[mid] == target) {
      return mid;
    }
  }
  return -1; // return -1 if not found
}