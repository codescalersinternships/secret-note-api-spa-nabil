<template>
    <div>
        <NavBar />
        <h1>Home</h1>
        <p>Welcome to the Secret Note App!</p>
    </div>
</template>
  
<script lang="ts">
    import { defineComponent } from 'vue';
    import NavBar from '../components/NavBar.vue';

    export default defineComponent({
        name: 'Home',
        components: {
            NavBar,
        },
    });
</script>
  
<style scoped>
  h1 {
    margin-top: 20px;
    text-align: center;
  }
</style>
  