import unittest
from singlyLinked import List as SinlgyLinkedList

class TestSinglyLinked(unittest.TestCase):
    list = SinlgyLinkedList()

    def test_0_add(self):
        self.list.add("first element")
        self.assertEqual(self.list.size, 1)
        self.assertEqual(self.list.tail, self.list.head)

        self.list.add("second element")
        self.assertEqual(self.list.size, 2)
        self.assertEqual(self.list.tail.data, "second element")

        self.list.add("third element")
        self.assertEqual(self.list.size, 3)
        self.assertEqual(self.list.tail.data, "third element")

    def test_1_search(self):
        element = self.list.search("non existent")
        self.assertEqual(element, None)

        element = self.list.search("second element")
        self.assertEqual(element.data, "second element")

    def test_2_remove(self):
        self.list.remove("non existent element")
        self.assertEqual(self.list.size, 3)

        self.list.remove("first element")
        self.assertEqual(self.list.size, 2)

        element = self.list.search("second element")
        self.assertEqual(self.list.head, element)
        self.assertNotEqual(self.list.tail, element)

        self.list.remove("third element")
        self.assertEqual(self.list.tail, element)
        self.assertEqual(self.list.tail.next, None)
        self.assertEqual(self.list.size, 1)

        self.list.remove("second element")
        self.assertEqual(self.list.head, self.list.tail)
        self.assertEqual(self.list.head, None)

if __name__ == '__main__':
    unittest.main()
