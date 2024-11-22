<template>
    <div class="search-result-card">
      <img :src="result.photo" alt="Фото" class="card-image" />
      <div class="card-content">
        <h3 class="title">{{ result.title }}</h3>
        <p class="description">{{ result.description }}</p>
        <router-link :to="`/details/${result.id}`" class="view-more">Подробнее</router-link>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: "SearchResultCard",
    props: {
      result: {
        type: Object,
        required: true,
      },
    },
  };
  </script>
  
  <style scoped>
  .search-result-card {
    border: 1px solid #ddd;
    border-radius: 5px;
    overflow: hidden;
    background-color: #fff;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    transition: transform 0.3s ease-in-out;
  }
  
  .card-image {
    width: 100%;
    height: 200px;
    object-fit: cover;
  }
  
  .card-content {
    padding: 10px;
  }
  
  .title {
    margin-top: 0;
    margin-bottom: 5px;
    font-weight: bold;
  }
  
  .view-more {
    text-decoration: none;
    color: #007bff;
    align-self: flex-end;
  }
  
  .search-result-card:hover {
    transform: scale(1.02);
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  }
  </style>