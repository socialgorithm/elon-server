package simulator

import (
	"log"
	"time"

	"github.com/socialgorithm/elon-server/domain"
	"github.com/socialgorithm/elon-server/track"
)

func StartSimulation() (<-chan domain.CarState, chan<- domain.CarControlState) {
	track := track.GenTrack()
	carControlStateChannel := make(chan domain.CarControlState)
	carStateChannel := startCarStateStream(track)
	return carStateChannel, carControlStateChannel
}

func startCarStateStream(track domain.Track) <-chan domain.CarState {
	carStateChannel := make(chan domain.CarState)
	go func() {
		for {
			log.Println("Generating new car state")
			carStateChannel <- genRandomCarState()
			time.Sleep(time.Second)
		}
	}()

	return carStateChannel
}

func genRandomCarState() domain.CarState {
	return domain.CarState{
		Position: domain.Position{X: 1.0, Y: 1.0},
		Velocity: 1,
		Sensors: []domain.Sensor{
			domain.Sensor{Angle: 1, Distance: 1},
		},
	}
}