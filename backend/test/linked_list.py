# linked_list.py
from node import Node  # Import the Node class from node.py

class LinkedList:
    def __init__(self):
        self.head = None  # Initialize the head of the linked list to None
    
    def append(self, data):
        new_node = Node(data)  # Create a new node with the given data
        if not self.head:  # If the list is empty, make the new node the head
            self.head = new_node
            return
        last_node = self.head
        while last_node.next:  # Traverse the list to find the last node
            last_node = last_node.next
        last_node.next = new_node  # Set the next of the last node to the new node
    
    def display(self):
        current_node = self.head
        while current_node:
            print(current_node.data, end=" -> ")
            current_node = current_node.next
        print("None")  # Indicate the end of the list
