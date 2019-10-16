import java.util.Stack;

// Design a stack that supports push,pop,getMin. 
// Time and space complexity of all operations must be O(1).

class Solution {
    private Stack<Integer> stack = new Stack<>();
    private int min;

    public void push(int x) {
        if (stack.isEmpty()) {
            stack.push(x);
            min = x;
        } else {
            min = x - min < 0 ? x : min;
            stack.push(x - min);
        }
    }

    public void pop(int x) {
        int top = stack.peek();
        min = top < 0 ? min - top : min;
        stack.pop();
    }

    public int getMin() {
        return min;
    }
}