<template>
  <form @submit.prevent="loginForm">
    <div class="form-group">
      <input v-model="email" type="email" placeholder="Email" required />
    </div>
    <div class="form-group">
      <input v-model="password" type="password" placeholder="Password" required />
    </div>
    <button type="submit">Login</button>
  </form>
</template>
  
<script lang="ts">
    import { defineComponent, ref } from 'vue';
    import { useRouter } from 'vue-router';
    import { login } from '../Client';

    export default defineComponent({
      name: 'LoginForm',
      setup() {
          const email = ref('');
          const password = ref('');
          const router = useRouter();

          const loginForm = async () => {
          try {
              await login({email:email.value, password: password.value});
              router.push('/');
          } catch (error) {
              console.error('Failed to login', error);
          }
          };

          return {
            email,
            password,
            loginForm,
          };
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