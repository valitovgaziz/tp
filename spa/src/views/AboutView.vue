<template>
  <Header />
  <About />
  <Developers />
  <Commits />
  <Footer />
</template>

<style></style>



<script>
import Header from '../components/header/header.vue';
import Footer from '../components/footer.vue';
import About from '../components/about/about.vue';
import Commits from '../components/about/commits.vue';
import Developers from '../components/about/developers.vue';

export default {
  components: {
    Header,
    Footer,
    About,
    Developers,
    Commits,
  },
};
</script>