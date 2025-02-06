<template>
    <div class="reviews-container">
      <h2>Отзывы о приложении</h2>
  
      <!-- Форма для добавления отзыва -->
      <div class="add-review">
        <textarea v-model="newReview" placeholder="Напишите ваш отзыв..." rows="4"></textarea>
        <button @click="submitReview" :disabled="!newReview.trim()">Оставить отзыв</button>
      </div>
  
      <!-- Список последних отзывов -->
      <div class="reviews-list">
        <div v-for="review in reviews" :key="review.id" class="review-item">
          <p>{{ review.text }}</p>
          <small>{{ formatDate(review.created_at) }}</small>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        reviews: [], // Список отзывов
        newReview: '', // Текст нового отзыва
      };
    },
    mounted() {
      this.loadReviews(); // Загружаем отзывы при монтировании компонента
    },
    methods: {
      // Загрузка отзывов из базы данных
      async loadReviews() {
        try {
          const response = await axios.get('http://213.108.4.63:3000/api/reviews');
          this.reviews = response.data.slice(0, 5); // Показываем только последние 5 отзывов
        } catch (error) {
          console.error('Ошибка при загрузке отзывов:', error);
        }
      },
      // Отправка нового отзыва
      async submitReview() {
        if (!this.newReview.trim()) return;
  
        try {
          const response = await axios.post('http://213.108.4.63:3000/api/reviews', {
            text: this.newReview,
          });
          this.reviews.unshift(response.data); // Добавляем новый отзыв в начало списка
          this.newReview = ''; // Очищаем текстовое поле
          if (this.reviews.length > 5) {
            this.reviews.pop(); // Удаляем самый старый отзыв, если их больше 5
          }
        } catch (error) {
          console.error('Ошибка при отправке отзыва:', error);
        }
      },
      // Форматирование даты
      formatDate(dateString) {
        const date = new Date(dateString);
        return date.toLocaleString(); // Форматируем дату в удобный формат
      },
    },
  };
  </script>
  
  <style scoped>
  .reviews-container {
    max-width: 600px;
    margin: 0 auto;
    padding: 20px;
  }
  
  .add-review textarea {
    width: 100%;
    padding: 10px;
    margin-bottom: 10px;
    border-radius: 5px;
    border: 1px solid #ccc;
  }
  
  .add-review button {
    padding: 10px 20px;
    background-color: #42b983;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
  }
  
  .add-review button:disabled {
    background-color: #ccc;
    cursor: not-allowed;
  }
  
  .reviews-list {
    margin-top: 20px;
  }
  
  .review-item {
    background: #f9f9f9;
    padding: 15px;
    border-radius: 5px;
    margin-bottom: 10px;
  }
  
  .review-item p {
    margin: 0;
  }
  
  .review-item small {
    color: #666;
    font-size: 0.9em;
  }
  </style>