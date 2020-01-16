import java.util.LinkedList;

class Solution {
    public String simplifyPath(String path) {
        LinkedList<String> stack = new LinkedList<>();
        for (String dir : path.split("/")) {
            if ("..".equals(dir) && !stack.isEmpty()) stack.pop();
            else if (!"".equals(dir) && !".".equals(dir) && !"..".equals(dir)) stack.push(dir);
        }
        String res = "";
        for (String dir : stack) {
            res = "/" + dir + res;
        }
        return res.isEmpty() ? "/" : res;
    }
}