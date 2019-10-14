import numbers

class Node:
	def __init__(self, value):
		if not isinstance(value, numbers.Number):
			return -1

		self.value = value
		self.left = None
		self.right = None

class Tree:
	def __init__(self):
		self.root = None
		self.size = 0

	def insert(self, value):
		if not isinstance(value, numbers.Number):
			return -1

		if self.root is None:
			self.root = Node(value)
		else:
			self.insertNode(self.root, value)

	def insertNode(self, node, value):
		if node.value > value:
			if node.left is None:
				node.left = Node(value)
			else:
				return self.insertNode(node.left, value)
		else:
			if node.right is None:
				node.right = Node(value)
			else:
				return self.insertNode(node.right, value)

		self.size += 1

	def search(self, value):
		if not isinstance(value, numbers.Number):
			return -1

		return self.contains(self.root, value)

	def contains(self, node, value):
		if node is None:
			return -1

		if node.value == value:
			return node
		elif node.value > value:
			return self.contains(node.left, value)

		return self.contains(node.right, value)

	def remove(self, value):
		node = self.search(value)
