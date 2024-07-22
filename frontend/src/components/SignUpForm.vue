<template>
  <form @submit.prevent="signUpForm">
    <div class="form-group">
      <input v-model="name" type="text" placeholder="Name" required />
    </div>
    <div class="form-group">
      <input v-model="email" type="email" placeholder="Email" required />
    </div>
    <div class="form-group">
      <input
        v-model="password"
        type="password"
        placeholder="Password"
        required
      />
    </div>
    <button type="submit">Sign Up</button>
  </form>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { useRouter } from "vue-router";
import { signUp } from "../Client";

export default defineComponent({
  name: "SignUpForm",
  setup() {
    const name = ref("");
    const email = ref("");
    const password = ref("");
    const router = useRouter();

    const signUpForm = async () => {
      try {
        await signUp({
          name: name.value,
          email: email.value,
          password: password.value,
        });
        router.push("/login");
      } catch (error) {
        console.error("Failed to login", error);
      }
    };

    return {
      name,
      email,
      password,
      signUpForm,
    };
  },
});
</script>

<style scoped>
.form-group {
  margin-bottom: 1rem;
}

button {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 4px;
  background-color: #007bff;
  color: white;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}
</style>
