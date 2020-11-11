package server

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/twitter/scoot/cloud/cluster"
	"github.com/twitter/scoot/common/log/tags"
	"github.com/twitter/scoot/common/stats"
	"github.com/twitter/scoot/runner"
	"github.com/twitter/scoot/scheduler/domain"
)

var (
	default_sc = SchedulerConfig{
		MaxRetriesPerTask:    0,
		DebugMode:            false,
		RecoverJobsOnStartup: false,
		DefaultTaskTimeout:   0,
		TaskTimeoutOverhead:  0,
		RunnerRetryTimeout:   0,
		RunnerRetryInterval:  0,
		ReadyFnBackoff:       0,
		MaxRequestors:        0,
		MaxJobsPerRequestor:  0,
		TaskThrottle:         0,
		Admins:               nil,
	}
)

type classState struct {
	loadPct              int
	numRunningTasks      int
	numWaitingTasks      int
	numJobs              int
	expectedTasksToStart int
	expectedTasksToStop  int
}

type testDef struct {
	totalWorkers int
	classes      map[string]classState
}

func Test_Class_Task_Start_Cnts(t *testing.T) {
	testsDefs := []testDef{
		{totalWorkers: 1000, classes: map[string]classState{
			"c0": {loadPct: 30, numRunningTasks: 200, numWaitingTasks: 290, numJobs: 10, expectedTasksToStart: 94},
			"c1": {loadPct: 25, numRunningTasks: 300, numWaitingTasks: 230, numJobs: 50, expectedTasksToStart: 0},
			"c2": {loadPct: 20, numRunningTasks: 0, numWaitingTasks: 150, numJobs: 20, expectedTasksToStart: 150},
			"c3": {loadPct: 15, numRunningTasks: 100, numWaitingTasks: 150, numJobs: 3, expectedTasksToStart: 46},
			"c4": {loadPct: 10, numRunningTasks: 110, numWaitingTasks: 90, numJobs: 2, expectedTasksToStart: 0},
			"c5": {loadPct: 0, numRunningTasks: 0, numWaitingTasks: 328, numJobs: 1, expectedTasksToStart: 0}},
		},
		{totalWorkers: 1000, classes: map[string]classState{
			"c0": {loadPct: 30, numRunningTasks: 200, numWaitingTasks: 290, numJobs: 10, expectedTasksToStart: 167},
			"c1": {loadPct: 25, numRunningTasks: 300, numWaitingTasks: 230, numJobs: 15, expectedTasksToStart: 53},
			"c2": {loadPct: 20, numRunningTasks: 0, numWaitingTasks: 0, numJobs: 0, expectedTasksToStart: 0},
			"c3": {loadPct: 15, numRunningTasks: 100, numWaitingTasks: 50, numJobs: 3, expectedTasksToStart: 50},
			"c4": {loadPct: 10, numRunningTasks: 110, numWaitingTasks: 90, numJobs: 2, expectedTasksToStart: 20}},
		},
		{totalWorkers: 1000, classes: map[string]classState{
			"c0": {loadPct: 30, numRunningTasks: 200, numWaitingTasks: 10, numJobs: 2, expectedTasksToStart: 10},
			"c1": {loadPct: 25, numRunningTasks: 300, numWaitingTasks: 230, numJobs: 15, expectedTasksToStart: 166},
			"c2": {loadPct: 20, numRunningTasks: 0, numWaitingTasks: 0, numJobs: 0, expectedTasksToStart: 0},
			"c3": {loadPct: 15, numRunningTasks: 100, numWaitingTasks: 50, numJobs: 10, expectedTasksToStart: 50},
			"c4": {loadPct: 10, numRunningTasks: 110, numWaitingTasks: 90, numJobs: 3, expectedTasksToStart: 64}},
		},
		{totalWorkers: 1000, classes: map[string]classState{
			"c0": {loadPct: 30, numRunningTasks: 0, numWaitingTasks: 300, numJobs: 30, expectedTasksToStart: 105},
			"c1": {loadPct: 25, numRunningTasks: 0, numWaitingTasks: 230, numJobs: 10, expectedTasksToStart: 81},
			"c2": {loadPct: 20, numRunningTasks: 0, numWaitingTasks: 400, numJobs: 40, expectedTasksToStart: 66},
			"c3": {loadPct: 15, numRunningTasks: 0, numWaitingTasks: 650, numJobs: 13, expectedTasksToStart: 48},
			"c4": {loadPct: 10, numRunningTasks: 700, numWaitingTasks: 800, numJobs: 40, expectedTasksToStart: 0}},
		},
		{totalWorkers: 1000, classes: map[string]classState{
			"c0": {loadPct: 35, numRunningTasks: 200, numWaitingTasks: 100, numJobs: 30, expectedTasksToStart: 100},
			"c1": {loadPct: 30, numRunningTasks: 300, numWaitingTasks: 50, numJobs: 10, expectedTasksToStart: 0},
			"c2": {loadPct: 20, numRunningTasks: 0, numWaitingTasks: 200, numJobs: 40, expectedTasksToStart: 159},
			"c3": {loadPct: 0, numRunningTasks: 100, numWaitingTasks: 300, numJobs: 13, expectedTasksToStart: 0},
			"c4": {loadPct: 15, numRunningTasks: 110, numWaitingTasks: 500, numJobs: 40, expectedTasksToStart: 31}},
		},
		{totalWorkers: 10000, classes: map[string]classState{
			"c0": {loadPct: 30, numRunningTasks: 1660, numWaitingTasks: 14220, numJobs: 300, expectedTasksToStart: 830},
			"c1": {loadPct: 25, numRunningTasks: 101, numWaitingTasks: 9401, numJobs: 100, expectedTasksToStart: 1282},
			"c2": {loadPct: 16, numRunningTasks: 420, numWaitingTasks: 16542, numJobs: 400, expectedTasksToStart: 641},
			"c3": {loadPct: 14, numRunningTasks: 14, numWaitingTasks: 4194, numJobs: 13, expectedTasksToStart: 754},
			"c4": {loadPct: 6, numRunningTasks: 404, numWaitingTasks: 15944, numJobs: 400, expectedTasksToStart: 76},
			"c5": {loadPct: 4, numRunningTasks: 42, numWaitingTasks: 11136, numJobs: 40, expectedTasksToStart: 187},
			"c6": {loadPct: 3, numRunningTasks: 977, numWaitingTasks: 9145, numJobs: 30, expectedTasksToStart: 0},
			"c7": {loadPct: 2, numRunningTasks: 2612, numWaitingTasks: 16781, numJobs: 40, expectedTasksToStart: 0}},
		},
		{totalWorkers: 10000, classes: map[string]classState{
			"c0": {loadPct: 30, numRunningTasks: 1660, numWaitingTasks: 14220, numJobs: 300, expectedTasksToStart: 830},
			"c1": {loadPct: 25, numRunningTasks: 101, numWaitingTasks: 29401, numJobs: 100, expectedTasksToStart: 1282},
			"c2": {loadPct: 16, numRunningTasks: 420, numWaitingTasks: 16542, numJobs: 400, expectedTasksToStart: 641},
			"c3": {loadPct: 14, numRunningTasks: 14, numWaitingTasks: 104194, numJobs: 13, expectedTasksToStart: 754},
			"c4": {loadPct: 6, numRunningTasks: 404, numWaitingTasks: 15944, numJobs: 400, expectedTasksToStart: 76},
			"c5": {loadPct: 4, numRunningTasks: 42, numWaitingTasks: 11136, numJobs: 40, expectedTasksToStart: 187},
			"c6": {loadPct: 3, numRunningTasks: 977, numWaitingTasks: 209145, numJobs: 30, expectedTasksToStart: 0},
			"c7": {loadPct: 2, numRunningTasks: 2612, numWaitingTasks: 416781, numJobs: 40, expectedTasksToStart: 0}},
		},
	}

	statsRegistry := stats.NewFinagleStatsRegistry()
	statsReceiver, _ := stats.NewCustomStatsReceiver(func() stats.StatsRegistry { return statsRegistry }, 0)

	config := &LoadBasedAlgConfig{stat: statsReceiver, minRebalanceTime: DefaultMinRebalanceTime}
	lbs := NewLoadBasedAlg(config, nil)

	runTests(t, testsDefs, lbs)
}

