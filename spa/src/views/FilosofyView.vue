<template>
    <Header />
    <Filosofy />
    <Footer />
</template>

<style>

</style>

<script>
import Header from '../components/header/header.vue';
import Footer from '../components/footer.vue';
import Filosofy from '../components/filosofy.vue';

export default {
    components: {
        Footer,
        Header,
        Filosofy,
    },
}
</script>