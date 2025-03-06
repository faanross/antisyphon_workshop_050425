<template>
  <div>
    <!-- Connection Status -->
    <div class="status" :class="{ connected: isConnected }">
      WebSocket Status: {{ connectionStatus }}
    </div>

    <!-- Messages received from the server will be shown here -->
    <div v-if="messages.length > 0" class="messages">
      <ul>
        <li v-for="(message, index) in messages" :key="index">{{ message }}</li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import {ref, onMounted, onUnmounted} from 'vue';

// WebSocket connection
const socket = ref(null);
const isConnected = ref(false);
const connectionStatus = ref('Disconnected');
const messages = ref([]);

// Connect to WebSocket server
const connectWebSocket = () => {
  // Close existing connection if any
  if (socket.value) {
    socket.value.close();
  }

  // Create new WebSocket connection
  const wsUrl = 'ws://localhost:8080/ws';
  socket.value = new WebSocket(wsUrl);

  // Connection opened
  socket.value.addEventListener('open', (event) => {
    console.log('Connected to WebSocket server');
    isConnected.value = true;
    connectionStatus.value = 'Connected';

    // Send a message to the server
    socket.value.send('Hello from Vue client!');
  });

  // Listen for messages
  socket.value.addEventListener('message', (event) => {
    console.log('Message from server:', event.data);
    messages.value.push(event.data);
  });

  // Connection closed
  socket.value.addEventListener('close', (event) => {
    console.log('Disconnected from WebSocket server');
    isConnected.value = false;
    connectionStatus.value = 'Disconnected';
  });

  // Connection error
  socket.value.addEventListener('error', (event) => {
    console.error('WebSocket error:', event);
    connectionStatus.value = 'Error';
  });
};

// Connect on component mount
onMounted(() => {
  connectWebSocket();
});

// Clean up on component unmount
onUnmounted(() => {
  if (socket.value) {
    socket.value.close();
  }
});
</script>

<style scoped>
.status {
  margin: 20px 0;
  padding: 10px;
  background-color: #5e5e5e;
}

.status.connected {
  background-color: #5e5e5e;
}

.messages {
  margin-top: 5px;
  padding: 10px;
  border: 1px solid #ddd;
  background-color: #5e5e5e;
}

.messages ul {
  margin: 0;
  padding-left: 50px;
}
</style>