func runTests(t *testing.T, testsDefs []testDef, lbs *LoadBasedAlg) {
	jobsByJobID := map[string]*jobState{}
	for _, testDef := range testsDefs {
		// reinitialize the task start times since this test will be creating new tasks
		lbs.tasksByJobClassAndStartTimeSec = tasksByClassAndStartTimeSec{}
		totalWorkers := testDef.totalWorkers
		usedWorkers := 0
		jobsByRequestor := map[string][]*jobState{}
		requestorToClass := map[string]string{}
		loadPcts := map[string]int32{}
		expectedNumTasks := 0
		expectedNumStopTasks := 0
		for className, state := range testDef.classes {
			usedWorkers += state.numRunningTasks
			js := makeJobStatesFromClassStates(t, className, state, jobsByJobID, lbs.tasksByJobClassAndStartTimeSec)
			if len(js) > 0 {
				jobsByRequestor[js[0].Job.Def.Requestor] = js
				requestorToClass[js[0].Job.Def.Requestor] = className
			}
			loadPcts[className] = int32(state.loadPct)
			expectedNumTasks += state.expectedTasksToStart
			expectedNumStopTasks += state.expectedTasksToStop
		}

		cluster := &clusterState{
			updateCh:         nil,
			nodes:            nil,
			suspendedNodes:   nil,
			offlinedNodes:    nil,
			nodeGroups:       makeIdleGroup(totalWorkers),
			maxLostDuration:  0,
			maxFlakyDuration: 0,
			readyFn:          nil,
			numRunning:       usedWorkers,
			stats:            nil,
		}
		cluster.nodes = cluster.nodeGroups["idle"].idle

		lbs.SetClassLoadPcts(loadPcts)
		lbs.SetRequestorToClassMap(requestorToClass)

		tasksToBeAssigned, stopTasks := lbs.GetTasksToBeAssigned(nil, lbs.config.stat, cluster, jobsByRequestor)

		assert.Equal(t, expectedNumTasks, len(tasksToBeAssigned), "wrong number of tasks in tasksToBeAssigned")
		assert.Equal(t, expectedNumStopTasks, len(stopTasks))

		// compute the number of tasks to start for each class from the tasks list
		numTasksByClassName := map[string]int{}
		for _, task := range tasksToBeAssigned {
			jobState := jobsByJobID[task.JobId]
			className := GetRequestorClass(jobState.Job.Def.Requestor, lbs.requestorReToClassMap)
			if _, ok := numTasksByClassName[className]; !ok {
				numTasksByClassName[className] = 1
			} else {
				numTasksByClassName[className]++
			}
		}

		// compute the number of tasks to stop for each class from the tasks list
		numStopTasksByClassName := map[string]int{}
		for _, task := range stopTasks {
			jobState := jobsByJobID[task.JobId]
			if jobState == nil {
				log.Errorf("************ jobState is nil")
			}
			if jobState.Job == nil {
				log.Errorf("************ jobState.Job is nil in %v", jobState)
			}
			className := GetRequestorClass(jobState.Job.Def.Requestor, lbs.requestorReToClassMap)
			if _, ok := numStopTasksByClassName[className]; !ok {
				numStopTasksByClassName[className] = 1
			} else {
				numStopTasksByClassName[className]++
			}
		}

		// verify we've computed the number of tasks to start for each task correctly and
		// have the correct number of tasks in the task list for each class
		for className, state := range testDef.classes {
			// verify the computed number of tasks to start for the class
			assert.Equal(t, state.expectedTasksToStart, numTasksByClassName[className], "wrong number of %s tasks in the task list", className)
			assert.Equal(t, state.expectedTasksToStop, numStopTasksByClassName[className], "wrong number of %s tasks to stop in the task list", className)
		}
	}
}
func TestEmptyRequestor(t *testing.T) {
	statsRegistry := stats.NewFinagleStatsRegistry()
	statsReceiver, _ := stats.NewCustomStatsReceiver(func() stats.StatsRegistry { return statsRegistry }, 0)

	tasksByClassAndStartMap := tasksByClassAndStartTimeSec{}

	jobsByJobID := map[string]*jobState{}
	jobsByRequestor := map[string][]*jobState{}
	jobsByRequestor[""] = makeJobStatesFromClassStates(t, "", classState{numRunningTasks: 0, numWaitingTasks: 215, numJobs: 37}, jobsByJobID,
		tasksByClassAndStartMap)

	cluster := &clusterState{
		updateCh:         nil,
		nodes:            nil,
		suspendedNodes:   nil,
		offlinedNodes:    nil,
		nodeGroups:       makeIdleGroup(0),
		maxLostDuration:  0,
		maxFlakyDuration: 0,
		readyFn:          nil,
		numRunning:       0,
		stats:            nil,
	}
	cluster.nodes = cluster.nodeGroups["idle"].idle

	config := &LoadBasedAlgConfig{stat: statsReceiver, minRebalanceTime: DefaultMinRebalanceTime}
	lbs := NewLoadBasedAlg(config, tasksByClassAndStartMap)
	lbs.SetClassLoadPcts(DefaultLoadBasedSchedulerClassPcts)
	lbs.SetRequestorToClassMap(DefaultRequestorToClassMap)

	tasksToBeAssigned, stopTasks := lbs.GetTasksToBeAssigned(nil, statsReceiver, cluster, jobsByRequestor)

	assert.Equal(t, 0, len(tasksToBeAssigned), "wrong number of tasks in tasksToBeAssigned")
	assert.Nil(t, stopTasks)
}

