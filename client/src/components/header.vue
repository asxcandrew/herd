<template>
  <el-header>
    <el-dialog v-if="modals.signin" :visible="modals.signin" @close="closeModal('signin')" customClass="dialog-custom">
      <login/>
    </el-dialog>
    <el-dialog v-if="modals.signup" :visible="modals.signup" @close="closeModal('signup')" customClass="dialog-custom">
      <signup/>
    </el-dialog>
    <el-row type="flex" class="row-bg header-wrapper" justify="center">
      <el-col :span="14">
        <ul class="menu-left ul-unstyled">
          <li>
            <router-link to="/" class="logo unstyled-link">Herd</router-link>
          </li>
          <li class="status-bar">
            <i class="el-icon-loading" v-if="statusBar.isLoading"></i>
            <span>{{ statusBar.notice }}</span>
          </li>
        </ul>
        <ul class="menu-right ul-unstyled">
          <li>
            <div :is="plugin"></div>
          </li>
          <li v-if="loggedIn">
            <el-dropdown trigger="hover">
              <span class="el-dropdown-link userinfo-inner">
                <avatar
                  :username="session.user.name"
                  :size="36"
                />
              </span>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item @click.native="$router.push({ name: 'new-story'})">
                  {{ $t('components.header.dropdown.newStory') }}
                </el-dropdown-item>
                <el-dropdown-item @click.native="$router.push({ name: 'stories'})">
                  {{ $t('components.header.dropdown.stories') }}
                </el-dropdown-item>
                <el-dropdown-item>Bookmarks</el-dropdown-item>
                <el-dropdown-item @click.native="$router.push({name: 'public-profile', params: { username: session.user.username }})">
                  {{ $t('components.header.dropdown.profile') }}
                </el-dropdown-item>
                <el-dropdown-item @click.native="$router.push({name: 'settings'})">
                  {{ $t('components.header.dropdown.settings') }}
                </el-dropdown-item>
                <el-dropdown-item divided @click.native="logout">
                  {{ $t('components.header.dropdown.signOut') }}
                </el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </li>
          <li v-else>
            <el-button type="text" @click="showModal('signin')">
              {{ $t('components.header.signInButton') }}
            </el-button>
            <el-button type="primary" @click="showModal('signup')" plain>
              {{ $t('components.header.signUpButton') }}
            </el-button>
          </li>
        </ul>
      </el-col>
    </el-row>
  </el-header>
</template>

<script>
import Avatar from 'vue-avatar';
import { mapGetters } from 'vuex';
import Login from '@/views/login';
import Signup from '@/views/signup';
import { SIGN_OUT } from '@/store/actions.type';

const plugins = {
  default: 'search-navbar-plugin',
  story: 'story-navbar-plugin',
};

export default {
  name: 'app-header',
  computed: {
    ...mapGetters({
      session: 'session',
      modals: 'modals',
      loggedIn: 'loggedIn',
      statusBar: 'statusBar',
    }),
    plugin(){
      return (plugins[this.$route.meta.navabrPlugin] || plugins.default);
    }
  },
  components: {
    Avatar,
    Login,
    Signup,
  },
  methods: {
    // TODO: Move this logic to shared customized modal
    showModal(modal) {
      this.$store.dispatch('showModal', modal);
    },
    closeModal(modal) {
      this.$store.dispatch('closeModal', modal);
    },
    logout() {
      this.$store.dispatch(SIGN_OUT).then(() => {
        this.$router.replace({ path: '/' });
      });
    },
  },
};
</script>
<style lang="scss">
.header-wrapper {
  line-height: 60px;
  .logo {
    font-size: 20pt;
    font-weight: bold;
  }
  .status-bar {
    font-size: 9pt;
    color: #606266;
  }
  .menu-left {
    float: left;
  }
  .menu-right {
    text-align: right;
    float: right;
  }
  .userinfo-inner {
    cursor: pointer;
    img {
      width: 40px;
      height: 40px;
      border-radius: 20px;
      margin: 10px 0px 10px 10px;
      float: right;
    }
  }
}

</style>
