<template>
  <fieldset>
    <legend>Please sign in</legend>
    <form @submit.prevent="submit">
      <div class="auth">
        <label for="inputEmail"></label>
        <input v-model="data.email" type="email" id="inputEmail" size="30" placeholder="Email address" required>
        <label for="inputPassword"></label>
        <input v-model="data.password" type="password" id="inputPassword" size="30" placeholder="Password" required>
      </div>
      <div class="checkbox">
        <label><input type="checkbox" value="Remember-me">Remember me</label>
        <input type="submit" name="send" value="send">
      </div>
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