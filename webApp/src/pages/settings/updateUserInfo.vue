<template>
    <div>
        <van-nav-bar
                title="修改用户信息"
                left-text="返回"
                left-arrow
                @click-left="goBack">
        </van-nav-bar>
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