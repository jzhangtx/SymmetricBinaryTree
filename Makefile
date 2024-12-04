SymmetricBinaryTree: SymmetricBinaryTree.o
	g++ -g -o SymmetricBinaryTree.exe SymmetricBinaryTree.o -pthread    

SymmetricBinaryTree.o: SymmetricBinaryTree/SymmetricBinaryTree.cpp
	g++ -g  -c -pthread SymmetricBinaryTree/SymmetricBinaryTree.cpp
