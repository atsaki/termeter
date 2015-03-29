# termeter
Visualize data in the terminal

## Description

termeter can visualize data in the terminal. Data can be passed by pipe or file.

```bash
$  seq 100 | awk 'BEGIN{OFS="\t"; print "x","sin(x)","cos(x)"}{x=$1/10; print x,sin(x),cos(x)}' | termeter
```

![screenshot01](https://qiita-image-store.s3.amazonaws.com/0/15114/d838dbcd-5629-3f7c-da0e-710a899dac20.png)

You can even draw charts from streaming data.  

```bash
$ seq 300 | awk 'BEGIN{OFS="\t"; print "x","sin(x)","cos(x)"}{x=$1/10; print x,sin(x),cos(x); system("sleep 0.1")}' | termeter
```

<a href="https://asciinema.org/a/18127"><img src="https://asciinema.org/a/18127.png" /></a>

## Installation

```bash
$ go get github.com/atsaki/termeter/cmd/termeter
```

## Input Data

You can input data with stdin or file.

```bash
$ cat data.txt | termeter
$ termeter data.txt
```

termeter can accept tabular data like CSV. 
Delimiter character can be specified with option '-d DELIMITER'. Default is tab.

## Chart types

termeter supports following chart types. 

* LINE
  * Plot values as line plot
* COUNTER
  * Bar chart of frequencies
* CDF
  * Cumulative distribution function

By default, termeter choose chart type automatically from second line of data.
If value is numeric LINE is choosed. Otherwise, COUNTER is choosed. 

You can specify chart type with option ```-t TYPESTRING```.
nth character of TYPESTRING corresponds to nth chart type.
Following charcters can be used.

* l: LINE
* c: COUNTER
* d: CDF
* other: auto

### Example of chart types

```bash
$ (echo "line counter cdf"; seq 1 1000 | awk '{x=int(6*rand())+1; print x,x,x}') | termeter -d " " -t lcd -S numerical
```

![charttype](https://qiita-image-store.s3.amazonaws.com/0/15114/653ddf3a-bc0f-6f76-f39f-984bd33eaff4.png)

## Use case

It is useful to draw chart of resouce in the terminal.
You can use tools like [dstat](https://github.com/dagwieers/dstat).

```bash
$ dstat --cpu --output dstat.log > /dev/null &
$ tail -f -n +7 dstat.log | termeter -d ,
```

<a href="https://asciinema.org/a/18129"><img src="https://asciinema.org/a/18129.png" /></a>

## License

MIT
