<template>
    <div class="registration-form">
      <h2>Регистрация</h2>
      <form @submit.prevent="handleSubmit">
        <label for="name">Имя:</label>
        <input type="text" id="name" v-model="formData.name" required />
  
        <label for="phone">Телефон:</label>
        <input type="tel" id="phone" v-model="formData.phone" required />
  
        <label for="password">Пароль:</label>
        <input type="password" id="password" v-model="formData.password" required />
  
        <button type="submit">Зарегистрироваться</button>
      </form>
    </div>
  </template>
  
  <script setup>
  import { reactive } from 'vue';
  import axios from 'axios';
  
  const formData = reactive({
    name: '',
    phone: '',
    password: ''
  });
  
  async function handleSubmit() {
    try {
      const response = await axios.post('http://localhost:8080/api/signup', formData);
      console.log(response.data); // Логируем ответ сервера
      alert('Успешная регистрация!');
    } catch (error) {
      console.error(error);
      alert('Ошибка при регистрации');
    }
  }
  </script>
  
  <style scoped>
  .registration-form {
    max-width: 400px;
    margin: 0 auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 10px;
  }
  
  label {
    display: block;
    margin-bottom: 5px;
  }
  
  input {
    width: 100%;
    padding: 8px;
    margin-bottom: 15px;
    border: 1px solid #ddd;
    border-radius: 4px;
  }
  
  button {
    background-color: #4CAF50;
    color: white;
    padding: 12px 20px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  </style>