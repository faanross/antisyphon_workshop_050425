<!-- ListenersTable.vue -->
<template>
  <!-- Listener Counter -->
  <div class="counter">
    Total Listeners: <span>{{ displayedListeners.length }}</span>
  </div>

  <!-- Listeners Table -->
  <table>
    <thead>
    <tr>
      <th>Start</th>
      <th>Duration</th>
      <th>ID</th>
      <th>Port</th>
    </tr>
    </thead>
    <tbody>
    <!-- Show a message when there are no listeners -->
    <tr v-if="displayedListeners.length === 0">
      <td colspan="4" style="text-align: center;">No listeners available</td>
    </tr>
    <tr v-for="listener in displayedListeners" :key="listener.id">
      <td>{{ formatDate(listener.createdAt) }}</td>
      <td>{{ calculateDuration(listener.createdAt) }}</td>
      <td>{{ listener.id }}</td>
      <td>{{ listener.port }}</td>
    </tr>
    </tbody>
  </table>
</template>

<script setup>
import {ref, inject, computed, watch, onMounted} from 'vue';

// Get the shared listeners state
const listenersState = inject('listenersState');

// Create a local copy of the listeners data for display
const localListeners = ref([]);

// Define a computed property that will always have the latest data
const displayedListeners = computed(() => {
  return localListeners.value;
});

// Watch for changes in the shared listeners state
watch(() => listenersState.getListeners(), (newListeners) => {
  console.log('ListenersTable - listeners state changed:', newListeners);
  // Update our local copy with the new data
  localListeners.value = [...newListeners];
}, {immediate: true, deep: true});

// On component mount, initialize with the current listener data
onMounted(() => {
  console.log('ListenersTable mounted, fetching listeners from shared state');
  localListeners.value = [...listenersState.getListeners()];
  console.log('Local listeners initialized:', localListeners.value);
});

// Format date function
const formatDate = (dateString) => {
  if (!dateString) return 'N/A';
  try {
    const date = new Date(dateString);
    return date.toLocaleString();
  } catch (e) {
    console.error('Error formatting date:', e);
    return 'Invalid date';
  }
};

// Calculate duration from creation time
const calculateDuration = (dateString) => {
  if (!dateString) return 'N/A';
  try {
    const startTime = new Date(dateString);
    const now = new Date();
    const diff = now - startTime;

    // Convert to seconds
    const seconds = Math.floor(diff / 1000);

    // Format as hours, minutes, seconds
    const hours = Math.floor(seconds / 3600);
    const minutes = Math.floor((seconds % 3600) / 60);
    const remainingSeconds = seconds % 60;

    return `${hours.toString().padStart(2, '0')}:${minutes.toString().padStart(2, '0')}:${remainingSeconds.toString().padStart(2, '0')}`;
  } catch (e) {
    console.error('Error calculating duration:', e);
    return 'Invalid date';
  }
};
</script>