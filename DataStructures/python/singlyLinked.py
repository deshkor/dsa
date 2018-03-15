class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

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
            self.tail.next = element
            self.tail = element

        self.size += 1

    def remove(self, data):
        if self.size == 0:
            return

        prev, element = self.search(data, True)

        if element is None:
            return        

        if element == self.head:
            self.head = element.next

            if element == self.tail:
                self.tail = self.head

        elif prev != None:
            prev.next = element.next

            if element == self.tail:
                self.tail = prev

        else:
            return

        self.size -= 1

    def search(self, data, returnPrev=False):
        if self.size == 0:
            if returnPrev:
                return (None, None)

            return None

        previous = None
        current = self.head
        while current != None:
            if current.data == data:
                break

            previous = current
            current = previous.next

        if returnPrev:
            return (previous, current)

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
