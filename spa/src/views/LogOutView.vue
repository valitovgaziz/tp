<template>
    <Header />
    <LogOut />
    <Footer />
</template>

<style></style>



<script>
import Header from '../components/header/header.vue';
import Footer from '../components/footer.vue';
import LogOut from '../components/inout/logout.vue';

export default {
    components: {
        Header,
        Footer,
        LogOut,
    },
};
</script>