
class Solution {
    public int[] plusOne(int[] digits) {
        int k = digits.length;
        for (int i=k-1;i>=0;i--) {
            digits[i]++;
            digits[i] %= 10;
            if (digits[i]!=0) return digits;
        }
        digits= new int[k+1];
        digits[0]=1;
        return digits;
    }
}