<template>
  {{ message }}
</template>

<script lang="ts">
import {onMounted, ref} from "vue";
import {useStore} from "vuex";

export default {
  name: "Home",
  setup() {
    const message = ref('You are on logged in');
    const store = useStore();

    onMounted(async () => {
      try {
        const response = await fetch('http://localhost:8000/api/user', {
          headers: {'Content-Type': 'application/json'},
          credentials: 'include',
        });
        const content = await response.json();
        message.value = `Welcome ${content.name}`;

        await store.dispatch('setAuth', true);
      } catch (e) {
        await store.dispatch('setAuth', false);
      }
    });

    return {
      message
    }
  }
}
</script>

<style scoped>

</style>