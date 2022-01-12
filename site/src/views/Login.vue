<template>
  <fieldset>
    <legend>Please sign in</legend>
    <form @submit.prevent="submit">
      <div class="authorize">
        <input v-model="data.email" type="email" placeholder="Email address" required>
        <input v-model="data.password" type="password" placeholder="Password" required>
      </div>
      <input type="submit" name="send" value="send">
    </form>
  </fieldset>
</template>

<script lang="ts">
import {reactive} from "vue";
import {useRouter} from "vue-router";

export default {
  name: "Login",
  setup() {
    const data = reactive({
      email: '',
      password: ''
    });

    const router = useRouter();

    const submit = async () => {
      await fetch('http://localhost:8000/api/login', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        // option for getting cookie
        credentials: 'include',
        body: JSON.stringify(data)
      });
      await router.push('/');
    }

    return {
      data,
      submit
    }
  }
}
</script>

<style scoped>

</style>