// TestRandomScenario generate random tests with 10k workers and verify that
// the idle workers are allocated
func TestRandomScenario(t *testing.T) {
	// set up the test scenario: set up 2 classes to get 75% of workers, then create random % for the remaining 25%
	loadPcts := generatePcts()

	aTest := testDef{totalWorkers: 10000, classes: map[string]classState{}}
	totalWorkers := aTest.totalWorkers
	// define a random set of class states for the loadPcts defined above
	// these classes will use up a random number of workers (not to exceed 5000) and the number of waiting
	// tasks for each class will be a random number, not to exceed 2 times the total number of workers
	workersToUse := totalWorkers - rand.Intn(5001)
	totalWaitingTasks := 0
	for className := range loadPcts {
		numRunningTasks := 0
		if workersToUse > 0 {
			numRunningTasks = rand.Intn(workersToUse + 1)
		}
		waitingTasks := rand.Intn(totalWorkers * 2)
		aTest.classes[className] = classState{numRunningTasks: numRunningTasks, numWaitingTasks: waitingTasks, numJobs: max(1, min(100, numRunningTasks))}
		workersToUse -= numRunningTasks
		totalWaitingTasks += waitingTasks
	}

	tasksByJobClassAndStartMap := tasksByClassAndStartTimeSec{}

	// create jobState objects for each class
	usedWorkers := 0
	jobsByRequestor := map[string][]*jobState{}
	jobsByJobID := map[string]*jobState{}
	requestorToClass := map[string]string{}
	for className, state := range aTest.classes {
		usedWorkers += state.numRunningTasks
		js := makeJobStatesFromClassStates(t, className, state, jobsByJobID, tasksByJobClassAndStartMap)
		jobsByRequestor[js[0].Job.Def.Requestor] = js
		requestorToClass[js[0].Job.Def.Requestor] = className
	}

	cluster := &clusterState{
		updateCh:         nil,
		nodes:            nil,
		suspendedNodes:   nil,
		offlinedNodes:    nil,
		nodeGroups:       makeIdleGroup(totalWorkers),
		maxLostDuration:  0,
		maxFlakyDuration: 0,
		readyFn:          nil,
		numRunning:       usedWorkers,
		stats:            nil,
	}
	cluster.nodes = cluster.nodeGroups["idle"].idle

	statsRegistry := stats.NewFinagleStatsRegistry()
	statsReceiver, _ := stats.NewCustomStatsReceiver(func() stats.StatsRegistry { return statsRegistry }, 0)

	// run the test
	config := &LoadBasedAlgConfig{stat: statsReceiver, minRebalanceTime: DefaultMinRebalanceTime}
	lbs := NewLoadBasedAlg(config, tasksByJobClassAndStartMap)
	lbs.SetClassLoadPcts(loadPcts)
	lbs.SetRequestorToClassMap(requestorToClass)
	tasks, stopTasks := lbs.GetTasksToBeAssigned(nil, statsReceiver, cluster, jobsByRequestor)

	// verify the results: we don't know what to expect for each class (since it was randomly generated), so
	// just verify that the number of tasks being started equals the min of number of idle workers, total waiting tasks
	numTasksStarting := 0
	for className := range aTest.classes {
		numTasksStarting += lbs.getNumTasksToStart(className)
	}

	expectedNumTasks := min(totalWorkers-usedWorkers, totalWaitingTasks)
	assert.Equal(t, totalWorkers-usedWorkers, numTasksStarting)
	assert.Equal(t, expectedNumTasks, len(tasks))
	assert.Nil(t, stopTasks)
}

