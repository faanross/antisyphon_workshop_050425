<template>
  <div class="container">
    <h1>OrlokC2</h1>

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
      <!-- Status Tab -->
      <div v-if="activeTab === 'status'" class="tab-pane">
        <WebSocketConnection/>
      </div>

      <!-- Listeners Tab -->
      <div v-if="activeTab === 'listeners'" class="tab-pane">
        <ListenersTable/>
      </div>
    </div>
  </div>
</template>

<script setup>
import {ref} from 'vue';
import WebSocketConnection from './components/WebSocketConnection.vue';
import ListenersTable from './components/ListenersTable.vue';

// Track active tab
const activeTab = ref('status'); // Default to status tab
</script>

<style>
.container {
  width: 100%;
  max-width: 900px;
  margin: 0 auto;
}

.tabs-container {
  position: relative;
  width: 100%;
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-bottom: 16px;
}

.tabs {
  display: flex;
  justify-content: center;
  width: 100%;
  z-index: 1;
}

.tab-line {
  position: absolute;
  bottom: 0;
  width: 100%;
  height: 1px;
  background-color: #ccc;
  z-index: 0;
}

.tabs button {
  padding: 8px 40px;
  border: none;
  background: none;
  cursor: pointer;
  font-size: 16px;
  border-radius: 4px 4px 0 0;
  margin: 0 10px;
  position: relative;
  color: #fffbfb;
}

.tabs button.active {
  background-color: #8a57e9;
  border: 1px solid #8a57e9;
  border-bottom: none;
  box-shadow: 0 0 10px rgba(138, 87, 233, 0.5);
}

.tab-content {
  padding: 16px;
  width: 100%;
  min-height: 400px;
  box-sizing: border-box;
  display: flex;
  justify-content: center;
}

.tab-pane {
  animation: fadeIn 0.3s;
  width: 100%;
  max-width: 800px;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
</style>