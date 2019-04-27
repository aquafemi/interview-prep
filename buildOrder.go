package main

import (
	"fmt"
)

func main() {
	projects := []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	dependencies := [][]string{
		[]string{"f", "c"},
		[]string{"f", "b"},
		[]string{"f", "a"},
		[]string{"d", "g"},
		[]string{"c", "a"},
		[]string{"b", "a"},
		[]string{"a", "e"},
		[]string{"b", "e"},
		[]string{"b", "h"},
	}
	fmt.Println(getBuildOrder(projects, dependencies))

	dependencies = [][]string{
		[]string{"a", "b"},
		[]string{"b", "a"},
	}
	fmt.Println(getBuildOrder(projects[0:3], dependencies))
}

func getBuildOrder(projects []string, dependencies [][]string) []string {
	buildOrder := []string{}
	numDependencies := make(map[string]int)
	dependencyGraph := make(map[string][]string)
	
	addProjectsFromDependencies(dependencies, numDependencies, dependencyGraph)
	addRemainingProjects(projects, numDependencies, dependencyGraph)

	//fmt.Println(dependencyGraph)
	//fmt.Println(numDependencies)

	for len(dependencyGraph) != 0 {
		projectAdded := false
		for requiredProject, _ := range dependencyGraph {
			numDeps := numDependencies[requiredProject]
			if numDeps == 0 {
				projectAdded = true
				for _, targetProject := range dependencyGraph[requiredProject] {
					numDependencies[targetProject]--
				}
				buildOrder = addToBuildOrder(buildOrder, requiredProject, dependencyGraph)
			}

			//fmt.Println(requiredProject, numDependencies)
		}

		// We went through the whole dependency graph without finding a project with no dependencies.
		// Since every project is dependent on something, we have a cyclic dependency
		if projectAdded == false {
			return []string{}
		}
	}

	return buildOrder
}

func addProjectsFromDependencies(dependencies [][]string, numDependencies map[string]int, dependencyGraph map[string][]string) {
	for _, dep := range dependencies {
		requiredProject := dep[0]
		targetProject := dep[1]

		// Add the dependency to the target project
		_, exists := numDependencies[targetProject]
		if !exists {
			numDependencies[targetProject] = 0
		}
		numDependencies[targetProject]++

		// Add the projects that can be built after the required project
		_, exists = dependencyGraph[requiredProject]
		if !exists {
			dependencyGraph[requiredProject] = []string{}
		}
		dependencyGraph[requiredProject] = append(dependencyGraph[requiredProject], targetProject)
	}
}

func addRemainingProjects(projects []string, numDependencies map[string]int, dependencyGraph map[string][]string) {
	// Add remaining projects to numDependencies map and graph
	for _, project := range projects {
		// Get projects with no dependencies
		_, exists := numDependencies[project]
		if !exists {
			numDependencies[project] = 0
		}

		_, exists = dependencyGraph[project]
		if !exists {
			dependencyGraph[project] = []string{}
		}
	}
}

func addToBuildOrder(buildOrder []string, requiredProject string, dependencyGraph map[string][]string) []string {
	buildOrder = append(buildOrder, requiredProject)

	// Remove from the graph so we don't consider it anymore
	delete(dependencyGraph, requiredProject)
	return buildOrder
}
