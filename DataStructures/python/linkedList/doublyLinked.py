class Node:
    def __init__(self, data):
        self.data = data
        self.next = None
        self.prev = None

class List:
    def __init__(self):
        self.size = 0
        self.head = None
        self.tail = None

    def add(self, data):
        element = Node(data)

        if self.head is None:
            self.head = self.tail = element
        else:
            element.prev = self.tail
            self.tail.next = element
            self.tail = element

        self.size += 1

    def remove(self, data):
        if self.size == 0:
            return

        element = self.search(data)

        if element is None:
            return

        if element == self.head:
            if element == self.tail:
                self.tail = self.head = None
            else:
                self.head = element.next

            if self.head != None:
                self.head.prev = None

        else:
            if element == self.tail:
                self.tail = element.prev

            element.prev.next = element.next

        self.size -= 1

    def search(self, data):
        if self.size == 0:
            return None

        current = self.head
        while current != None:
            if current.data == data:
                break

            current = current.next
        
        return current

    def __str__(self):
        elements = self.getElementsData()

        return ', '.join(elements)

    def getElementsData(self):
        nodes = []

        cur = self.head
        while cur != None:
            nodes.append(cur.data)
            cur = cur.next

        return nodes
