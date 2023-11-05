# hashring

Provides a simple implementation of a consistent hash ring.  By default it uses the 32-bit cyclic redundancy check (CRC-32) for calculating checksums, but can take any function that takes a byte slice and returns an unsigned 32bit integer.  The current implementation, is not as efficient as it could be and can also be prone to hotspots though it should consistently distribute across the nodes +/- 3 percent (dependent on the number of replicates).

## Install

```
go get stvz.io/hashring
```

## Usage

```go
import "stvz.io/hashring"

// Initialize the ring with 3 replicas, using the default checksum function and
// add some nodes to the ring.
ring := hashring.New(3, nil)
ring.Add("server1", "server2", "server3", "server4", "server5")
```

To get the node from the ring, use the `Get(key string)` function.  If there are no nodes available, the method returns an empty string.

```go
if node := ring.Get("id"); node != "" {
  sendFn(node, data)
}
```

For pull style distributed applications, the function `Mine(name, key string)` exists to determine whether or not the key that is being pulled belongs to the worker.  It returns a boolean value.  Assuming that the nodes use the hostname of the worker the following is an example:

```go
hostname, _ := os.Hostname()
if ring.Mine(hostname, "id") {
  processFn(data)
}
```

To remove a node (or nodes) use the `Remove(node string)` function.

```go
ring.Remove("server1", "server2")
```
