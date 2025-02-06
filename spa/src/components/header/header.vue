<template>
    <div class="header">
        <router-link to="/about" id="rl-about"><img src="../images/logo150x150.png" alt="logo"></router-link>
        <router-link to="/filosofy" id="rl-filosofy">Trip*Freedom*Adventure</router-link>
        <Menu />
    </div>
</template>


<style scoped>
.header {
    width: 100%;
    height: 100%;
    margin: 0;
    padding: 0;
    border-radius: 1rem;
    background-color: aquamarine;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: space-between;
    height: 3rem;
}

#rl-about {
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-items: center;
    border-radius: 50%;
    padding: 1rem;  
    height: 1rem;
    width: 1rem;
}

#rl-about:hover {
    cursor: pointer;
    background-color: rgb(85, 117, 117);
}

img {
    height: 1.5rem;
    padding: 0 1rem 0 0;
}

#rl-filosofy {
    all: unset;
    box-shadow: 1px 2px 3px rgb(63, 145, 130);
    border-radius: 1rem;
    padding: 0.5rem;
}

#rl-filosofy:hover {
    cursor: pointer;
    background-color: rgb(143, 202, 202);
    box-shadow: 2px 3px 4px rgb(64,146, 132);
}
</style>


<script>
import Menu from './menu.vue';

export default {
    components: {
        Menu,
    },
    name: 'header',
}
</script>