func Test_Rebalance(t *testing.T) {
	testsDefs := []testDef{
		{totalWorkers: 10, classes: map[string]classState{ // debuggable scenario
			"c0": {loadPct: 70, numRunningTasks: 2, numWaitingTasks: 20, numJobs: 3, expectedTasksToStart: 5},
			"c1": {loadPct: 20, numRunningTasks: 4, numWaitingTasks: 10, numJobs: 2, expectedTasksToStart: 0, expectedTasksToStop: 2},
			"c2": {loadPct: 10, numRunningTasks: 4, numWaitingTasks: 30, numJobs: 4, expectedTasksToStart: 0, expectedTasksToStop: 3}},
		},
		{totalWorkers: 10000, classes: map[string]classState{ // no rebalance - spread not large enough
			"c0": {loadPct: 30, numRunningTasks: 1660, numWaitingTasks: 14220, numJobs: 300, expectedTasksToStart: 830},
			"c1": {loadPct: 25, numRunningTasks: 101, numWaitingTasks: 9401, numJobs: 100, expectedTasksToStart: 1282},
			"c2": {loadPct: 16, numRunningTasks: 420, numWaitingTasks: 16542, numJobs: 400, expectedTasksToStart: 641},
			"c3": {loadPct: 14, numRunningTasks: 14, numWaitingTasks: 104194, numJobs: 13, expectedTasksToStart: 754},
			"c4": {loadPct: 6, numRunningTasks: 404, numWaitingTasks: 5944, numJobs: 400, expectedTasksToStart: 76},
			"c5": {loadPct: 4, numRunningTasks: 42, numWaitingTasks: 11136, numJobs: 40, expectedTasksToStart: 187},
			"c6": {loadPct: 3, numRunningTasks: 977, numWaitingTasks: 9145, numJobs: 30, expectedTasksToStart: 0},
			"c7": {loadPct: 2, numRunningTasks: 2612, numWaitingTasks: 16781, numJobs: 40, expectedTasksToStart: 0}},
		},
		{totalWorkers: 10000, classes: map[string]classState{ // rebalance, but no loaning workers
			"c0": {loadPct: 30, numRunningTasks: 166, numWaitingTasks: 14220, numJobs: 300, expectedTasksToStart: 2834},
			"c1": {loadPct: 25, numRunningTasks: 101, numWaitingTasks: 9401, numJobs: 100, expectedTasksToStart: 2399},
			"c2": {loadPct: 16, numRunningTasks: 420, numWaitingTasks: 16542, numJobs: 400, expectedTasksToStart: 1180},
			"c3": {loadPct: 14, numRunningTasks: 14, numWaitingTasks: 104194, numJobs: 13, expectedTasksToStart: 1386},
			"c4": {loadPct: 6, numRunningTasks: 404, numWaitingTasks: 15944, numJobs: 400, expectedTasksToStart: 196},
			"c5": {loadPct: 4, numRunningTasks: 42, numWaitingTasks: 11136, numJobs: 40, expectedTasksToStart: 358},
			"c6": {loadPct: 3, numRunningTasks: 977, numWaitingTasks: 209145, numJobs: 30, expectedTasksToStart: 0, expectedTasksToStop: 677},
			"c7": {loadPct: 2, numRunningTasks: 2612, numWaitingTasks: 416781, numJobs: 40, expectedTasksToStart: 0, expectedTasksToStop: 2412}},
		},
		{totalWorkers: 10000, classes: map[string]classState{ // rebalance and loan workers
			"c0": {loadPct: 30, numRunningTasks: 166, numWaitingTasks: 14220, numJobs: 300, expectedTasksToStart: 3034},
			"c1": {loadPct: 25, numRunningTasks: 101, numWaitingTasks: 9401, numJobs: 100, expectedTasksToStart: 2549},
			"c2": {loadPct: 16, numRunningTasks: 420, numWaitingTasks: 16542, numJobs: 400, expectedTasksToStart: 1275},
			"c3": {loadPct: 14, numRunningTasks: 14, numWaitingTasks: 104194, numJobs: 13, expectedTasksToStart: 1470},
			"c4": {loadPct: 6, numRunningTasks: 404, numWaitingTasks: 0, numJobs: 400, expectedTasksToStart: 0},
			"c5": {loadPct: 4, numRunningTasks: 42, numWaitingTasks: 0, numJobs: 40, expectedTasksToStart: 0},
			"c6": {loadPct: 3, numRunningTasks: 977, numWaitingTasks: 209145, numJobs: 30, expectedTasksToStart: 0, expectedTasksToStop: 660},
			"c7": {loadPct: 2, numRunningTasks: 2612, numWaitingTasks: 416781, numJobs: 40, expectedTasksToStart: 0, expectedTasksToStop: 2404}},
		},
	}

	statsRegistry := stats.NewFinagleStatsRegistry()
	statsReceiver, _ := stats.NewCustomStatsReceiver(func() stats.StatsRegistry { return statsRegistry }, 0)

	config := &LoadBasedAlgConfig{stat: statsReceiver, minRebalanceTime: 0 * time.Second}
	lbs := NewLoadBasedAlg(config, nil)

	runTests(t, testsDefs, lbs)
}

