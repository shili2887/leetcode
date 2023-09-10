package src;

class ListNode {
    int val;
    ListNode next = null;

    public ListNode(int val) {
        this.val = val;
    }
}

public class Q92_ReverseLinkedList {

    public ListNode reverseBetween(ListNode head, int m, int n) {
        // write code here
        if (head == null || head.next == null || m == n) {
            return head;
        }
        ListNode[] nodes = new ListNode[n - m + 1];
        ListNode cur = head;
        ListNode pre = null;
        int i = 1;
        while (i <= n) {
            if (i >= m && i <= n) {
                nodes[i - m] = cur;
            }
            if (i == m - 1) {
                pre = cur;
            }
            cur = cur.next;
            i++;
        }
        // reverse
        for (int j = nodes.length - 1; j > 0; j--) {
            nodes[j].next = nodes[j - 1];
        }
        // 接回去
        if (pre == null) {
            head = nodes[n - m];
        } else {
            pre.next = nodes[n - m];
        }
        nodes[0].next = cur;
        return head;
    }

}
