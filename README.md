# JobScheduler API Golang library

This is Golang library for SOS JobScheduler API.

# Requirements

* JobScheduler >= 1.10

# Install

```
$ go get github.com/ike-dai/go-jobscheduler
```

# Usage

Info:
* JobScheduler Web API URL -> http://localhost:4444
* standalone job name -> test/test_job
* jobchain name -> test/test_jobchain

Example1: Start Standalone Job


```
package main

import (
	"github.com/ike-dai/go-jobscheduler/jobscheduler"
)

func main() {
	client = jobscheduler.NewClient("http://localhost:4444"))
	params := &jobscheduler.StartJobInput{Job: "test/test_job"}
	_,err := client.StartJob(params)
	if err != nil {
		fmt.Printf("[ERROR] start job failure: %s \n", err.text")
	}
}

```

# Contact

Please send feedback to me.

Daisuke IKEDA

Twitter: [@ike_dai](https://twitter.com/ike_dai)

e-mail: <dai.ikd123@gmail.com>

# License

Licensed under the Apache License, Version 2.0. The Apache v2 full text is published at this [link](http://www.apache.org/licenses/LICENSE-2.0).

Copyright 2016 Daisuke IKEDA.
