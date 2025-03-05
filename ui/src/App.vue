<template>
  <div>
    <h1>OrlokC2</h1>

    <!-- Listener Counter -->
    <div class="counter">
      Total Listeners: <span>0</span>
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
      <!-- Empty table body as requested -->
      </tbody>
    </table>
  </div>
</template>

<script setup>
// Empty setup - we'll add functionality later
</script>

<style>
/* Minimal styling */
table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 20px;
}

th, td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}

th {
  background-color: #5e5e5e;
}

.counter {
  margin: 20px 0;
  padding: 10px;
  background-color: #5e5e5e;
}
</style>