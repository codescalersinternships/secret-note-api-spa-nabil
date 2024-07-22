<template>
  <div>
    <div class="note" v-if="note">
      <p>{{ note.text }}</p>
      <p>Expires at: {{ note.expiredat }}</p>
      <p>Remaining visits: {{ note.noteremvisits }}</p>
      <p>Link to note {{ noteLink }}</p>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import { getNote } from "../Client";
import { BASE_URL } from "../config";

export default defineComponent({
  name: "NoteView",
  props: {
    noteId: {
      type: String,
      required: true,
    },
  },
  setup(props) {
    const note = ref<any>(null);
    const router = useRouter();
    var noteLink: string = router.resolve({
      name: "ViewNote",
      params: { id: props.noteId },
    }).href;
    noteLink = `${BASE_URL}${noteLink}`;

    onMounted(async () => {
      try {
        note.value = await getNote(props.noteId);
      } catch (error) {
        console.error("Failed to fetch note", error);
      }
    });

    return {
      note,
      noteLink,
    };
  },
});
</script>

<style scoped>
.note {
  padding: 1rem;
  border: 1px solid #ccc;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}
</style>
