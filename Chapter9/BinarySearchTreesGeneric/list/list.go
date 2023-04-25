package list

type MatchFunc[T any] func(key1, key2 *T) int

type DestroyFunc[T any] func(data *T)

type ListElmt[T any] struct {
	Data *T
	Next *ListElmt[T]
}

type List[T any] struct {
	Size    int
	Match   MatchFunc[T]
	Destroy DestroyFunc[T]
	Head    *ListElmt[T]
	Tail    *ListElmt[T]
}

func ListInsNext[T any](list *List[T], element *ListElmt[T], data *T) int {
	// Allocate storage for the element.
	newElement := &ListElmt[T]{}
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
