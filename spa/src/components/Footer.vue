<template>
    <div class="footer">
        <div class="navigation" v-on:click="goBack">
            <img src="./images/icons/arrow_back_24dp_E8EAED_FILL0_wght400_GRAD0_opsz24.svg" alt="back">
        </div>
        <router-link to="/">
            <img src="./images/icons/home_24dp_E8EAED_FILL0_wght400_GRAD0_opsz24.svg" alt="home">
        </router-link>
        <div class="navigation" v-on:click="goForward">
            <img src="./images/icons/arrow_forward_24dp_E8EAED_FILL0_wght400_GRAD0_opsz24.svg" alt="forward">
        </div>
    </div>
</template>

<style scoped>
.footer {
    height: 3rem;
    background-color: darkmagenta;
    border-radius: 1rem;
    display: flex;
    flex-direction: row;
    align-items: center;
    padding: 0 2rem;
    margin: 0 0 0.5rem 0;
}
.navigation {
    padding: 0 1rem;
    margin: 0 1rem;
}
.navigation:hover {
    cursor: pointer;
    background-color: rgb(112, 48, 171);
}
</style>

<script>
import MenuIcon from 'vue-material-design-icons/Menu.vue';
export default {
    components: {
        MenuIcon,
    },
    name: 'footer',
    methods: {
        goBack() {
            this.$router.go(-1);
        },
        goForward() {
            this.$router.go(1);
        }
    }
};

</script>