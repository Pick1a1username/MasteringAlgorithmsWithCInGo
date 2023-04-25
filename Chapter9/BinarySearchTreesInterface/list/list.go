package list

type MatchFunc func(key1, key2 interface{}) int

type DestroyFunc func(data interface{})

type ListElmt struct {
	Data interface{}
	Next *ListElmt
}

type List struct {
	Size    int
	Match   MatchFunc
	Destroy DestroyFunc
	Head    *ListElmt
	Tail    *ListElmt
}

func ListInsNext(list *List, element *ListElmt, data interface{}) int {
	// Allocate storage for the element.
	newElement := &ListElmt{}
	// Insert the element into the list
	newElement.Data = data
	if element == nil {
		// Handle insertion at the head of the list.
		if list.Size == 0 {
			list.Tail = newElement
		}
		newElement.Next = list.Head
		list.Head = newElement
	} else {
		// Handle insertion somewhere other than at the head.
		if element.Next == nil {
			list.Tail = newElement
		}
		newElement.Next = element.Next
		element.Next = newElement
	}
	// Adjust the size of the list to account for the inserted element.
	list.Size++

	return 0
}
