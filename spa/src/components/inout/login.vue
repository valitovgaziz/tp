<template>
    <div class="login-form">
      <h1>Логин</h1>
      <form @submit.prevent="handleSubmit">
        <div class="form-group">
          <label for="email">Email:</label>
          <input v-model.trim="email" type="email" id="email" required />
        </div>
        
        <div class="form-group">
          <label for="password">Пароль:</label>
          <input v-model.trim="password" type="password" id="password" required />
        </div>
        
        <button type="submit">Войти</button>
      </form>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        email: '',
        password: ''
      };
    },
    methods: {
      handleSubmit() {
        if (!this.isValid(this.email, this.password)) {
          alert("Пожалуйста, заполните все поля корректно.");
          return;
        }
  
        this.loginUser({
          email: this.email,
          password: this.password
        }).then(() => {
          this.$router.push('/'); // Переход на главную страницу
        }).catch((error) => {
          console.error(error);
          alert('Неверный email или пароль. Попробуйте снова.');
        });
      },
      isValid(email, password) {
        if (email.length === 0 || password.length === 0) {
          return false;
        }
  
        const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        if (!emailRegex.test(email)) {
          return false;
        }
  
        if (password.length < 6) {
          return false;
        }
  
        return true;
      },
      async loginUser(data) {
        console.log("Login by this data: ", data);
        // try {
        //   const response = await fetch('/login', {
        //     method: 'POST',
        //     headers: {
        //       'Content-Type': 'application/json'
        //     },
        //     body: JSON.stringify(data)
        //   });
  
        //   if (!response.ok) {
        //     throw new Error('Неверный email или пароль');
        //   }
        // } catch (error) {
        //   throw error;
        // }
      }
    }
  };
  </script>
  
  <style scoped>
  .login-form {
    max-width: 400px;
    margin: 50px auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 10px;
  }
  
  .form-group {
    margin-bottom: 15px;
  }
  
  .form-group label {
    display: block;
    margin-bottom: 5px;
  }
  
  .form-group input {
    width: 100%;
    padding: 10px;
    border: 1px solid #ddd;
    border-radius: 5px;
  }
  
  button {
    padding: 10px 20px;
    background-color: #4CAF50;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
  }
  </style>