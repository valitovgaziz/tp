<template>
    <div class="search-bar">
      <input
        type="text"
        placeholder="Поиск..."
        v-model="searchQuery"
        @keyup.enter="onSubmit"
      />
      <button @click="onSubmit">Найти</button>
    </div>
  </template>
  
  <script>
  export default {
    name: 'SearchBar',
    data() {
      return {
        searchQuery: ''
      };
    },
    methods: {
      onSubmit() {
        this.$emit('submited', this.searchQuery);

      }
    }
  };
  </script>
  
  <style scoped>
  .search-bar {
    display: flex;
    gap: 10px;
    align-items: center;
    margin-bottom: 20px;
  }
  
  .search-bar input {
    flex-grow: 1;
    padding: 10px;
    border: 1px solid #ccc;
    border-radius: 5px;
  }
  
  .search-bar button {
    padding: 10px 20px;
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
  }
  </style>