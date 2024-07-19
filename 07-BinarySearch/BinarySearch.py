# binary search algorithim

# must declare function before using it 
def binarySearch(array, target):
  high = len(array) -1
  low = 0

  while low <= high:
    mid = (high+low) // 2 #use integer division
    if array[mid] < target:
      low = mid + 1
    elif array[mid] > target:
      high = mid - 1
    elif array[mid] == target:
      return mid
    
  return -1 # return -1 if not found 

nums = [1, 3, 5, 7, 8, 10]
index = binarySearch(nums, 10)
print(f"The number 10 is found at index: {index}")