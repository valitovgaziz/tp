<template>
    <Header />
    <LogIn />
    <Footer />
</template>

<style></style>



<script>
import Header from '../components/header/header.vue';
import Footer from '../components/footer.vue';
import LogIn from '../components/inout/login.vue';

export default {
    components: {
        Header,
        Footer,
        LogIn,
    },
};
</script>