func generatePcts() map[string]int32 {
	// set up the test scenario: set up 2 classes to get 75% of workers, then create random % for the remaining 25%
	rand.Seed(time.Now().UnixNano())
	loadPcts := map[string]int32{
		"c0": 50,
		"c1": 25,
	}
	// generate random %s to make up the remaining 25 %
	i := 2
	var remainingPct int32 = 25
	for remainingPct > 0 {
		var pct int32
		if remainingPct < 3 {
			pct = remainingPct
		} else {
			pct = int32(rand.Intn(10)) // pct will be 0-9
		}
		loadPcts[fmt.Sprintf("c%d", i)] = pct
		i++
		remainingPct -= pct
	}

	return loadPcts
}

// makeJobStateFromClassStates make a list of jobStates for the class.  The classState will contain the number of
// jobStates to create and the total number of running and waiting tasks to distribute across the jobStates.
func makeJobStatesFromClassStates(t *testing.T, className string, cState classState, jobsByJobID map[string]*jobState,
	tasksByClassAndStartMap tasksByClassAndStartTimeSec) []*jobState {
	jobStates := make([]*jobState, cState.numJobs)
	requestor := fmt.Sprintf("requestor%s", className)

	var runningTasksDist []int
	if cState.numRunningTasks > 0 && cState.numJobs > 0 {
		runningTasksDist = createTaskDistribution(cState.numRunningTasks, cState.numJobs)
	}
	var waitingTasksDist []int
	if cState.numWaitingTasks != 0 && cState.numJobs > 0 {
		waitingTasksDist = createTaskDistribution(cState.numWaitingTasks, max(1, cState.numJobs-1)) // last job will have 0 waiting tasks unless there is only 1 job
	}

	totalWaitingTasks := 0
	totalRunningTasks := 0
	for i := 0; i < cState.numJobs; i++ {
		var numRunningTasks int = 0
		if cState.numRunningTasks != 0 {
			numRunningTasks = runningTasksDist[i]
		}
		var numWaitingTasks int = 0
		if cState.numWaitingTasks != 0 {
			if i < len(waitingTasksDist) {
				numWaitingTasks = waitingTasksDist[i]
			}
		}
		j := &domain.Job{
			Id: fmt.Sprintf("job_%s_%d", className, i),
			Def: domain.JobDefinition{
				JobType:   "dummyJobType",
				Requestor: requestor,
				Basis:     "",
				Tag:       "",
				Priority:  domain.Priority(0),
			},
		}
		js := &jobState{
			Job:                            j,
			Saga:                           nil,
			EndingSaga:                     false,
			TasksCompleted:                 0,
			TasksRunning:                   numRunningTasks,
			JobKilled:                      false,
			TimeCreated:                    time.Time{},
			TimeMarker:                     time.Time{},
			Completed:                      make(map[string]*taskState),
			Running:                        nil,
			NotStarted:                     nil,
			jobClass:                       className,
			tasksByJobClassAndStartTimeSec: tasksByClassAndStartMap,
		}
		totalRunningTasks += numRunningTasks
		_, tStates := makeTestTasks(j.Id, numRunningTasks)
		startSeed := time.Now()
		for i, rtState := range tStates {
			// distribute task start times across prior 10 minutes
			timeDelta := time.Duration(time.Minute * time.Duration(i%3))
			startTime := startSeed.Add(-1 * timeDelta)
			rtState.TimeStarted = startTime
			js.addTaskToStartTimeMap(className, rtState, startTime.Truncate(time.Second))
		}

		t, ts := makeTestTasks(j.Id, numWaitingTasks)
		j.Def.Tasks = t
		js.Tasks = ts
		taskMap := makeTaskMap(ts)
		js.NotStarted = taskMap
		jobStates[i] = js
		jobsByJobID[js.Job.Id] = js

		totalWaitingTasks += len(taskMap)
	}

	if (cState.numWaitingTasks != totalWaitingTasks) || (cState.numRunningTasks != totalRunningTasks) {
		// this is an error print the configuration so we can debug it
		assert.Equal(t, cState.numWaitingTasks, totalWaitingTasks, "invalid test setup, did not create correct number of waiting tasks for %s", className)
		assert.Equal(t, cState.numRunningTasks, totalRunningTasks, "invalid test setup, did not create correct number of running tasks for %s", className)
	}

	return jobStates
}

