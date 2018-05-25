<template>
  <div id="signup">
    <h1>{{ $t("views.signup.title") }}</h1>
    <p>{{ $t("views.signup.message") }}</p>
    <el-form ref="form" :model="form" :rules="rules">
      <el-form-item prop="username">
        <el-input v-model="form.username" :placeholder="$t('form.username.name')">
          <template slot="append">@herd.something.com</template>
        </el-input>
      </el-form-item>
      <el-form-item prop="email">
        <el-input v-model="form.email" :placeholder="$t('form.email.name')"></el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input type="password" v-model="form.password" :placeholder="$t('form.password.name')" maxlength=20></el-input>
      </el-form-item>
      <el-form-item prop="passwordConfirmation">
        <el-input type="password" v-model="form.passwordConfirmation" :placeholder="$t('form.passwordConfirmation.name')"></el-input>
      </el-form-item>
    </el-form>
    <p>
      {{ $t("views.signup.signInProposal") }}
      <el-button type="text" @click="loginDialogVisible = true">{{ $t("components.header.signInButton") }}</el-button>.
    </p>
  </div>
</template>

<script>
export default {
  name: 'signup',
  data() {
    return {
      form: {
        email: '',
        username: '',
        password: '',
        passwordConfirmation: '',
      },
      rules: {
        username: [
          { required: true, message: this.$t('form.username.required'), trigger: 'blur' },
          { min: 3, max: 20, message: this.$t('form.username.length', {min: 3, max: 20}), trigger: 'blur' },
        ],
        email: [
          { required: true, message: this.$t('form.email.required'), trigger: 'blur' },
          { type: 'email', message: this.$t('form.email.incorrect'), trigger: 'blur' }
        ],
        password: [
          { required: true, message: this.$t('form.password.required'), trigger: 'blur' },
          { min: 8, message: this.$t('form.password.length', {min: 8}), trigger: 'blur' }
        ],
        passwordConfirmation: [
          { required: true, message: this.$t('form.passwordConfirmation.required'), trigger: 'blur' },
        ],
      }
    }
  },
};
</script>

<style>
</style>
