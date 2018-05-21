<template>
  <div id="login">
    <el-form ref="form" :model="form" label-width="200px">
      <el-input v-model="form.email" placeholder="Email"></el-input>
      <el-input type="password" v-model="form.password" placeholder="Password" maxlength=20></el-input>
    </el-form>
  </div>
</template>

<script>
export default {
  name: 'login',
  data() {
    return {
      form: {
        email: '',
        password: '',
      }
    }
  },
};
</script>

<style>
</style>
