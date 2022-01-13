<template>
  <header>
    <Nav/>
  </header>
  <br>
  <div class="auth-container">
    <h1>YOUR NOTE SYSTEM</h1>
    <router-view/>
  </div>
</template>

<script>
import Nav from "@/components/Nav";

export default {
  components: {
    Nav
  }
}
</script>

<style>
body {
  font-family: Helvetica, Arial, sans-serif;
  margin: 0;
  padding: 0;
  background-color: #fafafa;
  font-size: 20px;
}

header nav {
  position: absolute;
  top: 0;
  right: 0;
}

h1 {
  color: #3F51B5;
  clear: both;
}

li {
  padding: 0.5em;
  margin: 0;
  display: inline-block;
}

a, router-link {
  padding: 0.5em;
  display: block;
}

a {
  color: #282828;
  text-decoration: none;
}

a:hover {
  color: #dc6454;
}

li {
  list-style-type: none;
}

.auth-container {
  padding: 3em;
  margin: auto auto;
  text-align: center;
  width: 30%;
}

fieldset {
  border-radius: 10px;
  border: 0;
  text-align: left;
  box-shadow: 0 0 15px 4px rgba(0, 0, 0, 0.06);
}

input {
  margin-bottom: 1em;
}

.authorize input {
  text-align: center;
  height: 2em;
  width: 100%;
  border: 0;
  border-radius: 10px;
  box-shadow: 0 0 15px 4px rgba(0, 0, 0, 0.06);
}

.authorize input:hover {
  background-color: #ccc;
}

.authorize input:focus {
  border: 0;
  background-color: #fff;
  font-weight: 600;
}

input[type="submit"] {
  height: 2em;
  border: none;
  width: 100%;
  background-color: #3F51B5;
  color: #fff;
  border-radius: 10px;
}

input[type="submit"]:active {
  font-weight: 600;
}

</style>
