# simulationExample

You can run the server using `docker-compose up -d` inside `fleetState` folder

You can build the vehicle simulator using inside the `dataSimulator` using `go build .` and then use it with paramter `number` to specify the number of vehicle simulations

You can then use the binary to simulate the vehicles use the argument `./dataSimulator`

You can view the stream of vehicle movement by building `displayData` using `go build .` inside the folder and then executing it with paramter `vin` to specify the vin of the vehicle to display


VIN of vehicles are created of the form simulated_<numbers starting from 0 to n>

