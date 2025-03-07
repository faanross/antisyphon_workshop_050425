<template>
  <div>
    <!-- Connection Status (this won't be visible due to the parent div being hidden) -->
    <div class="status" :class="{ connected: isConnected }">
      WebSocket Status: {{ connectionStatus }}
    </div>

    <!-- Messages (also not visible) -->
    <div v-if="messages.length > 0" class="messages">
      <ul>
        <li v-for="(message, index) in messages" :key="index">{{ message }}</li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import {ref, onMounted, onUnmounted, inject, defineEmits} from 'vue';

// Define emits for status changes and messages
const emit = defineEmits(['status-change', 'message-received']);

// Get the shared listeners state
const listenersState = inject('listenersState');

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

    // Emit status change
    emit('status-change', {connected: true, status: 'Connected'});

    // Send a message to the server
    socket.value.send('Hello from Vue client!');
  });

  // Listen for messages
  socket.value.addEventListener('message', (event) => {
    console.log('Message from server (raw):', event.data);

    // Add the raw message to messages list
    messages.value.push(event.data);

    // Emit message received
    emit('message-received', event.data);

    // Try to parse as JSON
    try {
      const data = JSON.parse(event.data);
      console.log('Parsed message data:', data);

      // Handle different message types
      if (data.type === 'listener_created') {
        // Use the shared state to add the new listener
        listenersState.addListener(data.payload);
        console.log('New listener added, current listeners:', listenersState.getListeners());
      } else if (data.type === 'listener_status') {
        // Use the shared state to update all listeners
        listenersState.updateListeners(data.payload);
        console.log('Updated listeners list:', listenersState.getListeners());
      }
    } catch (e) {
      console.log('Received non-JSON message:', e.message);
    }
  });

  // Connection closed
  socket.value.addEventListener('close', (event) => {
    console.log('Disconnected from WebSocket server');
    isConnected.value = false;
    connectionStatus.value = 'Disconnected';

    // Emit status change
    emit('status-change', {connected: false, status: 'Disconnected'});
  });

  // Connection error
  socket.value.addEventListener('error', (event) => {
    console.error('WebSocket error:', event);
    connectionStatus.value = 'Error';

    // Emit status change
    emit('status-change', {connected: false, status: 'Error'});
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

// Define what properties are exposed to the parent component
defineExpose({
  isConnected,
  connectionStatus,
  messages
});
</script>

<style scoped>
.status {
  margin: 20px 0;
  padding: 10px;
  background-color: #5e5e5e;
}

.status.connected {
  background-color: #4CAF50; /* Green for connected state */
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