<!-- eslint-disable vue/require-v-for-key -->
<template>
  <div>
    <NavBar />
    <h1>Home</h1>
    <p>Welcome to the Secret Note App!</p>
    <div>
      <ul>
        <div v-if="notes">
          <li v-for="note in notes">
            <p>{{ note.text }}</p>
            <p>Expires at: {{ note.expiredat }}</p>
            <p>Remaining visits: {{ note.noteremvisits }}</p>
            <p>
              Link to note
              {{
                `${BASE_URL}${
                  router.resolve({
                    name: "ViewNote",
                    params: { id: note.noteId },
                  }).href
                }`
              }}
            </p>
          </li>
        </div>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from "vue";
import NavBar from "../components/NavBar.vue";
import { useRouter } from "vue-router";
import { getUserNotes } from "../Client";
import { BASE_URL } from "../config";

export default defineComponent({
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Home",
  components: {
    NavBar,
  },
  setup() {
    const notes = ref<any>(null);
    const router = useRouter();
    onMounted(async () => {
      try {
        notes.value = await getUserNotes();
      } catch (error) {
        console.error("Failed to fetch note", error);
      }
    });

    return {
      notes,
      BASE_URL,
      router,
    };
  },
});
</script>

<style scoped>
h1 {
  margin-top: 20px;
  text-align: center;
}
</style>