// createTaskDistribution a distribution of n tasks to m jobs.  Create list of m integers such that the sum
// of the integers add up to n.
func createTaskDistribution(nTasks int, mJobs int) []int {
	rand.Seed(time.Now().UnixNano())

	// over m iterations generate random numbers making sure the sum doesn't go over nTasks
	totalTaskCnt := 0
	taskCnts := []int{}
	if nTasks < mJobs {
		for i := 0; i < mJobs; i++ {
			if totalTaskCnt < nTasks {
				taskCnts = append(taskCnts, 1)
				totalTaskCnt++
			} else {
				taskCnts = append(taskCnts, 0)
			}
		}
	} else {
		aveTasksPerJob := int(nTasks/mJobs) * 2
		// generate random numbers up to aveTasksPerJob, forcing the final entries to add up to the sum
		for i := 0; i < mJobs; i++ {
			t := rand.Intn(aveTasksPerJob) + 1
			if (nTasks - (totalTaskCnt + t)) <= (mJobs - i) {
				taskCnts = append(taskCnts, 1)
				totalTaskCnt++
			} else {
				taskCnts = append(taskCnts, t)
				totalTaskCnt += t
			}
		}
		if totalTaskCnt < nTasks {
			taskCnts[len(taskCnts)-1] += nTasks - totalTaskCnt
		}
	}
	return taskCnts
}

