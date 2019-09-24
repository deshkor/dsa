import unittest
from unbalanced import Tree as UnbalancedTree

class TestUnbalanced(unittest.TestCase):
	tree = UnbalancedTree()

	def test_0_add(self):
		self.assertEqual(self.tree.insert("shouldn't work"), -1)

		self.assertEqual(self.tree.insert(5), None)
		self.assertEqual(self.tree.insert(3), None)
		self.assertEqual(self.tree.insert(8), None)
		self.assertEqual(self.tree.insert(8), None)
		self.assertEqual(self.tree.insert(7), None)
		self.assertEqual(self.tree.insert(4), None)
		self.assertEqual(self.tree.insert(10), None)
		self.assertEqual(self.tree.root.right.value, 8)
		self.assertEqual(self.tree.root.right.right.value, 8)
		self.assertEqual(self.tree.root.right.left.value, 7)
		self.assertEqual(self.tree.root.left.value, 3)

	def test_1_search(self):
		self.assertEqual(self.tree.search("nope"), -1)
		self.assertEqual(self.tree.search(3).value, 3)
		self.assertEqual(self.tree.search(3).right.value, 4)
		self.assertEqual(self.tree.search(7).value, 7)
		self.assertEqual(self.tree.search(8).left.value, 7)
		self.assertEqual(self.tree.search(8).right.right.value, 10)

if __name__ == '__main__':
	unittest.main()

