<template>
    <div class="search-container">
      <div class="search-input-container">
        <!-- Иконка для вызова меню параметров -->
        <div class="menu-icon" @click="toggleMenu">
          <svg viewBox="0 0 24 24" width="24" height="24">
            <path d="M3 18h18v-2H3v2zm0-5h18v-2H3v2zm0-7v2h18V6H3z" fill="currentColor"/>
          </svg>
        </div>
        <!-- Поле ввода -->
        <input
          v-model="searchQuery"
          @keyup.enter="performSearch"
          class="search-input"
          placeholder="Введите место для поиска"
        />
        <!-- Кнопка поиска -->
        <button @click="performSearch" class="search-button">
          <svg viewBox="0 0 24 24" width="24" height="24">
            <path d="M10 18a8 8 0 1 1 0-16 8 8 0 0 1 0 16zm8-8a7.95 7.95 0 0 1-1.757 4.95l4.95 4.95-1.414 1.414-4.95-4.95A7.95 7.95 0 0 1 10 18z" fill="currentColor"/>
          </svg>
        </button>
      </div>
      <!-- Выпадающее меню для параметров поиска -->
      <div v-if="isMenuVisible" class="search-menu">
        <div class="menu-item" v-for="option in searchOptions" :key="option" @click="selectOption(option)">
          {{ option }}
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        searchQuery: '',
        isMenuVisible: false,
        searchOptions: ['Пляжи', 'Горы', 'Отели', 'Рестораны']
      };
    },
    methods: {
      performSearch() {
      if (this.searchQuery.trim()) {
        // Переход на страницу результатов поиска с передачей запроса
        this.$router.push({
          name: 'results', // Имя маршрута для страницы результатов
          query: { q: this.searchQuery } // Передача поискового запроса через query-параметры
        });
      }
      },
      toggleMenu() {
        this.isMenuVisible = !this.isMenuVisible;
      },
      selectOption(option) {
        this.searchQuery = option;
        this.isMenuVisible = false;
        this.performSearch();
      }
    }
  };
  </script>
  
  <style scoped>
  .search-container {
    position: relative;
    width: 100%;
    max-width: 600px;
    margin: auto;
  }
  
  .search-input-container {
    display: flex;
    align-items: center;
    background: #fff;
    border-radius: 25px;
    padding: 10px;
    box-shadow: 0 2px 5px rgba(0,0,0,0.2);
  }
  
  .menu-icon {
    cursor: pointer;
    margin-right: 10px;
  }
  
  .search-input {
    flex-grow: 1;
    border: none;
    outline: none;
    padding: 10px;
    border-radius: 20px;
  }
  
  .search-button {
    background: none;
    border: none;
    cursor: pointer;
    padding: 10px;
  }
  
  .search-menu {
    position: absolute;
    top: 60px;
    left: 0;
    right: 0;
    background: #fff;
    border-radius: 10px;
    box-shadow: 0 4px 6px rgba(0,0,0,0.1);
    z-index: 1000;
  }
  
  .menu-item {
    padding: 10px;
    cursor: pointer;
  }
  
  .menu-item:hover {
    background-color: #f0f0f0;
  }
  </style>