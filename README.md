# termeter
Visualize data in the terminal


## Examples

ruby -e 'puts "x\tsin(x)\tcos(x)"; (1..1000).each{|i| x=i/10.0; puts"#{x}\t#{Math.sin(x)}\t#{Math.cos(x)}"; STDOUT.flush; sleep 0.1}' | LANG=C go run cmd/termeter/termeter.go -L first
ruby -e 'puts "x\ty\tx+y"; (1..1000).each{x=rand(1..6); y=rand(1..6); puts "#{x}\t#{y}\t#{x+y}"; STDOUT.flush; sleep 0.1}' | LANG=C go run cmd/termeter/termeter.go -t ccc -l 1,2,3,4,5,6,7,8,9,10,11,12 -T helloworld
