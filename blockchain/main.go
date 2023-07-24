package main

func main() {
	bc := NewBlockChain()
	bc.AddBlock(5, nil)
	bc.AddBlock(7, nil)
	bc.addTransaction(NewTransaction("B", "A", 1.1))
	bc.addTransaction(NewTransaction("C", "A", 2.2))
	bc.addTransaction(NewTransaction("D", "A", 3.3))
	bc.addTransaction(NewTransaction("E", "A", 4.4))
	bc.addTransaction(NewTransaction("F", "A", 5.5))
	bc.Print()
}
