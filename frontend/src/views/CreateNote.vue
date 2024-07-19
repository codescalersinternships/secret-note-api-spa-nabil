<template>
    <div class="create-note-container">
      <NavBar />
      <h1>Create Note</h1>
      <NoteForm />
    </div>
</template>
  
<script lang="ts">
  import { defineComponent } from 'vue';
  import NoteForm from '../components/NoteForm.vue';
  import NavBar from '../components/NavBar.vue';
  
  export default defineComponent({
    name: 'CreateNote',
    components: {
      NoteForm,
      NavBar,
    },
  });
</script>
  
<style scoped>
  .create-note-container {
    max-width: 600px;
    margin: 0 auto;
    padding: 1rem;
    border: 1px solid #ccc;
    border-radius: 8px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  }
</style>
  