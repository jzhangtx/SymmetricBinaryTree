// SymmetricBinaryTree.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>

class Node
{
public:
    Node* left;
    Node* right;
    int data;

    Node(int d)
        : data(d)
        , left(nullptr)
        , right(nullptr)
    {}
};

bool Symmetric(Node* n1, Node* n2)
{
    if (n1 == nullptr && n2 == nullptr)
        return true;
    if (n1 == nullptr || n2 == nullptr)
        return false;

    if (n1->data != n2->data)
        return false;
    return Symmetric(n1->left, n2->right) && Symmetric(n1->right, n2->left);
}

bool IsSymmetric(Node* root)
{
    return Symmetric(root->left, root->right);
}

int main()
{
    Node* root = new Node(1);
    root->left = new Node(2);
    root->right = new Node(2);
    root->left->left = new Node(4);
    root->left->right = new Node(5);
    root->right->left = new Node(5);
    root->right->right = new Node(4);

    std::cout << (IsSymmetric(root) ? "True" : "False") << std::endl;
}
