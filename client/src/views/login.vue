<template>
  <div id="login">
    <h1>{{ $t("views.login.title") }}</h1>
    <p>{{ $t("views.login.message") }}</p>
    <el-form ref="form" :model="form" :rules="rules">
      <el-form-item>
        <el-alert v-if="loginError"
          :title="loginError"
          type="error"
          :closable="false">
        </el-alert>
      </el-form-item>
      <el-form-item prop="email">
        <el-input v-model="form.email" :placeholder="$t('form.email.name')"></el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input type="password" v-model="form.password" placeholder="Password" maxlength=20>
        </el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">{{ $t("views.login.submitButton") }}</el-button>
      </el-form-item>
    </el-form>
    <p>
      {{ $t("views.login.signUpProposal") }}
      <el-button type="text" @click="showSignup()">
        {{ $t("views.login.signUpButton") }}
      </el-button>.
    </p>
  </div>
</template>

<script>
export default {
  name: 'login',
  methods: {
    showSignup() {
      this.$store.dispatch('showModal', 'signup');
    },
    onSubmit() {
      this.$refs.form.validate((valid) => {
        if (valid) {
          this.$store.dispatch('signIn', this.form)
            .then(() => {
              // this.$router.replace({ path: this.$route.query.redirect || '/' })
              this.$store.dispatch('closeModal', 'signin');
            }).catch((error) => {
              this.loginError = error.response.data.message;
            });
        }
      });
    },
  },
  data() {
    return {
      loginError: '',
      form: {
        email: '',
        password: '',
      },
      rules: {
        email: [
          { required: true, message: this.$t('form.email.required'), trigger: 'blur' },
          { type: 'email', message: this.$t('form.email.incorrect'), trigger: 'blur' },
        ],
        password: [
          { required: true, message: this.$t('form.password.required'), trigger: 'blur' },
        ],
      },
    };
  },
};
</script>
