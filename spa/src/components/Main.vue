<template>
    <div class="content-wrapper" :class="{ centered: isCentered }">
      <slot>
        <SearchBar @submited="queryStr" />
        <SearchResults ref="sdf" />
      </slot>
    </div>
  </template>
  
  <script>
  import SearchBar from './SearchBar.vue';
  import SearchResults from './SearchResults.vue';

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
        SearchBar,
        SearchResults
    },
    methods: {
        queryStr(searchQuery) {
            alert(this.$refs.sdf)
            this.isCentered = false;
            this.$refs.searchResultsRef.showContent = true;
            this.$refs.searchResultsRef.results = [
                {
                    "id": "asldfjk",
                    "title": "title1",
                    "description": "long description about this result"
                },
                {
                    "id": "asladsdfjk",
                    "title": "title2",
                    "description": "long description about this result" 
                },{
                    "id": "asldfasdfjk",
                    "title": "title3",
                    "description": "long description about this result"
                }
            ];
        }
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