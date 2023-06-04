# Warehouse Robot

Warehouse Robot is a GoLang project that simulates the movement of two robots in a warehouse. 
The robots work together and move concurrently in the same direction. 
They follow instructions given by the points of the compass 
(N, E, S, W) and collect objects along the way.

## Instructions

The robots' movements are determined by the instructions provided. 
The goal is to write the directions for the robots to collect specific objects in the warehouse. 
The objects are represented by different symbols:

 - 't': Hat
 - 'l': Ball
 - 'x': Box
 - 'L': Empty space

The robots are represented as Robot A and Robot B. 
They collect objects based on the given instructions 
and move towards the objects in the same direction.

## Examples

Here are the directions for the robots to collect specific objects:

 1. Robot A: ball   Robot B: ball

   All Steps: NEE

 2. Robot A: hat, ball   Robot B: box, hat

   All Steps: NNSSE

 3. Robot A: ball, ball   Robot B: hat, ball

   All Steps: ENE

## Usage

 1. Ensure you have GoLang installed on your system.
 2. Clone the repository and navigate to the project directory.
 3. Run the following command to execute the program:

        go run warehouse_robot.go


The output will display the directions for the robots to collect the specified objects.
Feel free to modify the robot's starting locations, warehouse layout, 
and object collection instructions to explore different scenarios.

Happy robot navigating and object collection in the warehouse!
