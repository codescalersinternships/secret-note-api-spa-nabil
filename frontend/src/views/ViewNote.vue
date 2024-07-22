<template>
  <div class="view-note-container">
    <NavBar />
    <h1>View Note</h1>
    <NoteView :noteId="noteId" />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import NoteView from "../components/NoteView.vue";
import NavBar from "../components/NavBar.vue";
import { useRoute } from "vue-router";

export default defineComponent({
  name: "ViewNote",
  components: {
    NoteView,
    NavBar,
  },
  setup() {
    const route = useRoute();
    const noteId = route.params.id as string;
    return {
      noteId,
    };
  },
});
</script>

<style scoped>
.view-note-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 1rem;
  border: 1px solid #ccc;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}
</style>
