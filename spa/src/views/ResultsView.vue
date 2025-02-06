<template>
    <Header />
    <SearchLine />
    <Results />
    <Footer />
</template>

<style></style>



<script>
import Header from '../components/header/header.vue';
import Footer from '../components/footer.vue';
import Results from '../components/saerch_results/results.vue';
import searchLine from '@/components/searchLine/searchLine.vue';
import SearchLine from '@/components/searchLine/searchLine.vue';

export default {
    components: {
        Header,
        Footer,
        Results,
        SearchLine,
    },
};
</script>