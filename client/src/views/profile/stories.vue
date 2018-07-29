<template>
  <div>
    <p>test</p>
    <el-tabs v-model="activeTab">
      <el-tab-pane label="Drafts" name="drafts">Drafts</el-tab-pane>
      <el-tab-pane label="Published" name="Published">Published</el-tab-pane>
    </el-tabs>
  </div>
</template>

<script>
export default {
  name: 'stories',
  data(){
    activeTab: 'drafts'
  }
};
</script>

<style>

</style>
