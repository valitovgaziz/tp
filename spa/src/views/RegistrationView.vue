<template>
    <Header />
    <Registration />
    <Footer />
</template>

<style></style>



<script>
import Header from '../components/header/header.vue';
import Footer from '../components/footer.vue';
import Registration from '../components/inout/registration.vue';

export default {
    components: {
        Header,
        Footer,
        Registration,
    },
};
</script>