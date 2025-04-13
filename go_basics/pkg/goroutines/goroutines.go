package goroutines

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var mutex = sync.Mutex{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var saveDBValues []string

func Init() {
	t0 := time.Now()
	for _, value := range dbData {
		wg.Add(1)
		go dbCall(value)
	}

	wg.Wait()
	fmt.Printf("\n Total execution time: %v", time.Since(t0))
	fmt.Printf("\n Db values obtained %v", saveDBValues)
}

func dbCall(dbValue string) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)

	// updating values into slice by acquiring lock first
	// if locks removed results are pretty scary
	mutex.Lock()
	saveDBValues = append(saveDBValues, dbValue)
	mutex.Unlock()
	wg.Done()
}
