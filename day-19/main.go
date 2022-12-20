package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Minerals struct {
	Ore            int
	Clay           int
	Obsidian       int
	Geodes         int
	OreRobots      int
	ClayRobots     int
	ObsidianRobots int
	GeodeRobots    int
}

type Blueprint struct {
	OreRobotOreReq        int
	ClayRobotOreReq       int
	ObsidianRobotOreReq   int
	ObsidianRobotClayReq  int
	GeodeRobotOreReq      int
	GeodeRobotObsidianReq int
}

func main() {
	blueprints, err := readBlueprints("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partOne(blueprints))
	fmt.Println(partTwo(blueprints))
}

func partOne(blueprints []Blueprint) int {
	total := 0
	for index, blueprint := range blueprints {
		total += findMaxGeodes(blueprint, 24) * (index + 1)
	}
	return total
}

func partTwo(blueprints []Blueprint) int {
	return findMaxGeodes(blueprints[0], 32) * findMaxGeodes(blueprints[1], 32) * findMaxGeodes(blueprints[2], 32)
}

func findMaxGeodes(blueprint Blueprint, days int) int {
	minerals := []Minerals{{Ore: 0, Clay: 0, Obsidian: 0, Geodes: 0, OreRobots: 1, ClayRobots: 0, ObsidianRobots: 0, GeodeRobots: 0}}
	completeMinerals := []Minerals{}
	minMax := 0
	for i := 1; i <= days; i++ {
		daysRemaining := days - i + 1
		for _, mineral := range completeMinerals {
			if minMax < mineral.Geodes+daysRemaining*mineral.GeodeRobots {
				minMax = mineral.Geodes + daysRemaining*mineral.GeodeRobots
			}
		}
		for _, mineral := range minerals {
			if minMax < mineral.Geodes+daysRemaining*mineral.GeodeRobots {
				minMax = mineral.Geodes + daysRemaining*mineral.GeodeRobots
			}
		}
		newMinerals := []Minerals{}
		for _, mineral := range minerals {
			if cannotBuildGeodeRobotBeforeLastDay(mineral, blueprint, daysRemaining) {
				completeMinerals = append(completeMinerals, mineral)
			} else {
				progressedMinerals := progressMinerals(mineral, blueprint, daysRemaining)
			CheckMins:
				for _, progressedMineral := range progressedMinerals {
					if findMinMaxGeodes(progressedMineral, blueprint, daysRemaining-1) <= minMax {
						continue
					}
					for newMineralIndex, newMineral := range newMinerals {
						if inferior(progressedMineral, newMineral) {
							continue CheckMins
						} else if inferior(newMineral, progressedMineral) {
							newMinerals = append(newMinerals[:newMineralIndex], newMinerals[newMineralIndex+1:]...)
							newMinerals = append(newMinerals, progressedMineral)
							continue CheckMins
						}
					}
					newMinerals = append(newMinerals, progressedMineral)
				}
			}
		}
		for index := range completeMinerals {
			completeMinerals[index].Geodes += completeMinerals[index].GeodeRobots
		}
		minerals = newMinerals
	}
	max := 0
	for _, mineral := range completeMinerals {
		if mineral.Geodes > max {
			max = mineral.Geodes
		}
	}
	return max
}

func findMinMaxGeodes(minerals Minerals, blueprint Blueprint, daysRemaining int) int {
	minDaysToCompleteNextGeodeRobot := minDaysToGetGeodeResources(minerals, blueprint) + 1
	remainingDays := daysRemaining - minDaysToCompleteNextGeodeRobot
	if remainingDays < 0 {
		remainingDays = 0
	}
	total := minerals.Geodes
	for i := 1; i <= daysRemaining; i++ {
		total += minerals.GeodeRobots
	}
	for i := 1; i <= remainingDays; i++ {
		total += i
	}
	return total
}

func minDaysToGetGeodeResources(minerals Minerals, blueprint Blueprint) int {
	if minerals.Obsidian >= blueprint.GeodeRobotObsidianReq && minerals.Obsidian >= blueprint.GeodeRobotOreReq {
		return 0
	}
	obsidian := minerals.Obsidian
	minDaysToNewRobot := minDaysToGetObsidianResources(minerals, blueprint) + 1
	for i := 1; i >= 0; i++ {
		obsidian += minerals.ObsidianRobots
		if i > minDaysToNewRobot {
			obsidian += i - minDaysToNewRobot
		}
		if obsidian >= blueprint.GeodeRobotObsidianReq {
			return i
		}
	}
	return -1
}

func minDaysToGetObsidianResources(minerals Minerals, blueprint Blueprint) int {
	if minerals.Clay >= blueprint.ObsidianRobotClayReq && minerals.OreRobots >= blueprint.ObsidianRobotOreReq {
		return 0
	}
	clay := minerals.Clay
	minDaysToNewRobot := minDaysToGetClayResources(minerals, blueprint) + 1
	for i := 1; i >= 0; i++ {
		clay += minerals.ClayRobots
		if i > minDaysToNewRobot {
			clay += i - minDaysToNewRobot
		}
		if clay >= blueprint.ObsidianRobotClayReq {
			return i
		}
	}
	return -1
}

func minDaysToGetClayResources(minerals Minerals, blueprint Blueprint) int {
	if minerals.Ore >= blueprint.ClayRobotOreReq {
		return 0
	}
	ore := minerals.Ore
	for i := 1; i >= 0; i++ {
		ore += minerals.OreRobots + i
		if ore >= blueprint.ClayRobotOreReq {
			return i
		}
	}
	return -1
}

func cannotBuildGeodeRobotBeforeLastDay(minerals Minerals, blueprint Blueprint, daysRemaining int) bool {
	if daysRemaining == 1 {
		return true
	} else if daysRemaining == 2 {
		if minerals.Obsidian < blueprint.GeodeRobotObsidianReq || minerals.Ore < blueprint.GeodeRobotOreReq {
			return true
		} else {
			return false
		}
	} else {
		return minerals.Ore+triangleNumber(daysRemaining-3)+(daysRemaining-2)*minerals.ObsidianRobots < blueprint.GeodeRobotOreReq ||
			minerals.Obsidian+triangleNumber(daysRemaining-3)+(daysRemaining-2)*minerals.ObsidianRobots < blueprint.GeodeRobotObsidianReq
	}
}

func triangleNumber(value int) int {
	return value * (value + 1) / 2
}

func inferior(minerals1, minerals2 Minerals) bool {
	return minerals1.Ore <= minerals2.Ore && minerals1.Clay <= minerals2.Clay &&
		minerals1.Obsidian <= minerals2.Obsidian && minerals1.Geodes <= minerals2.Geodes &&
		minerals1.OreRobots <= minerals2.OreRobots && minerals1.ClayRobots <= minerals2.ClayRobots &&
		minerals1.ObsidianRobots <= minerals2.ObsidianRobots && minerals1.GeodeRobots <= minerals2.GeodeRobots
}

func progressMinerals(oldMinerals Minerals, blueprint Blueprint, daysRemaining int) []Minerals {
	newMinerals := findNewValues(oldMinerals, blueprint, daysRemaining)
	for index := range newMinerals {
		newMinerals[index].Ore += oldMinerals.OreRobots
		newMinerals[index].Clay += oldMinerals.ClayRobots
		newMinerals[index].Obsidian += oldMinerals.ObsidianRobots
		newMinerals[index].Geodes += oldMinerals.GeodeRobots
	}
	return newMinerals
}

func findNewValues(oldMinerals Minerals, blueprint Blueprint, daysRemaining int) []Minerals {
	newMinerals := []Minerals{}
	if daysRemaining > 2 {
		if !mustBuild(oldMinerals, blueprint) {
			newMinerals = append(newMinerals, oldMinerals)
		}
		if oldMinerals.Ore >= blueprint.OreRobotOreReq {
			newMin := oldMinerals
			newMin.OreRobots++
			newMin.Ore -= blueprint.OreRobotOreReq
			newMinerals = append(newMinerals, newMin)
		}
		if oldMinerals.Ore >= blueprint.ClayRobotOreReq {
			newMin := oldMinerals
			newMin.ClayRobots++
			newMin.Ore -= blueprint.ClayRobotOreReq
			newMinerals = append(newMinerals, newMin)
		}
		if oldMinerals.Ore >= blueprint.ObsidianRobotOreReq && oldMinerals.Clay >= blueprint.ObsidianRobotClayReq {
			newMin := oldMinerals
			newMin.ObsidianRobots++
			newMin.Ore -= blueprint.ObsidianRobotOreReq
			newMin.Clay -= blueprint.ObsidianRobotClayReq
			newMinerals = append(newMinerals, newMin)
		}
		if oldMinerals.Ore >= blueprint.GeodeRobotOreReq && oldMinerals.Obsidian >= blueprint.GeodeRobotObsidianReq {
			newMin := oldMinerals
			newMin.GeodeRobots++
			newMin.Ore -= blueprint.GeodeRobotOreReq
			newMin.Obsidian -= blueprint.GeodeRobotObsidianReq
			newMinerals = append(newMinerals, newMin)
		}
	} else if daysRemaining == 2 {
		if oldMinerals.Ore >= blueprint.GeodeRobotOreReq && oldMinerals.Obsidian >= blueprint.GeodeRobotObsidianReq {
			newMin := oldMinerals
			newMin.GeodeRobots++
			newMin.Ore -= blueprint.GeodeRobotOreReq
			newMin.Obsidian -= blueprint.GeodeRobotObsidianReq
			newMinerals = append(newMinerals, newMin)
		} else {
			newMinerals = append(newMinerals, oldMinerals)
		}
	}
	return newMinerals
}

func mustBuild(minerals Minerals, blueprint Blueprint) bool {
	if minerals.ClayRobots == 0 {
		return minerals.Ore >= blueprint.OreRobotOreReq && minerals.Ore >= blueprint.ClayRobotOreReq
	} else if minerals.ObsidianRobots == 0 {
		return minerals.Ore >= blueprint.OreRobotOreReq && minerals.Ore >= blueprint.ClayRobotOreReq &&
			minerals.Ore >= blueprint.ObsidianRobotOreReq && minerals.Clay >= blueprint.ObsidianRobotClayReq
	} else {
		return minerals.Ore >= blueprint.OreRobotOreReq && minerals.Ore >= blueprint.ClayRobotOreReq &&
			minerals.Ore >= blueprint.ObsidianRobotOreReq && minerals.Clay >= blueprint.ObsidianRobotClayReq &&
			minerals.Ore >= blueprint.GeodeRobotOreReq && minerals.Obsidian >= blueprint.GeodeRobotObsidianReq
	}
}

func readBlueprints(fileName string) ([]Blueprint, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	blueprints := []Blueprint{}
	for scanner.Scan() {
		splitString := strings.Split(scanner.Text(), " ")
		stringValues := []string{splitString[6], splitString[12], splitString[18], splitString[21], splitString[27], splitString[30]}
		values := []int{}
		for _, stringValue := range stringValues {
			value, err := strconv.Atoi(stringValue)
			if err != nil {
				return nil, err
			}

			values = append(values, value)
		}
		blueprint := Blueprint{OreRobotOreReq: values[0], ClayRobotOreReq: values[1], ObsidianRobotOreReq: values[2], ObsidianRobotClayReq: values[3], GeodeRobotOreReq: values[4], GeodeRobotObsidianReq: values[5]}
		blueprints = append(blueprints, blueprint)
	}

	return blueprints, nil
}
