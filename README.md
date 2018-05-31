# simple_websocket
websocket, golang, echo framework 


```

// Create a new WebSocket
ws = new WebSocket((window.location.protocol == 'http') ? 'ws://' : 'ws://' +  '127.0.0.1:1323' + '/ws')
// Make it show an alert when a message is received
ws.onmessage = function(message) {
  alert(message.data);

}
// Send a new message when the WebSocket opens
ws.onopen = function() {
  ws.send('Hello, world');

}


```
