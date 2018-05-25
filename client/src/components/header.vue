<template>
  <el-header>
    <el-dialog :visible.sync="loginDialogVisible" customClass="dialog-custom">
      <login/>
    </el-dialog>
    <el-dialog :visible.sync="signupDialogVisible" customClass="dialog-custom">
      <signup/>
    </el-dialog>
    <el-row type="flex" class="row-bg header-wrapper" justify="center">
      <el-col :span="14">
        <div class="menu-right" v-if="session.user">
          <el-dropdown trigger="hover">
            <span class="el-dropdown-link userinfo-inner">{{"session.user.name"}}</span>
            <el-dropdown-menu slot="dropdown">
              <el-dropdown-item>Profile</el-dropdown-item>
              <el-dropdown-item>Settings</el-dropdown-item>
              <el-dropdown-item divided @click.native="logout">Logout</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
        </div>
        <div class="menu-right" v-else>
          <el-button type="text" @click="loginDialogVisible = true">{{ $t('components.header.signInButton') }}</el-button>
          <el-button type="primary" @click="signupDialogVisible = true" plain>{{ $t('components.header.signUpButton') }}</el-button>
        </div>
      </el-col>
    </el-row>
  </el-header>
</template>

<script>
import { mapGetters } from 'vuex';
import Login from '../views/login.vue';
import Signup from '../views/signup.vue';

export default {
  name: 'app-header',

  computed: mapGetters({ session: 'session' }),

  components: {
    'login': Login,
    'signup': Signup,
  },

  created() {
    this.$store.dispatch('getCurrentUser')
  },

  data() {
    return {
      loginDialogVisible: false,
      signupDialogVisible: false,
    };
  },

  methods: {
    logout() {
      this.$store.dispatch('deleteToken');
      this.$router.replace({ path: '/' });
    },
    signIn() {
      console.log('signIn');
    },
  },
};
</script>
<style lang="scss">
.header-wrapper {
  color: #333;
  line-height: 60px;
  .menu-right {
    text-align: right;
    padding-right: 35px;
    float: right;
    .userinfo-inner {
      cursor: pointer;
      color:#fff;
      img {
        width: 40px;
        height: 40px;
        border-radius: 20px;
        margin: 10px 0px 10px 10px;
        float: right;
      }
    }
  }
}

</style>
