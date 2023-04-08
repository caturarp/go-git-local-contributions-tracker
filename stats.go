package main

func stats(email string) {
	commits := processRepositories(email)
	printStats(commits)
}

func processRepositories(email string) map[int]int {
	filepath := getDotFilePath()
	repos := parseFileLinesToSlice(filepath)
	daysInMap := daysInLastQuarter

	commits := make(map[int]int, daysInMap)
	for i := daysInMap; i > 0; i-- {
		commits[i] = 0
	}
	for _, path := range repos {
		commits := listCommits(email, path, commits)
	}
	return commits
}
