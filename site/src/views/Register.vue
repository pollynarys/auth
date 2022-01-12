<template>
  <fieldset>
    <legend>Please register</legend>
    <form @submit.prevent="submit">
      <div class="authorize">
        <input v-model="data.name" placeholder="Name" required>
        <input v-model="data.email" type="email" placeholder="Email address" required>
        <input v-model="data.password" type="password" placeholder="Password" required>
      </div>
      <input type="submit" name="submit" value="submit">
    </form>
  </fieldset>
</template>

<script lang="ts">
import {reactive} from "vue";
import {useRouter} from "vue-router";

export default {
  name: "Register",
  setup() {
    const data = reactive({
      name: '',
      email: '',
      password: ''
    });
    const router = useRouter();

    const submit = async () => {
      await fetch('http://localhost:8000/api/register', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(data)
      });
      // console.log(data);
      await router.push('/login');
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