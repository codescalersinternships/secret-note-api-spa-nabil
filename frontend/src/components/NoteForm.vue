<template>
    <form @submit.prevent="submitNote">
      <div class="form-group">
        <textarea v-model="text" placeholder="Note text" required></textarea>
      </div>
      <div class="form-group">
        <input v-model="expireDate" type="datetime-local" placeholder="Expiration date" required />
      </div>
      <div class="form-group">
        <input v-model="maxRemDays" type="number" placeholder="Max remaining days" required />
      </div>
      <button type="submit">Create Note</button>
    </form>
  </template>
  
  <script lang="ts">
  import { defineComponent, ref } from 'vue';
  import { useRouter } from 'vue-router';
  import { createNote } from '../Client.ts'; 
  
  export default defineComponent({
    name: 'NoteForm',
    setup() {
      const text = ref('');
      const expireDate = ref('');
      const maxRemDays = ref(0);
      const router = useRouter();
      
      const submitNote = async () => {
        try {
          expireDate.value = expireDate.value.replace("T"," ");
          expireDate.value = `${expireDate.value}:00`
          await createNote({
            text: text.value,
            noteremvisits: maxRemDays.value,
            expiredat: expireDate.value.replace("T"," "),
            userid: null
          });
          router.push('/');
        } catch (error) {
          console.error('Failed to create note', error);
        }
      };
  
      return {
        text,
        expireDate,
        maxRemDays,
        submitNote,
      };
    },
  });
  </script>
  
<style scoped>
  .form-group {
    margin-bottom: 1rem;
  }
  
  textarea {
    width: 100%;
    height: 100px;
    padding: 0.5rem;
  }
  
  input, button {
    width: 100%;
    padding: 0.5rem;
    margin-top: 0.5rem;
    border-radius: 4px;
  }
  
  button {
    background-color: #007bff;
    color: white;
    cursor: pointer;
    border: none;
  }
  
  button:hover {
    background-color: #0056b3;
  }
</style>
  