# Century
Golang game server for 21st Century (prototype)

## Description

It's just a skeleton of what I call "Next-generation game server". But it's a runnable prototype, with basic error handlings, and benchmarkable performance maturity.

I started with a chat application scenario, but you can simply modify it to match the need of any realtime game server, or even general realtime server, just by replacing the "broadcast" method with your own "ProcessPacket" method.

It's both simple enough and complete enough to demonstrate the next-generation network programs, which are expected to have the feature of --

## Feature

* High throughput
* High concurrency
* (Automatic) High scalability, especially on many-core computers. (Think of 64-core computers, as much as 4-core ones.)

## Structure

`century.go` is the main server source file.

It contains a benchmarker named `chat_bencher.go`.

## 

## Detailed Information

You can find a even simpler chat server on:

[https://gist.github.com/drewolson/3950226](https://gist.github.com/drewolson/3950226)

(In fact I started my scratch from that.)

----------------

If you are looking for a "real" golang game server, you may find the following repos helpful:

* [gonet/2](https://github.com/gonet2) (website: [gonet/2 website](http://gonet2.github.io/))
* [gonet](https://github.com/xtaci/gonet) (which is the predecessor of the above one.)

(Both are described in **Chinese only**.)

And also less mature (IMO) one:

* [go4game](https://github.com/kasworld/go4game)

----------------

Q: Why a chat server?

A: Many such kinds of server frameworks choose chat server to demo usage, such as Boost.asio, Node.js, [Pomelo distributed game server](https://github.com/NetEase/pomelo), etc.



## Benchmark Result

```
Benchmarking: 127.0.0.1:6666
3 clients, running 8 bytes, 5 sec.

Speed: 85009 request/sec, 56940 response/sec
Requests: 425048
Responses: 284704
```

(Explain: I tuned the bencher not to wait for all the responses.)
