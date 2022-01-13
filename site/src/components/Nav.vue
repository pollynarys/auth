<template>
  <nav>
    <ul class="nav">
      <li>
        <router-link to="/">MAIN</router-link>
      </li>
      <!--      <ul v-if="auth">-->
      <li>
        <router-link to="/login">LOGIN</router-link>
      </li>
      <li>
        <router-link to="/register">REGISTER</router-link>
      </li>
      <!--      </ul>-->
      <!--      <ul v-if="!auth">-->
      <li>
        <a href="#" @click="logout">LOGOUT</a>
      </li>
      <!--      </ul>-->
    </ul>
  </nav>
</template>

<script lang="ts">
import {useStore} from "vuex";
import {computed} from "vue";

export default {
  name: "Nav",
  setup() {
    const store = useStore();
    const auth = computed(() => store.state.authenticated);

    const logout = async () => {
      await fetch('http://localhost:8000/api/logout', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        // option for getting cookie
        credentials: 'include',
      });
    }
    return {
      auth,
      logout
    }
  }
}
</script>

<style scoped>

</style>