<template>
  <div class="container">
    <h1>OrlokC2</h1>

    <!-- Hidden persistent WebSocket connection -->
    <div style="display: none;">
      <WebSocketConnection ref="wsConnection" @status-change="updateStatus" @message-received="addMessage"/>
    </div>

    <!-- Tabs Navigation -->
    <div class="tabs-container">
      <div class="tabs">
        <button
            :class="{ active: activeTab === 'status' }"
            @click="activeTab = 'status'"
        >
          Status
        </button>
        <button
            :class="{ active: activeTab === 'listeners' }"
            @click="activeTab = 'listeners'"
        >
          Listeners
        </button>
      </div>
      <div class="tab-line"></div>
    </div>

    <!-- Tab Content -->
    <div class="tab-content">
      <!-- Status Tab - Now just shows connection status info, not the component itself -->
      <div v-if="activeTab === 'status'" class="tab-pane">
        <div class="status" :class="{ connected: isConnected }">
          WebSocket Status: {{ connectionStatus }}
        </div>
        <div v-if="messages.length > 0" class="messages">
          <ul>
            <li v-for="(message, index) in messages" :key="index">{{ message }}</li>
          </ul>
        </div>
      </div>

      <!-- Listeners Tab -->
      <div v-if="activeTab === 'listeners'" class="tab-pane">
        <ListenersTable/>
      </div>
    </div>
  </div>
</template>

<script setup>
import {ref, provide, onMounted} from 'vue';
import WebSocketConnection from './components/WebSocketConnection.vue';
import ListenersTable from './components/ListenersTable.vue';

// Track active tab
const activeTab = ref('status'); // Default to status tab

// WebSocket connection reference
const wsConnection = ref(null);

// Status information to display in the Status tab
const isConnected = ref(false);
const connectionStatus = ref('Disconnected');
const messages = ref([]);

// Create a shared state for listeners that will be accessible to both components
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

// Add message received from WebSocketConnection
function addMessage(message) {
  messages.value.push(message);
}

// Log for debugging
onMounted(() => {
  console.log('App mounted, WebSocketConnection ref:', wsConnection.value);
});
</script>

<style>
/* Add your styles here */
.status {
  margin: 20px 0;
  padding: 10px;
  background-color: #5e5e5e;
}

.status.connected {
  background-color: #4CAF50; /* Green color for connected state */
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