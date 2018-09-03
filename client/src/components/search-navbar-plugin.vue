<template>
  <div class="navbar-plugin search-plugin">
    <el-input
      size="small"
      class="search-input hidden-sm-and-down"
      placeholder="Search herd"
      >
      <i
        class="el-icon-search el-input__icon"
        slot="suffix"
      />
    </el-input>
    <el-button
      icon="el-icon-search"
      class="hidden-md-and-up"
      type="text">
    </el-button>
  </div>
</template>

<script>
export default {
  name: 'search-navbar-plugin',
  methods: {
  },
}
</script>

<style>
.search-input input{
  border: 0;
}
</style>
