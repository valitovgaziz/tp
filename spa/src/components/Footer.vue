<template>
    <footer class="footer">
      <nav class="footer-nav">
        <ul>
          <li><a href="#about">О нас</a></li>
          <li><a href="#requisites">Реквизиты</a></li>
          <li><a href="#vacancies">Вакансии</a></li>
          <li><a href="#donates">Донаты</a></li>
        </ul>
      </nav>
      <p class="copyright">© Все права защищены, 2023</p>
    </footer>
  </template>
  
  <script>
  export default {
    name: 'AppFooter'
  };
  </script>
  
  <style scoped>
  .footer {
    background-color: #f9f9f9;
    padding: 20px;
    text-align: center;
  }
  
  .footer-nav ul {
    list-style-type: none;
    padding: 0;
    margin: 0;
    display: flex;
    justify-content: center;
  }
  
  .footer-nav li {
    margin-right: 20px;
  }
  
  .footer-nav li:last-child {
    margin-right: 0;
  }
  
  .footer-nav a {
    text-decoration: none;
    color: #333;
  }
  
  .copyright {
    font-size: 14px;
    margin-top: 10px;
  }
  </style>