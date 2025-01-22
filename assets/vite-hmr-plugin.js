// Function to check if server is ready
async function isServerReady() {
  try {
    const response = await fetch("/health");
    return response.ok;
  } catch {
    return false;
  }
}

// Function to wait for server with timeout
async function waitForServer(maxRetries = 10, interval = 150) {
  // sleep for 200ms to wait for debounce
  await new Promise((resolve) => setTimeout(resolve, 500));
  for (let i = 0; i < maxRetries; i++) {
    if (await isServerReady()) {
      return true;
    }
    await new Promise((resolve) => setTimeout(resolve, interval));
  }
  return false;
}

export default function templHmrPlugin() {
  return {
    name: "wait-for-server-hmr",
    async handleHotUpdate({ server, modules }) {
      if (!modules.length) return [];
      // Wait for server to be ready before reloading
      await waitForServer();
      // Reload the page
      server.ws.send({
        type: "full-reload",
      });
      return [];
    },
  };
}
