<template>
    <div class="content-wrapper" :class="{ centered: isCentered }">
      <slot>
        <SearchBar />
      </slot>
    </div>
  </template>
  
  <script>
  import SearchBar from './SearchBar.vue';

  export default {
    name: 'ContentWrapper',
    props: {
      // Проп для управления положением контента (центрирование или выравнивание вверх)
      isCentered: {
        type: Boolean,
        default: true
      }
    },
    components: {
        SearchBar
    }
  };
  </script>

  
  <style scoped>
  .content-wrapper {
    display: flex;
    flex-direction: column;
    justify-content: flex-start; /* Выравниваем контент к верху */
    align-items: stretch;
    min-height: 80vh; /* Если нужно растянуть контейнер на всю высоту экрана */
  }
  
  .centered {
    justify-content: center; /* Центрирование контента вертикально */
    align-items: center;     /* Центрирование контента горизонтально */
  }
  </style>