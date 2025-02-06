<template>
  <Header />
  <Profile />
  <Footer />
</template>


<style></style>

<script>
import Header from '../components/header/header.vue';
import Footer from '../components/footer.vue';
import Profile from '../components/profile.vue';

export default {
  components: {
    Header,
    Footer,
    Profile,
  },
};
</script>