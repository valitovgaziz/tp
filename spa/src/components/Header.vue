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
            <router-link :to="item.tooo">{{ item.id }} - {{ item.title }}</router-link>
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
          { id: 1, title: 'Главная', tooo: 'home' },
          { id: 2, title: 'О нас', tooo: 'about' },
          { id: 3, title: 'Услуги', tooo: 'home' },
          { id: 4, title: 'Контакты', tooo: 'home' },
          { id: 5, title: 'Блог', tooo: 'home' },
          { id: 6, title: 'SignUp', tooo: 'signup' }
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