func makeTestTasks(jobId string, numTasks int) ([]domain.TaskDefinition, []*taskState) {
	tasks := make([]domain.TaskDefinition, int(numTasks))
	tasksState := make([]*taskState, int(numTasks))
	for i := 0; i < numTasks; i++ {
		td := runner.Command{
			Argv:           []string{""},
			EnvVars:        nil,
			Timeout:        0,
			SnapshotID:     "",
			LogTags:        tags.LogTags{TaskID: fmt.Sprintf("%d", i), Tag: "dummyTag"},
			ExecuteRequest: nil,
		}
		tasks[i] = domain.TaskDefinition{Command: td}

		tasksState[i] = &taskState{
			JobId:  jobId,
			TaskId: fmt.Sprintf("task%d", i),
			Status: domain.NotStarted,
		}
	}

	return tasks, tasksState
}

func makeIdleGroup(n int) map[string]*nodeGroup {
	idle := make(map[cluster.NodeId]*nodeState)
	for i := 0; i < n; i++ {
		idle[cluster.NodeId(fmt.Sprintf("node%d", i))] = &nodeState{}
	}
	idleGroup := &nodeGroup{}
	idleGroup.idle = idle

	rVal := make(map[string]*nodeGroup)
	rVal["idle"] = idleGroup
	return rVal
}

func makeTaskMap(tasks []*taskState) map[string]*taskState {
	rVal := make(map[string]*taskState)
	for _, t := range tasks {
		rVal[t.TaskId] = t
	}
	return rVal
}

func ppTasksByJobClassAndStartTimeSec(tasksByJobClassAndStartTimeSec map[string]map[time.Time]map[string]*taskState) {
	log.Infof("********** tasks by job class and start time")
	for k, v := range tasksByJobClassAndStartTimeSec {
		log.Infof("class:%s has %d time buckets:", k, len(v))
		for timeKey, tasks := range v {
			log.Infof("%s has %d tasks", timeKey.Format("2006-01-02 15:04:05 MST"), len(tasks))
		}
	}
}