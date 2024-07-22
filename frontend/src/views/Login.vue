<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <div class="login-container">
    <NavBar />
    <h1>Login</h1>
    <LoginForm />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import LoginForm from "../components/LoginForm.vue";
import NavBar from "../components/NavBar.vue";

export default defineComponent({
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Login",
  components: {
    LoginForm,
    NavBar,
  },
});
</script>

<style scoped>
.login-container {
  max-width: 400px;
  margin: 0 auto;
  padding: 1rem;
  border: 1px solid #ccc;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}
</style>
