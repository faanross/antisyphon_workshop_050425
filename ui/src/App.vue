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
</style>