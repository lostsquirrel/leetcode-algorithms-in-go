package main

func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]bool)
	walker := head
	for walker != nil {
		_, hasKey := m[walker]
		if hasKey {
			return true
		} else {
			m[walker] = true
			walker = walker.Next

		}
	}
	return false
}
