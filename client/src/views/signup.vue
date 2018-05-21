<template>
  <div id="signup">
    <el-form ref="form" :model="form" label-width="200px">
      <el-input v-model="form.username" placeholder="Username" data-rules="'required|email'" name="email">
        <template slot="append">@herd.something.com</template>
      </el-input>
      <el-input v-model="form.email" placeholder="Email"></el-input>
      <el-input type="password" v-model="form.password" placeholder="Password" maxlength=20></el-input>
      <el-input type="password" v-model="form.email" placeholder="Confirm password"></el-input>
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
        username: '',
        password: '',
      }
    }
  },
};
</script>

<style>
</style>
