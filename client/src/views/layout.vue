<template>
    <el-container direction="vertical">
      <app-header />
      <el-main>
        <el-row type="flex" class="row-bg" justify="center">
          <el-col :lg="16" :sm="24">
            <transition name="content">
              <div class="inner">
                <slot></slot>
              </div>
            </transition>
          </el-col>
        </el-row>
      </el-main>
      <el-footer class="center">
        <el-row type="flex" class="row-bg substring" justify="center">
          <el-col>asxcandrew@gmail.com Herd - federated publishing platform</el-col>
        </el-row>
        <el-row type="flex" class="row-bg" justify="center">
          <el-col><a href="https://github.com/asxcandrew/herd">Source code</a></el-col>
        </el-row>
      </el-footer>
    </el-container>
</template>

<script>
import Header from '@/components/header';

export default {
  name: 'layout',
  title: 'Herd - federated publishing platform',
  components: {
    'app-header': Header,
  },
};
</script>
