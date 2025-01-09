package main

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	directFriends, exists := friendships[username]
	if !exists {
		return nil
	}

	directFriendsSet := make(map[string]struct{}, len(directFriends))
	for _, friend := range directFriends {
		directFriendsSet[friend] = struct{}{}
	}

	suggestedFriendsSet := make(map[string]struct{})

	for _, friend := range directFriends {
		for _, potentialFriend := range friendships[friend] {
			if potentialFriend == username {
				continue
			}
			if _, isDirectFriend := directFriendsSet[potentialFriend]; isDirectFriend {
				continue
			}
			suggestedFriendsSet[potentialFriend] = struct{}{}
		}
	}

	result := make([]string, 0, len(suggestedFriendsSet))
	for friend := range suggestedFriendsSet {
		result = append(result, friend)
	}

	if len(result) == 0 {
		return nil
	}

	return result
}


func main() {
	runTestCases()
}
// algorithm
// ------------------------------------
// 1. Initialize an empty map or set to track suggested friends.
//    This ensures we avoid duplicate suggested friends.

// 2. Retrieve the slice of direct friends for the given username from the friendships map.
//    If the username is not present in the map, return nil immediately as there are no friends to process.

// 3. Create another map or set to store all direct friends of the username. 
//    This will help in quickly checking if a potential suggestion is already a direct friend.

// 4. Loop through the slice of the user's direct friends:
//    a. For each direct friend, retrieve their direct friends from the friendships map.
//    b. For each friend of the direct friend, check if:
//       i. They are not the username themselves (to avoid suggesting the user as their own friend).
//       ii. They are not already a direct friend of the username (use the map/set created in Step 3).
//    c. If both conditions are satisfied, add them to the suggested friends map or set.

// 5. Convert the suggested friends map or set to a slice. 
//    This ensures the function returns the results in the required format.

// 6. Return the resulting slice of suggested friends.
//    If no suggestions were added, return nil.
