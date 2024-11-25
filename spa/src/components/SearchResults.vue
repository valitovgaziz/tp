<template>
  <div  v-show="showContent" class="search-results">
    <h2>Результаты поиска</h2>
    <div class="results-container">
      <SearchResultCard v-for="result in results" :key="result.id" :result="result" />
    </div>
  </div>
</template>

<script>
import SearchResultCard from "./SearchResultCard.vue";

export default {
  data() {
    return {
      showContent: false,
    };
  },
  name: "SearchResults",
  components: {
    SearchResultCard,
  },
  props: {
    results: {
      type: Array,
      required: true,
    },
  },
};

</script>

<style scoped>
.search-results {
  max-width:src="result.photo" 800px;
  margin: 0 auto;
  padding: 20px;
}

.results-container {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 20px;
}
</style>