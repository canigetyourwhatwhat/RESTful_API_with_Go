package platform

import "testing"

func TestAdd(t *testing.T){
	repo := New()

	repo.add(Item{"it is title", "it is detail", false})

	if len(repo.Items) != 1{
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T){
	repo := New()

	repo.add(Item{"wash dishes", "do it in 15 min", false})
	repo.add(Item{"wash laundry", "do it in 15 min", false})
	repo.add(Item{"wash a dog", "do it in 15 min", false})

	if len(repo.getAll()) != 3{
		t.Errorf("Not all items are displayed")
	}
}

