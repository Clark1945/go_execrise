class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
linked_data = ListNode()
data_1 = ListNode(1)
linked_data.next = data_1 
data_2 = ListNode(2)
linked_data.next = data_1 
data_3 = ListNode(3)
linked_data.next = data_1 

print(data_1)