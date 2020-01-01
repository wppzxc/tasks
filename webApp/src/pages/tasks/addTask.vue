<template>
    <div>
        <van-nav-bar title="添加任务"
                     left-text="返回"
                     left-arrow
                     @click-left="goBack"/>
    </div>
</template>
<script>
    export default {
        methods: {
            goBack: function () {
                this.$router.go(-1)
            }
        }
    }
</script>