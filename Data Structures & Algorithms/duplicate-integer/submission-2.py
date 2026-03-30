class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
         
        numset = set()

        for number in nums:
            if number not in numset:
                numset.add(number)
            else:
                return True
        return False
        