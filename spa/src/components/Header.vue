<template>
    <header class="header">
      <!-- Логотип -->
      <div class="logo">
        <img src="@/assets/logo.png" alt="Логотип" />
      </div>
      
      <!-- Выпадающее меню -->
      <div class="dropdown-menu">
        <button @click="isOpen = !isOpen">Меню</button>
        
        <ul v-show="isOpen" class="menu-list">
          <li v-for="item in menuItems" :key="item.id">
            <router-link :to="item.to">{{ item.id }} - {{ item.title }}</router-link>
          </li>
        </ul>
      </div>
    </header>
  </template>
  
  <script>
  export default {
    name: 'AppHeader',
    data() {
      return {
        isOpen: false,
        menuItems: [
          { id: 1, title: 'Главная', to: '#' },
          { id: 2, title: 'О нас', to: '/about' },
          { id: 3, title: 'Услуги', to: '#' },
          { id: 4, title: 'Контакты', to: '#' },
          { id: 5, title: 'Блог', to: '#' },
          { id: 6, title: 'SignUp', to: '#' }
        ]
      };
    }
  };
  </script>
  
  <style scoped>
  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 10px;
  }
  
  .logo img {
    width: 100px;
    height: auto;
  }
  
  .dropdown-menu button {
    background-color: #f0f0f0;
    border: none;
    cursor: pointer;
    padding: 8px 16px;
  }
  
  .menu-list {
    list-style-type: none;
    margin: 0;
    padding: 0;
    position: absolute;
    right: 0;
    background-color: white;
    box-shadow: 0 2px 6px rgba(0, 0, 0, 0.15);
    z-index: 99;
  }
  
  .menu-list li a {
    display: block;
    text-decoration: none;
    color: black;
    padding: 12px 20px;
  }
  
  .menu-list li a:hover {
    background-color: #eaeaea;
  }
  </style>