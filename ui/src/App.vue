<template>
  <div class="container">
    <h1>OrlokC2</h1>

    <!-- Hidden persistent WebSocket connection -->
    <div style="display: none;">
      <WebSocketConnection ref="wsConnection" @status-change="updateStatus"/>
    </div>

    <!-- Connection Status -->
    <div class="status" :class="{ connected: isConnected }">
      WebSocket Status: {{ connectionStatus }}
    </div>

    <!-- Listeners Table -->
    <ListenersTable/>
  </div>
</template>

<script setup>
import {ref, provide, onMounted} from 'vue';
import WebSocketConnection from './components/WebSocketConnection.vue';
import ListenersTable from './components/ListenersTable.vue';

// WebSocket connection reference
const wsConnection = ref(null);

// Status information
const isConnected = ref(false);
const connectionStatus = ref('Disconnected');

// Create a shared state for listeners that will be accessible to components
const sharedListeners = ref([]);

// Provide methods to access and update the listeners
provide('listenersState', {
  getListeners: () => sharedListeners.value,
  addListener: (listener) => {
    console.log('Adding listener to shared state:', listener);
    sharedListeners.value.push(listener);
  },
  updateListeners: (listeners) => {
    console.log('Updating all listeners in shared state:', listeners);
    sharedListeners.value = listeners;
  }
});

// Update status based on events from WebSocketConnection
function updateStatus(status) {
  isConnected.value = status.connected;
  connectionStatus.value = status.status;
}

// Log for debugging
onMounted(() => {
  console.log('App mounted, WebSocketConnection ref:', wsConnection.value);
});
</script>

<style>
.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}

h1 {
  text-align: center;
  margin-bottom: 30px;
}

.status {
  margin: 20px 0;
  padding: 10px;
  background-color: #5e5e5e;
  color: white;
  border-radius: 4px;
}

.status.connected {
  background-color: #4CAF50; /* Green color for connected state */
}
</style>