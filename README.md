# MetaSV stream connect performance

This program is to test the performance for connections on the new MetaSV stream api.

## Usage

Change the request header , add the api key to finish the test, if the key is not present, you will get 429

```
// todo replace the api-key
request.Header.Add("x-api-key", "ask-me-for-the-key")
request.Header.Set("Content-Type", "application/json")
```

Change the number of go routines to see how is the performance

## Test Result

Maximum 2000 active connections running on local PC. Getting out of memory locally if exceed this number.

So the server can handle at least 2000 stream connections
