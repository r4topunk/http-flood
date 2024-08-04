// worker.js

self.onmessage = function (event) {
  const { url, numberOfRequests } = event.data;

  let requestsSent = 0;

  function sendRequest() {
    fetch(url)
      .then((response) => {
        if (response.ok) {
          requestsSent++;
          self.postMessage(1); // Send a message back to the main script
        }
      })
      .catch((error) => {
        console.error("Request failed:", error);
      });
  }

  for (let i = 0; i < numberOfRequests; i++) {
    setTimeout(sendRequest, i * 10); // Send a request every 10ms
  }
};
