<template>
  <div class="signup-container">
    <NavBar />
    <h1>Sign Up</h1>
    <SignUpForm />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import SignUpForm from "../components/SignUpForm.vue";
import NavBar from "../components/NavBar.vue";

export default defineComponent({
  name: "SignUp",
  components: {
    SignUpForm,
    NavBar,
  },
});
</script>

<style scoped>
.signup-container {
  max-width: 400px;
  margin: 0 auto;
  padding: 1rem;
  border: 1px solid #ccc;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}
</style>
