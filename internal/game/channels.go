package game

type StateEvent struct {
	FoodDelta   int
	WorkerDelta int
}

type GameChannels struct {
	EggChan   chan int
	LogChan   chan string
	StateChan chan StateEvent
	FoodQuery chan chan int
}
